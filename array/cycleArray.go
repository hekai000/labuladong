package main

import "errors"

type CycleArray[T any] struct {
	arr   []T
	start int
	end   int
	count int
	size  int
}

func NewCycleArray[T any]() *CycleArray[T] {
	return NewCycleArrayWithSize[T](1)
}

func NewCycleArrayWithSize[T any](size int) *CycleArray[T] {
	return &CycleArray[T]{
		arr:   make([]T, size),
		start: 0,
		end:   0,
		count: 0,
		size:  size,
	}
}

func (ca *CycleArray[T]) resize(newSize int) {
	newArr := make([]T, newSize)
	for i := 0; i < ca.count; i++ {
		newArr[i] = ca.arr[(ca.start+i)%ca.size]
	}
	ca.arr = newArr
	ca.start = 0
	ca.end = ca.count
	ca.size = newSize
}

func (ca *CycleArray[T]) AddFirst(val T) {
	if ca.isFull() {
		ca.resize(ca.size * 2)
	}
	ca.start = (ca.start - 1 + ca.size) % ca.size
	ca.arr[ca.start] = val
	ca.count++
}

func (ca *CycleArray[T]) isFull() bool {
	return ca.count == ca.size
}

func (ca *CycleArray[T]) Size() int {
	return ca.count
}

func (ca *CycleArray[T]) isEmpty() bool {
	return ca.count == 0
}

func (ca *CycleArray[T]) RemoveFirst() error {
	if ca.isEmpty() {
		return errors.New("array is empty")
	}
	ca.arr[ca.start] = *new(T)
	ca.start = (ca.start + 1) % ca.size
	ca.count--
	if ca.count > 0 && ca.count == ca.size/4 {
		ca.resize(ca.size / 2)
	}
	return nil
}

func (ca *CycleArray[T]) AddLast(val T) {
	if ca.isFull() {
		ca.resize(ca.size * 2)
	}

	ca.arr[ca.end] = val
	ca.end = (ca.end + 1) % ca.size
	ca.count++
}

func (ca *CycleArray[T]) RemoveLast() error {
	if ca.isEmpty() {
		return errors.New("array is empty")
	}
	ca.end = (ca.end - 1 + ca.size) % ca.size

	ca.arr[ca.end] = *new(T)

	ca.count--

	if ca.count > 0 && ca.count == ca.size/4 {
		ca.resize(ca.size / 2)
	}
	return nil
}

func (ca *CycleArray[T]) GetFirst() (T, error) {
	if ca.isEmpty() {
		return *new(T), errors.New("array is empty")
	}
	return ca.arr[ca.start], nil
}
func (ca *CycleArray[T]) GetLast() (T, error) {
	if ca.isEmpty() {
		return *new(T), errors.New("array is empty")
	}
	return ca.arr[(ca.end-1+ca.size)%ca.size], nil
}
