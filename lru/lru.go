package lru

import (
	"container/list"
	"github.com/tangwenhai/demo/common"
	"sync"
)

type Node struct {
	Key string
	Value interface{}
}

type LRU struct {
	List	*list.List
	Map		sync.Map
	Size	int
	Cap		int
}

func New(n int) *LRU {
	return &LRU{
		List: list.New(),
		Size: 0,
		Cap:  n,
	}
}

func (lru *LRU)Put(key string, value interface{}) error {
	lru.List.PushFront(Node{
		Key:   key,
		Value: value,
	})
	lru.Map.Store(key, lru.List.Front())
	lru.Size++
	if lru.Size > lru.Cap {
		lru.Map.Delete(lru.List.Back().Value.(Node).Key)
		lru.List.Remove(lru.List.Back())
		lru.Size--
	}
	return nil
}

func (lru *LRU)Get(key string) (interface{}, error) {
	if addr, ok := lru.Map.Load(key); ok && addr != nil {
		lru.List.MoveToFront(addr.(*list.Element))
		return lru.List.Front().Value.(Node).Value, nil
	}
	return nil, common.ErrLRUKeyInValid
}