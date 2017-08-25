package main

import (
	"errors"
	"log"
	"time"
)

type TestOpResize struct {
	id int
}

func (rt *TestOpResize) ID() int {
	return rt.id
}

func (rt *TestOpResize) Execute() (TestResult, error) {
	res := TestResult{
		TestID:   rt.id,
		TestPass: true,
	}

	log.Printf("Test Started: %d\n", rt.ID())
	time.Sleep(1 * time.Second)

	//Tests
	if rt.id == 5 {
		res.TestPass = false
	}
	if rt.id == 7 {
		res.TestPass = false
		return res, errors.New("Test errored for some reason")
	}

	log.Printf("Test Finished: %d\n", rt.ID())
	return res, nil

}
