package tests

import (
	"fmt"
	"testing"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Append(value T) *List[T] {
	newNode := &List[T]{val: value}
	if l == nil {
		return newNode
	}

	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return l
}

func (l *List[T]) PrintAll() {
	current := l
	for current != nil {
		fmt.Println(">", current.val)
		current = current.next
	}
}

func TestExercise(t *testing.T) {
	// 使用 * 號
	var list *List[int]
	list = list.Append(10)
	list = list.Append(20)
	list = list.Append(30)
	list.PrintAll()

	// 不使用 * 號
	var list2 = List[string]{val: "s"}
	list2.Append("Leo")
	list2.Append("Evonne")
	list2.Append("Mou")
	list2.PrintAll()

}
