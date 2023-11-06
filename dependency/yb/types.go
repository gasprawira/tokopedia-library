package yb

// Created this to remove external Dependency for the wrapper

// BatchType - the batching type mode
type BatchType byte

const (
	LoggedBatch BatchType = iota
	UnloggedBatch
	CounterBatch
)
