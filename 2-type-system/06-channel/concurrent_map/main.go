package main

import (
	"context"
	"sync"
	"time"
)

/*
 要求实现一个map
（1）面向高并发；
（2）只存在插入和查询操作O(1)；
（3）查询时，若 key 存在，直接返回 val；若 key 不存在，阻塞直到 key val； 对被放入后，获取 val 返回；等待指定的时长依然未放入，返回超时错误；
（4）用go语言实现，不能有死锁或者 panic 风险。
*/

type ConcurrentMap struct {
	sync.Mutex
	mp      map[int]int
	keyToCh map[int]chan struct{}
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		mp:      make(map[int]int),
		keyToCh: make(map[int]chan struct{}),
	}
}

func (m *ConcurrentMap) Set(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v
	ch, ok := m.keyToCh[k]
	if !ok {
		return
	}
	select {
	case <-ch:
		return
	default:
		close(ch)
	}
}

func (m *ConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.Lock()
	v, ok := m.mp[k]
	if ok {
		m.Unlock()
		return v, nil
	}
	ch, ok := m.keyToCh[k]
	if !ok {
		ch = make(chan struct{})
		m.keyToCh[k] = ch
	}

	tCtx, cancel := context.WithTimeout(context.Background(), maxWaitingDuration)
	defer cancel()
	m.Unlock()

	select {
	case <-tCtx.Done(): //超时控制
		return -1, tCtx.Err()
	case <-ch: //读阻塞

	}

	m.Lock()
	defer m.Unlock()
	return m.mp[k], nil
}

func main() {

}
