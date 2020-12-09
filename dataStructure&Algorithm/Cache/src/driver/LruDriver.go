package driver

import (
	conf "../confParser"
	"container/list"
	"sync"
)

type LruDriver struct {
	maxKeyCnt 	  int
	lru           *list.List
	data          map[interface{}]*list.Element
	lock          sync.RWMutex

}

func (l* LruDriver) InitDriver(conf conf.Conf) {
	l.maxKeyCnt = conf.GetMaxKeyCnt()
	l.data = make(map[interface{}]*list.Element)
	l.lru = list.New()
}

func (l* LruDriver) Set(key interface{}, value interface{}) {
	if l.maxKeyCnt == 0 {
		return
	}
	l.lock.Lock()
	e, ok := l.data[key]
	if ok {
		l.lru.MoveToFront(e)
	} else {
		for (l.lru.Len() > l.maxKeyCnt) {
			e = l.lru.Back()
			v := e.Value.(*Value)
			delete(l.data, v.key)
			l.lru.Remove(e)
		}
		vv := &Value{
			key:       key,
			value:     value,
		}
		e := l.lru.PushFront(vv)
		l.data[key] = e
	}
	l.lock.Unlock()
}


func (l* LruDriver) Get(key interface{}) (interface{}, bool) {
	l.lock.Lock()
	var (
		elem interface{} = nil
		found = false
	)
	e, ok := l.data[key]
	if ok {
		l.lru.MoveToFront(e)
		v := e.Value.(*Value)
		elem = v.value
		found = true
	}
	l.lock.Unlock()
	return elem, found
}

