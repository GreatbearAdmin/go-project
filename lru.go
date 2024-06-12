package main

import "container/list"

type Value interface {
	Len() int
}

type Cache struct {
	maxBytes   int64
	nBytes     int64
	ll         *list.List
	cache      map[string]*list.Element
	OnEviction func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:   maxBytes,
		ll:         list.New(),
		cache:      make(map[string]*list.Element),
		OnEviction: onEvicted,
	}
}

func (c *Cache) Get(key string) (val Value, have bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToBack(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}
