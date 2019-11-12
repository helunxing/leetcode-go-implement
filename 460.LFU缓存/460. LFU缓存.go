package leetcode

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

import (
	"container/list"
)

type LFUCache struct {
	Cap    int
	Cntm   map[int]*list.List
	KV     map[int]*list.Element
	Lowest int
}

type node struct {
	Key   int
	Value int
	Cnt   int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Cap:    capacity,
		Cntm:   make(map[int]*list.List, capacity),
		KV:     make(map[int]*list.Element, capacity),
		Lowest: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	e, ok := this.KV[key]
	if !ok {
		return -1
	}

	n := e.Value.(*node)
	this.Cntm[n.Cnt].Remove(e)
	// 需要同时满足
	if this.Cntm[n.Cnt].Len() == 0 && n.Cnt == this.Lowest {
		this.Lowest += 1
	}
	n.Cnt += 1

	l, ok := this.Cntm[n.Cnt]
	if !ok {
		l = list.New()
		this.Cntm[n.Cnt] = l
	}
	// 需要更新节点地址
	this.KV[key] = l.PushFront(n)
	return n.Value
}

func (this *LFUCache) Put(key int, value int) {
	if this.Cap <= 0 {
		return
	}
	e, ok := this.KV[key]
	if ok {

		n := e.Value.(*node)
		this.Cntm[n.Cnt].Remove(e)
		if this.Cntm[n.Cnt].Len() == 0 && n.Cnt == this.Lowest {
			this.Lowest += 1
		}
		n.Cnt += 1

		n.Value = value
		l, ok := this.Cntm[n.Cnt]
		if !ok {
			l = list.New()
			this.Cntm[n.Cnt] = l
		}
		// 需要更新地址
		this.KV[key] = l.PushFront(n)
	} else {
		n := &node{
			Key:   key,
			Value: value,
			Cnt:   0,
		}
		var l *list.List
		var ok bool
		if len(this.KV) < this.Cap {
			l, ok = this.Cntm[0]
			if !ok {
				l = list.New()
				this.Cntm[0] = l
			}
			l.PushFront(n)
		} else {
			for len(this.KV) >= this.Cap {
				rmE := this.Cntm[this.Lowest].Back()
				reK := rmE.Value.(*node).Key
				this.Cntm[this.Lowest].Remove(rmE)
				delete(this.KV, reK)
			}
			// 如果满了，则必然存在cntm[0]
			l = this.Cntm[0]
			l.PushFront(n)
		}
		this.KV[key] = l.Front()
		this.Lowest = 0
	}
}

func (this *LFUCache) Upgarde() {

}
