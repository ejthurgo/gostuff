package main

import (
	"log"
	"reflect"
	"runtime"
	"sync"
)

type TestGroup struct {
	wg         sync.WaitGroup
	Name       string
	Tests      []TestOp //TODO: This needs to be synchronized with a RW lock
	testfinish chan interface{}
	stoplisten chan bool
	abort      bool
}

func NewTestGroup(name string) *TestGroup {
	tg := TestGroup{
		Name: name,
	}
	return &tg
}

func (tg *TestGroup) Add(t TestOp) {
	tg.Tests = append(tg.Tests, t)
}

func (tg *TestGroup) Run() {
	log.Printf("Group Run Started: %s\n", tg.Name)
	tg.listen()
	wkr := make(chan TestOp)
	tg.testfinish = make(chan interface{})
	tg.stoplisten = make(chan bool)
	numWorkers := runtime.GOMAXPROCS(1)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for t := range wkr {
				tg.wg.Add(1)
				t.Start(tg.testfinish)
			}
		}()
	}

	for _, t := range tg.Tests {
		if tg.abort {
			break
		}
		wkr <- t
	}
	tg.wg.Wait()
	tg.stoplisten <- true
	close(wkr)
	close(tg.testfinish)
	close(tg.stoplisten)
	log.Printf("Group Run Finished: %s\n", tg.Name)
}

func (tg *TestGroup) listen() {
	log.Printf("Starting Listener\n")
	go func(tg *TestGroup) {
		for {
			select {
			case <-tg.stoplisten:
				return
			/*case <-tg.Abort:	//TODO: Handle abort at TestGroup level.
			log.Printf("Group Run Aborted: %s\n", tg.Name)*/
			case t := <-tg.testfinish:
				// Write the output
				switch t := t.(type) {
				case *TestOpResize:
					log.Printf("Test completed %d\n", t.ID())
					tg.HandleCompletedTest(t)
				default:
					log.Printf("Unknown completed: %s\n", reflect.TypeOf(t))
				}
				tg.wg.Done()
			}
		}
	}(tg)
}

func (tg *TestGroup) HandleCompletedTest(t TestOp) {
	switch t.State() {
	case TOSPassed:
		log.Printf("Test Passed %d\n", t.ID())
	case TOSFailed:
		log.Printf("Test Failed %d\n", t.ID())
	case TOSErrored:
		log.Printf("Test Errored %d\n", t.ID())
		tg.abort = true
	}
}
