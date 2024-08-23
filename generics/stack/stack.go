package stack

import (
	"fmt"
	"github.com/fadet/data-structures/generics"
	"github.com/fadet/data-structures/generics/vector"
	"reflect"
)

type Stack[T any] struct {
	data generics.BackContainer[T]
}

func New[T any](container ...generics.BackContainer[T]) *Stack[T] {
	if len(container) == 0 {
		return &Stack[T]{vector.New[T]()}
	}

	if len(container) > 1 {
		panic(fmt.Sprintf("invalid arguments: stack.New() expects 0 or 1 arguments; found %d", len(container)))
	}

	return &Stack[T]{container[0]}
}

func (st *Stack[T]) Push(v T) {
	st.data.PushBack(v)
}

func (st *Stack[T]) Pop() (t T, e error) {
	res, err := st.data.PopBack()
	if err != nil {
		return t, fmt.Errorf("invalid operation: empty stack")
	}
	return res, nil
}

func (st *Stack[T]) Peek() (t T, e error) {
	res, err := st.data.Back()
	if err != nil {
		return t, fmt.Errorf("invalid operation: empty stack")
	}
	return res, nil
}

func (st *Stack[T]) Len() int {
	return st.data.Len()
}

func (st *Stack[T]) Copy() *Stack[T] {
	fn := reflect.ValueOf(st.data).MethodByName("Copy")
	if fn.IsValid() {
		res := fn.Call(nil)
		return &Stack[T]{res[0].Interface().(generics.BackContainer[T])}
	} else {
		panic("invalid operation: container must have Copy method")
	}
}
