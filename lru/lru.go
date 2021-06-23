package lru

import "container/list"

type Cache struct {
	// 缓存能容纳的最大数量，如果超过这个数，就会驱逐数据
	// 0 代表没有限制
	MaxEnteries int

	ll    *list.List
	cache map[interface{}]*list.Element
}

type Key interface{}

type entry struct {
	Key   Key
	value interface{}
}

func New(maxEnteries int) *Cache {
	return &Cache{
		MaxEnteries: maxEnteries,
		ll:          list.New(),
		cache:       make(map[interface{}]*list.Element),
	}
}

func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}

	if entryEle, ok := c.cache[key]; ok {
		c.ll.MoveToFront(entryEle)            // 最近访问过的元素移动到链表头部
		entryEle.Value.(*entry).value = value // 覆盖value的值
		return
	}

	entryEle := c.ll.PushFront(&entry{Key: key, value: value})
	c.cache[key] = entryEle
	// 判断是否需要驱逐
	if c.MaxEnteries != 0 && c.ll.Len() > c.MaxEnteries {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key Key) interface{} {
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return nil
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.Key)
	}
}
