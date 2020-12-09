package driver

import (
	conf "../confParser"
	"container/list"
	"sync"
)

type FiFoDriver struct {
	maxKeyCnt 	  int
	fifo           *list.List
	data          map[interface{}]*list.Element
	lock          sync.RWMutex
}

func (l* FiFoDriver) InitDriver(conf conf.Conf) {
	l.maxKeyCnt = conf.GetMaxKeyCnt()
	l.data = make(map[interface{}]*list.Element)
	l.fifo = list.New()
}

func (l* FiFoDriver) Set(key interface{}, value interface{}) {
	if l.maxKeyCnt == 0 {
		return
	}
	l.lock.Lock()
	e, ok := l.data[key]
	if !ok {
		for (l.fifo.Len() > l.maxKeyCnt) {
			l.fifo.Remove(l.fifo.Front())
			delete(l.data, key)
		}
		e := l.lru.PushFront(value)
		l.data[key] = e
	}
	l.lock.Unlock()
}


func (l* FiFoDriver) Get(key interface{}) (interface{}, bool) {
	l.lock.RLocker()
	var (
		elem interface{} = nil
		found = false
	)
	e, ok := l.data[key]
	if ok {
		v := e.Value.(*Value)
		elem = v.value
		found = true
	}
	l.lock.RUnlock()
	return elem, found
}

