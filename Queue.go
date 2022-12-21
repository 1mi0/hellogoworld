package main

import "sync"

type queue[T any] struct {
    head, tail *queueElement[T]
}

type queueElement[T any] struct {
    value *T
    next, privious *queueElement[T]
}

func makeQueueElement[T any](value *T) *queueElement[T] {
    return &queueElement[T] { value, nil, nil }
}

func (q *queue[T])push_back(val *T) {
    q.push_back_el(makeQueueElement(val))
}

func (q *queue[T])push_back_el(el *queueElement[T]) {
    if q.head == nil && q.tail == nil {
        q.head = el
        q.tail = el
        return
    } 

    el.privious = q.head
    q.head.next = el
    q.head = el
}

func (q *queue[T])pop() *T {
    if q.tail != nil {
        defer func() {
            q.tail = q.tail.next
            if q.tail != nil {
                q.tail.privious = nil
            } else {
                q.head = nil
            }
        }()

        return q.tail.value
    }

    return nil
}

func (q *queue[T])peek() *T {
    if q.tail == nil {
        return nil
    }

    return q.tail.value
}

func (q *queue[T])empty() bool {
    return q.tail == nil
}

type syncQueue[T any] struct {
    vals *queue[T]
    lck *sync.RWMutex
    cond *sync.Cond
}

func (sq *syncQueue[T])push_back(val *T) {
    sq.lck.Lock()
    defer sq.lck.Unlock()
    defer sq.cond.Signal()
    sq.vals.push_back(val)
}

func (sq *syncQueue[T])try_push_back(val *T) bool {
    if sq.lck.TryLock() {
        defer sq.lck.Unlock()
        sq.vals.push_back(val)
        sq.cond.Signal()
        return true
    }

    return false
}

func (sq *syncQueue[T])pop() *T {
    sq.lck.Lock()
    defer sq.lck.Unlock()
    return sq.vals.pop()
}

func (sq *syncQueue[T])try_pop() *T {
    if sq.lck.TryLock() {
        defer sq.lck.Unlock()
        return sq.vals.pop()
    }

    return nil
}

func (sq *syncQueue[T])peek() *T {
    sq.lck.RLock()
    defer sq.lck.RUnlock()
    return sq.vals.peek()
}

func (sq *syncQueue[T])wait_pop() *T {
    sq.cond.L.Lock()
    defer sq.cond.L.Unlock()
    for sq.empty() {
        sq.cond.Wait()
    }

    return sq.vals.pop()
}

func (sq *syncQueue[T])empty() bool {
    sq.lck.RLock()
    defer sq.lck.RUnlock()
    return sq.vals.empty()
}
