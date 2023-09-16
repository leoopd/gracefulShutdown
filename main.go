package main

import (
	"os"
	"os/signal"
	"sync"

	gfShutdown "github.com/leoopd/goProjects/gracefulShutdown/util"
)

func main() {

	var list string
	// shutdownCh and wg are used to guarantee that the gofuncs can return.
	var shutdownCh = make(chan struct{})
	var wg = &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		gfShutdown.FillingListAndSaving(&list, shutdownCh)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(shutdownCh)
	wg.Wait()
}
