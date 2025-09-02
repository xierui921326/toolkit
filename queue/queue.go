package queue

import "sync"

type Queue[T any] struct {
	items []T
	cond  *sync.Cond
}

// GetDefaultQueue 创建一个空队列
func GetDefaultQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: []T{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

// Len 队列长度
func (q *Queue[T]) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return len(q.items)
}

// Enqueue 入队
func (q *Queue[T]) Enqueue(item T) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.items = append(q.items, item)
	q.cond.Broadcast()
}

// Dequeue 出队（非阻塞）
// 返回值 + 是否存在
func (q *Queue[T]) Dequeue() (T, bool) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// BlockingDequeue 出队（阻塞，直到有值）
func (q *Queue[T]) BlockingDequeue() T {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.items) == 0 {
		q.cond.Wait()
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// DequeueN 取出前 n 个元素（不足则返回空 slice）
func (q *Queue[T]) DequeueN(n int) []T {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.items) < n {
		return []T{}
	}
	result := q.items[:n]
	q.items = q.items[n:]
	return result
}

// DequeueAll 取出全部元素
func (q *Queue[T]) DequeueAll() []T {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	result := q.items
	q.items = []T{}
	return result
}

// Reset 清空队列
func (q *Queue[T]) Reset() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.items = []T{}
}
