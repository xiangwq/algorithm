package others

import (
	"container/list"
	"fmt"
	"sync"
)

type Lru struct {
	Cap  int
	Map  map[interface{}]*list.Element
	L    *list.List
	lock *sync.Mutex
}

type Node struct {
	Key   interface{}
	Value interface{}
}

func LruNew(max int) *Lru {
	return &Lru{
		Cap:  max,
		Map:  make(map[interface{}]*list.Element),
		L:    list.New(),
		lock: &sync.Mutex{},
	}
}

func (l *Lru) Get(key interface{}) (interface{}, bool) {
	if e, ok := l.Map[key]; ok {
		l.L.MoveToFront(e)
		return e.Value.(*Node).Value, true
	}
	return nil, false
}

func (l *Lru) Put(key interface{}, value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if e, ok := l.Map[key]; ok {
		e.Value.(*Node).Value = value
		l.L.MoveToFront(e)
		return
	}

	ele := l.L.PushFront(&Node{Key: key, Value: value})
	l.Map[key] = ele
	if l.L.Len() > l.Cap {
		if e := l.L.Back(); e != nil {
			l.L.Remove(e)
			node := e.Value.(*Node)
			delete(l.Map, node.Key)
			return
		}
	}
	return
}

func (l *Lru) Output() {
	if l.L == nil {
		return
	}
	node := l.L.Front()
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next()
	}
}
