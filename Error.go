package main

type Result[T any] interface {
    Get() *T
    HasError() bool
    Error() string
}

type Success[T any] struct {
    result T
}

func (s Success[T]) Get() *T {
    return &s.result
}

func (_ Success[T]) HasError() bool {
    return false
}

func (_ Success[T]) Error() string {
    return ""
}

type Error[T any] struct {
    error string
}

func (_ Error[T]) Get() *T {
    return nil
}

func (_ Error[T]) HasError() bool {
    return true
}

func (e Error[T]) Error() string {
    return e.error
}
