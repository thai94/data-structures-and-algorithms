package main

type Node struct {
	key  int
	val  int
	pre  *Node
	post *Node
}

type LRUCache struct {
	kv       map[int]*Node
	lruHead  *Node
	lruTail  *Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		kv:       map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.kv[key]
	if !ok {
		return -1
	}

	// one element
	if this.lruHead == this.lruTail {
		return v.val
	}

	if this.lruHead == v {
		return v.val
	}

	if v == this.lruTail {
		tmp := v
		this.lruTail = v.pre
		this.lruTail.post = nil

		this.lruHead.pre = tmp
		tmp.post = this.lruHead
		tmp.pre = nil
		this.lruHead = tmp

		return v.val
	}

	tmp := v
	tmp.pre.post = tmp.post
	tmp.post.pre = tmp.pre

	this.lruHead.pre = tmp
	tmp.post = this.lruHead
	tmp.pre = nil
	this.lruHead = tmp

	return v.val
}

func (this *LRUCache) Put(key int, value int) {
	var node *Node
	v, ok := this.kv[key]
	if ok {
		node = v
		node.val = value
	} else {
		node = &Node{
			key: key,
			val: value,
		}
	}
	this.kv[key] = node
	if len(this.kv) == 1 {
		this.lruHead = node
		this.lruTail = node
		return
	}

	if ok {
		tmp := node
		tmp.pre.post = tmp.post
		tmp.post.pre = tmp.pre

		this.lruHead.pre = tmp
		tmp.post = this.lruHead
		tmp.pre = nil
		this.lruHead = tmp
	} else {
		this.lruHead.pre = node
		node.post = this.lruHead
		this.lruHead = node
	}

	if len(this.kv) > this.capacity {
		// remove lru element
		delete(this.kv, this.lruTail.key)
		this.lruTail = this.lruTail.pre
		this.lruTail.post = nil
		return
	}
}

func main() {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	obj.Get(1)
	obj.Get(2)
}
