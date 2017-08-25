package main

import (
	"log"
	"runtime"
	"sync"
)

type TestGroup struct {
	wg    sync.WaitGroup
	Name  string
	Tests []TestOp //TODO: This needs to be synchronized with a RW lock
	abort bool
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
	wkr := make(chan TestOp)
	numWorkers := runtime.GOMAXPROCS(1)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for t := range wkr {
				tg.wg.Add(1)
				r, err := t.Execute()
				if err != nil {
					tg.abort = true
					log.Printf("Group Run Aborting: %s\n", tg.Name)
				}
				tg.handleResult(r)
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
	close(wkr)
	log.Printf("Group Run Finished: %s\n", tg.Name)
}

func (tg *TestGroup) handleResult(res TestResult) {
	defer tg.wg.Done()
	if res.TestPass {
		//Passed
		log.Printf("Test Passed %d\n", res.TestID)
	} else {
		//Failed
		log.Printf("Test Failed %d\n", res.TestID)
	}

}
