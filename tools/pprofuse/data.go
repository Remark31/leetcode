package pprofuse

import (
	// "github.com/emirpasic/gods/maps/hashmap"
	"sync"
)

type state struct {
	value interface{}
	next  *state
	done  chan struct{}
}

func newState(value interface{}) *state {
	return &state{
		value: value,
		done:  make(chan struct{}),
	}
}

func (s *state) update(value interface{}) *state {
	s.next = newState(value)
	close(s.done)
	return s.next
}

type Stream interface {

	//获取当前的值
	Value() interface{}
	//新内容到达时尝试去改变
	Changes() chan struct{}
	//寻找下一个订阅者
	Next() interface{}
	//判断是否还存在下一个订阅者
	HasNext() bool
	// 等待当前订阅被关闭
	WaitNext() interface{}
	// 复制当前订阅
	Clone() Stream
}

type stream struct {
	state *state
}

func (s *stream) Clone() Stream {
	return &stream{state: s.state}
}

func (s *stream) Value() interface{} {
	return s.state.value
}

func (s *stream) Changes() chan struct{} {
	return s.state.done
}

func (s *stream) Next() interface{} {
	s.state = s.state.next
	return s.state.value
}

func (s *stream) HasNext() bool {
	select {
	case <-s.state.done:
		return true
	default:
		return false
	}
}

func (s *stream) WaitNext() interface{} {
	<-s.state.done
	s.state = s.state.next
	return s.state.value
}

type Property interface {
	Value() interface{}

	Update(value interface{})

	Observe() Stream
}

func NewProperty(value interface{}) Property {
	return &property{state: newState(value)}
}

type property struct {
	sync.RWMutex
	state *state
}

func (p *property) Value() interface{} {
	p.RLock()
	defer p.RUnlock()
	return p.state.value
}

func (p *property) Update(value interface{}) {
	p.Lock()
	defer p.Unlock()
	p.state = p.state.update(value)
}

func (p *property) Observe() Stream {
	p.RLock()
	defer p.RUnlock()
	return &stream{state: p.state}
}
