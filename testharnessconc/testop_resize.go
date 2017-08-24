package main

import (
	"log"
	"time"
)

type TestOpResize struct {
	id    int
	state TestOpState
}

func (rt *TestOpResize) ID() int {
	return rt.id
}

func (rt *TestOpResize) State() TestOpState {
	return rt.state
}

func (rt *TestOpResize) Start(complete chan interface{}) {
	rt.state = TOSRunning
	log.Printf("Test Started: %d\n", rt.ID())
	time.Sleep(1 * time.Second)

	rt.state = TOSPassed

	//Tests
	if rt.id == 5 {
		rt.state = TOSFailed
	}
	if rt.id == 7 {
		rt.state = TOSErrored
	}

	log.Printf("Test Finished: %d\n", rt.ID())
	complete <- rt

}
