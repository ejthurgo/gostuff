package main

type TestOpState uint

const (
	TOSPending TestOpState = (1 << iota) >> 1
	TOSRunning
	TOSPassed
	TOSFailed
	TOSErrored
)
