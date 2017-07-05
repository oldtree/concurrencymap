package ConcurrencyMap

import (
	"sync"
)

type ConcurrencyMapWithLock struct {
	data map[interface{}]interface{}
	sync.Mutex
}

func (c *ConcurrencyMapWithLock) Add(key, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}
func (c *ConcurrencyMapWithLock) Remove(key interface{}) {
	c.Lock()
	defer c.Unlock()
	delete(c.data, key)
}
func (c *ConcurrencyMapWithLock) Get(key interface{}) interface{} {
	c.Lock()
	defer c.Unlock()
	return c.data[key]
}

type Operation struct {
	OpType int
	key    interface{}
	value  interface{}
}

type ConcurrencyMapWithChan struct {
	data       map[interface{}]interface{}
	ChanOp     chan *Operation
	ChanResult chan *Operation
}

func (c *ConcurrencyMapWithChan) OperationLoop() {
	c.ChanOp = make(chan *Operation, 1024)
	c.ChanResult = make(chan *Operation, 1024)
	for {
		select {
		case o := <-c.ChanOp:
			switch o.OpType {
			case 1: //get
				o.value = c.data[o.key]
				c.ChanResult <- o
			case 2: //add
				c.data[o.key] = o.value
				c.ChanResult <- o
			case 3: ///remove
				delete(c.data, o.key)
				c.ChanResult <- o
			}
		}
	}
}

func (c *ConcurrencyMapWithChan) Add(key, value interface{}) {
	op := new(Operation)
	op.key = key
	op.value = value
	op.OpType = 2
	c.ChanOp <- op
	op = <-c.ChanResult
}

func (c *ConcurrencyMapWithChan) Remove(key interface{}) {
	op := new(Operation)
	op.key = key
	op.value = nil
	op.OpType = 3
	c.ChanOp <- op
	op = <-c.ChanResult
}

func (c *ConcurrencyMapWithChan) Get(key interface{}) interface{} {
	op := new(Operation)
	op.key = key
	op.value = nil
	op.OpType = 1
	c.ChanOp <- op
	op = <-c.ChanResult
	return op.value
}

type ConcurrencyMapWithRWLock struct {
	data map[interface{}]interface{}
	sync.RWMutex
}

func (c *ConcurrencyMapWithRWLock) Add(key, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}
func (c *ConcurrencyMapWithRWLock) Remove(key interface{}) {
	c.Lock()
	defer c.Unlock()
	delete(c.data, key)
}
func (c *ConcurrencyMapWithRWLock) Get(key interface{}) interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.data[key]
}
