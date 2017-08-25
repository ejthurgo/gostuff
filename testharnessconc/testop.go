package main

// Operation defines an operation
type TestOp interface {
	ID() int
	Execute() (TestResult, error)
}
