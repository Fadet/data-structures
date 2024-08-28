package queue

import (
	"fmt"
	"github.com/fadet/data-structures/generics"
	"github.com/fadet/data-structures/generics/vector"
	"reflect"
)

type Queue[T any] struct {
	data generics.BiContainer[T]
}

func New[T any](container ...generics.BiContainer[T]) *Queue[T] {
	if len(container) == 0 {
		return &Queue[T]{vector.New[T]()}
	}

	if len(container) > 1 {
		panic(fmt.Sprintf("invalid arguments: queue.New() expects 0 or 1 arguments; found %d", len(container)))
	}

	return &Queue[T]{container[0]}
}

func (q *Queue[T]) Push(v T) {
	q.data.PushBack(v)
}

func (q *Queue[T]) Pop() (t T, e error) {
	res, err := q.data.PopFront()
	if err != nil {
		return t, fmt.Errorf("invalid operation: empty queue")
	}
	return res, nil
}

func (q *Queue[T]) Peek() (t T, e error) {
	res, err := q.data.Front()
	if err != nil {
		return t, fmt.Errorf("invalid operation: empty queue")
	}
	return res, nil
}

func (q *Queue[T]) Len() int {
	return q.data.Len()
}

func (q *Queue[T]) Copy() *Queue[T] {
	fn := reflect.ValueOf(q.data).MethodByName("Copy")
	if fn.IsValid() {
		res := fn.Call(nil)
		return &Queue[T]{res[0].Interface().(generics.BiContainer[T])}
	} else {
		panic("invalid operation: container must have Copy method")
	}
}
