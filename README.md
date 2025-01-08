# Go Concurrency Bug: Race Condition in Channel Close

This repository demonstrates a subtle race condition that can occur in Go programs using channels and wait groups for concurrency. The bug is related to the synchronization of closing a channel and the receiving goroutine.  The program may exhibit unpredictable behaviour such as missed messages or panics. The `bug.go` file contains the buggy code, while `bugSolution.go` provides a corrected version.

## Bug Description

The original program has a race condition where the sending goroutine closes the channel before the receiving goroutine has finished reading all the values.  This is particularly insidious because it might not appear to be a problem in simple test cases.