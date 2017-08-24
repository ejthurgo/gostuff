package main

// Operation defines an operation
type TestOp interface {
	ID() int
	Start(chan interface{})
	State() TestOpState
}
