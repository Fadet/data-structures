package vector

import (
	"fmt"
)

type Vector[T any] struct {
	data []T
}

func New[T any](args ...int) *Vector[T] {
	switch len(args) {
	case 0:
		return &Vector[T]{data: make([]T, 0)}
	case 1:
		return &Vector[T]{data: make([]T, args[0])}
	case 2:
		return &Vector[T]{data: make([]T, args[0], args[1])}
	default:
		panic(fmt.Sprintf("invalid arguments: vector.New() expects from 0 to 2 arguments; found %d", len(args)))
	}
}

func (v *Vector[T]) Len() int {
	return len(v.data)
}

func (v *Vector[T]) Cap() int {
	return cap(v.data)
}

func (v *Vector[T]) Clear() {
	v.data = make([]T, 0)
}

func (v *Vector[T]) PushBack(elem T) {
	v.Insert(v.Len(), elem)
}

func (v *Vector[T]) PushFront(elem T) {
	v.Insert(0, elem)
}

func (v *Vector[T]) PopBack() (t T, e error) {
	if v.Len() == 0 {
		return t, fmt.Errorf("invalid operation: empty vector")
	}

	t = v.data[len(v.data)-1]
	v.data = v.data[:len(v.data)-1]
	return t, nil
}

func (v *Vector[T]) PopFront() (t T, e error) {
	if v.Len() == 0 {
		return t, fmt.Errorf("invalid operation: empty vector")
	}

	t = v.data[0]
	v.data = v.data[1:]
	return t, nil
}

func (v *Vector[T]) Back() (t T, e error) {
	if v.Len() == 0 {
		return t, fmt.Errorf("invalid operation: empty vector")
	}

	return v.data[len(v.data)-1], nil
}

func (v *Vector[T]) Front() (t T, e error) {
	if v.Len() == 0 {
		return t, fmt.Errorf("invalid operation: empty vector")
	}

	return v.data[0], nil
}

func (v *Vector[T]) Insert(index int, elems ...T) {
	v.data = append(v.data[:index], append(elems, v.data[index:]...)...)
}

func (v *Vector[T]) Delete(index1 int, index2 ...int) error {
	if len(index2) > 1 {
		panic(fmt.Sprint("invalid arguments: index2 must not be a single value"))
	}

	if len(index2) > 0 && index1 > index2[0] {
		panic(fmt.Sprint("invalid arguments: index1 must be less than or equal to index2"))
	}

	if index1 < 0 || index1 >= v.Len() {
		return fmt.Errorf("invalid operation: index out of bounds")
	}

	if len(index2) == 1 {
		v.data = append(v.data[:index1], v.data[index2[0]+1:]...)
	} else {
		v.data = append(v.data[:index1], v.data[index1+1:]...)
	}

	return nil
}

func (v *Vector[T]) Data() []T {
	return v.data
}

func (v *Vector[T]) Copy() *Vector[T] {
	res := make([]T, len(v.data))
	copy(res, v.data)
	return &Vector[T]{res}
}

func (v *Vector[T]) At(i int) (t T, e error) {
	if i < 0 || i >= v.Len() {
		return t, fmt.Errorf("invalid operation: index out of bounds")
	}

	return v.data[i], nil
}
