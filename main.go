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
	var shutdownCh = make(chan os.Signal, 1)
	var wg = &sync.WaitGroup{}

	wg.Add(1)
	go gfShutdown.FillingListAndSaving(&list, shutdownCh, wg)

	signal.Notify(shutdownCh, os.Interrupt)
	wg.Wait()
}
