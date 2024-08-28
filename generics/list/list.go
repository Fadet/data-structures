package list

import (
	"fmt"
)

type Element[T any] struct {
	next  *Element[T]
	prev  *Element[T]
	value T
}

func (e *Element[T]) Next() *Element[T] {
	return e.next
}

func (e *Element[T]) Prev() *Element[T] {
	return e.prev
}

type List[T any] struct {
	head Element[T]
	tail Element[T]
	len  int
}

func New[T any](args ...int) *List[T] {
	l := new(List[T])
	l.Clear()

	if len(args) == 1 {
		l.len = args[0]
		for i := 0; i < args[0]; i++ {
			var t T
			l.Insert(0, t)
		}
	} else if len(args) != 0 {
		panic(fmt.Sprintf("invalid arguments: list.New expects 0 or 1 arguments; found: %d", len(args)))
	}

	return l
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) Insert(index int, elems ...T) error {
	if index > l.len || index < 0 {
		return fmt.Errorf("invalid operation: index out of bounds")
	}

	if len(elems) == 0 {
		panic("invalid arguments: at least one argument must be provided")
	}

	nextEl, _ := l.AtElement(index)
	prevEl := nextEl.Prev()

	for _, elem := range elems {
		l.len++
		newEl := &Element[T]{next: nextEl, prev: prevEl, value: elem}
		nextEl.prev = newEl
		prevEl.next = newEl

		prevEl = newEl
	}

	return nil
}

func (l *List[T]) Delete(index1 int, index2 ...int) error {
	if len(index2) > 1 {
		panic(fmt.Sprint("invalid arguments: index2 must be a single value"))
	}

	if len(index2) > 1 && index1 > index2[0] {
		panic(fmt.Sprint("invalid arguments: index1 must be less than or equal to index2"))
	}

	el, err := l.AtElement(index1)
	if err != nil {
		return err
	}

	var count int
	if len(index2) == 1 {
		count = index2[0] - index1 + 1
	} else {
		count = 1
	}

	for i := 0; i < count; i++ {
		prevEl := el.Prev()
		nextEl := el.Next()
		el.next = nil
		el.prev = nil
		el = nextEl

		prevEl.next = nextEl
		nextEl.prev = prevEl
		l.len--
	}

	return nil
}

func (l *List[T]) PushBack(v T) {
	l.Insert(l.Len(), v)
}

func (l *List[T]) PushFront(v T) {
	l.Insert(0, v)
}

func (l *List[T]) PopBack() (t T, e error) {
	value, err := l.Back()
	if err != nil {
		return t, err
	}
	l.Delete(l.Len() - 1)
	return value, nil
}

func (l *List[T]) PopFront() (t T, e error) {
	value, err := l.Front()
	if err != nil {
		return t, err
	}
	l.Delete(0)
	return value, nil
}

func (l *List[T]) Back() (T, error) {
	return l.At(l.Len() - 1)
}

func (l *List[T]) Front() (T, error) {
	return l.At(0)
}

func (l *List[T]) At(index int) (t T, e error) {
	res, err := l.AtElement(index)

	if err == nil {
		return res.value, err
	} else {
		return t, err
	}
}

func (l *List[T]) AtElement(index int) (*Element[T], error) {
	if index >= l.Len() || index < 0 {
		return &l.tail, fmt.Errorf("invalid operation: index out of bounds")
	}

	el := l.Begin()
	for i := 0; i < index; i++ {
		el = el.Next()
	}
	return el, nil
}

func (l *List[T]) Begin() *Element[T] {
	return l.head.next
}

func (l *List[T]) End() *Element[T] {
	return &l.tail
}

func (l *List[T]) Clear() {
	l.head.next = &l.tail
	l.head.prev = nil
	l.tail.next = nil
	l.tail.prev = &l.head
	l.len = 0
}

func (l *List[T]) Copy() *List[T] {
	newList := New[T]()
	for el := l.Begin(); el != l.End(); el = el.Next() {
		newList.PushBack(el.value)
	}
	return newList
}
