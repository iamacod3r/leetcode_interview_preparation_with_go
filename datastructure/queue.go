package datastructure

import (
	"container/list"
	"fmt"
	"sync"
)

type Queue struct {
	queue *list.List
}

func (q *Queue) Enqueue(value any) {
	q.queue.PushBack(value)
}

func (q *Queue) Dequeue() (any, error) {
	if q.queue.Len() > 0 {
		element := q.queue.Front()
		q.queue.Remove(element)
		return element, nil
	}
	return nil, fmt.Errorf("dequeue error : queue is empty")
}

func (q *Queue) Peek() (any, error) {
	if q.queue.Len() > 0 {
		return q.queue.Front().Value, nil
	}

	return nil, fmt.Errorf("peek error: queue is empty")
}

func (q *Queue) Size() int {
	return q.queue.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.queue.Len() == 0
}

type QueueSlice struct {
	queue []any
	lock  sync.RWMutex
}

func (q *QueueSlice) Enqueue(value any) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, value)
}

func (q *QueueSlice) Dequeue() (any, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) > 0 {
		element := q.queue[0]
		q.queue = q.queue[1:]
		return element, nil
	}

	return nil, fmt.Errorf("dequeue error: queue is empty")
}

func (q *QueueSlice) Peek() (any, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) > 0 {
		return q.queue[0], nil
	}

	return nil, fmt.Errorf("peek error: queue is empty")
}

func (q *QueueSlice) Size() int {
	return len(q.queue)
}

func (q *QueueSlice) IsEmpty() bool {
	return len(q.queue) == 0
}
