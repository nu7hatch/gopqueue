// This package provides a priority queue implementation and
// scaffold interfaces.
//
// Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>
package pqueue

import (
	"container/heap"
	"sync"
	"errors"
	"time"
)

type Interface interface {
	Priority() int
}

type sorter struct {
	m []Interface
}

func (s *sorter) Push(i interface{}) {
	item, ok := i.(Interface)
	if !ok {
		return
	}
	s.m = append(s.m[:], item)
}

func (s *sorter) Pop() (x interface{}) {
	if s.Len() > 0 {
		x, s.m = s.m[s.Len()-1], s.m[:s.Len()-1]
	}
	return
}

func (s *sorter) Len() int {
	return len(s.m[:])
}

func (s *sorter) Less(i, j int) bool {
	return s.m[i].Priority() < s.m[j].Priority()
}

func (s *sorter) Swap(i, j int) {
	if s.Len() > 0 {
		s.m[i], s.m[j] = s.m[j], s.m[i]
	}
}
	
type Queue struct {
	Limit    int
	items    *sorter
	mtx      sync.Mutex
}

func New(max int) (q *Queue) {
	q = &Queue{Limit: max}
	q.items = new(sorter)
	heap.Init(q.items)
	return
}

func (q *Queue) Enqueue(item Interface) (err error) {
	if q.Limit > 0 && q.Len() >= q.Limit {
		return errors.New("Queue limit reached")
	}
	q.mtx.Lock()
	heap.Push(q.items, item)
	q.mtx.Unlock()
	return
}

func (q *Queue) Dequeue() (item Interface) {
start:
	q.mtx.Lock()
	x := heap.Pop(q.items)
	q.mtx.Unlock()
	if x == nil {
		<-time.After(1)
		goto start
	}
	item = x.(Interface)
	return
}

func (q *Queue) Len() int {
	return q.items.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}