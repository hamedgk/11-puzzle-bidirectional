package data_structures

import (
	"container/list"
)

type Queue struct {
	queue *list.List
}

func NewQueue() Queue {
	return Queue{queue: list.New()}
}

func (list Queue) Enqueue(v Node) {
	list.queue.PushBack(v)
}

func (list Queue) IsEmpty() bool {
	return list.queue.Len() == 0
}

func (list Queue) Dequeue() (Node, bool) {
	first := list.queue.Front()
	ret, ok := list.queue.Remove(first).(Node)
	return ret, ok
}

func (list Queue) Front() *list.Element {
	return list.queue.Front()
}

func (list Queue) Len() int {
	return list.queue.Len()
}
