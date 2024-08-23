package generics

type BackContainer[T any] interface {
	inter[T]

	PushBack(T)
	PopBack() (T, error)
	Back() (T, error)
}

type FrontContainer[T any] interface {
	inter[T]

	PushFront(T)
	PopFront() (T, error)
	Front() (T, error)
}

type BiContainer[T any] interface {
	BackContainer[T]
	FrontContainer[T]
}

type inter[T any] interface {
	Len() int
	Clear()
	At(int) (T, error)
}
