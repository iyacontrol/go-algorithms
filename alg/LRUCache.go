package alg

import "container/list"

type LRUCache struct {
	capacity int
	list     *list.List
	vals     map[int]*list.Element
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		vals:     make(map[int]*list.Element),
	}
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.vals[key]; ok {
		c.list.MoveToFront(v)
		return v.Value.(entry).value
	}

	return -1
}

func (c *LRUCache) Put(key, value int) {
	if v, ok := c.vals[key]; ok {
		v.Value = entry{key, value}
		c.list.MoveToFront(v)
		return
	}

	c.vals[key] = c.list.PushFront(entry{key, value})
	if len(c.vals) > c.capacity {
		delete(c.vals, c.list.Remove(c.list.Back()).(entry).key)
	}

}
