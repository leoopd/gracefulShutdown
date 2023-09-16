# gracefulShutdown

#### Description
a simple implementation of a graceful shutdown utilizing signal.Notify()

#### Mechanism
```golang
...
    var list string
	// shutdownCh and wg guarantee, that the goroutine can return.
	var shutdownCh = make(chan os.Signal, 1)
	var wg = &sync.WaitGroup{}

	wg.Add(1)
	go gfShutdown.FillingListAndSaving(&list, shutdownCh, wg)

	// Notifies on shutdownCh once Strg+C interrupts which makes FillingListAndSaving() return
	signal.Notify(shutdownCh, os.Interrupt)
	wg.Wait()
...
```