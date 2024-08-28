package queue_test

import (
	"github.com/fadet/data-structures/generics/list"
	"github.com/fadet/data-structures/generics/queue"
	"testing"
)

func assert(t *testing.T, fn func() bool) {
	if !fn() {
		t.Fatal("assert failed")
	}
}

func TestQueueLen(t *testing.T) {
	q := queue.New[int](list.New[int]())
	assert(t, func() bool {
		return q.Len() == 0
	})
	q.Push(42)
	assert(t, func() bool {
		return q.Len() == 1
	})
}

func TestQueuePeek(t *testing.T) {
	q := queue.New[int](list.New[int]())
	q.Push(3)
	q.Push(42)
	assert(t, func() bool { res, _ := q.Peek(); return res == 3 })
}

func TestQueuePeekError(t *testing.T) {
	q := queue.New[int](list.New[int]())
	assert(t, func() bool { _, err := q.Peek(); return err != nil })
}

func TestQueuePush(t *testing.T) {
	q := queue.New[int](list.New[int]())
	q.Push(42)
	assert(t, func() bool {
		res, _ := q.Peek()
		return res == 42 && q.Len() == 1
	})
}

func TestQueuePop(t *testing.T) {
	q := queue.New[int](list.New[int]())
	q.Push(42)
	q.Push(3)
	assert(t, func() bool {
		res, _ := q.Pop()
		return res == 42 && q.Len() == 1
	})
	assert(t, func() bool {
		res, _ := q.Pop()
		return res == 3 && q.Len() == 0
	})
}

func TestQueuePopError(t *testing.T) {
	q := queue.New[int](list.New[int]())
	assert(t, func() bool { _, err := q.Pop(); return err != nil })
}

func TestQueueCopy(t *testing.T) {
	q := queue.New[int](list.New[int]())
	q.Push(42)
	q.Push(3)
	q.Push(5)
	q1 := q.Copy()
	assert(t, func() bool {
		if q.Len() != q1.Len() {
			return false
		}
		for q.Len() != 0 {
			el, _ := q.Pop()
			el1, _ := q1.Pop()
			if el != el1 {
				return false
			}
		}
		return true
	})
}
