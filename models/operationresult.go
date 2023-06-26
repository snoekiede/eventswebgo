package models

type OperationResult[T any] struct {
	Value  *T
	Result error
}

func ConstructWithoutError[T any](value *T) OperationResult[T] {
	return OperationResult[T]{
		Result: nil,
		Value:  value,
	}
}

func ConstructWithError[T any](errorResult error) OperationResult[T] {
	return OperationResult[T]{
		Result: errorResult,
		Value:  nil,
	}
}
