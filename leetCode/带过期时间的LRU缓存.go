package leetCode

import "time"

type LRUNode struct {
	key, value int
	prev, next *LRUNode
	expire_at  int64
}

type LRUCache struct {
	hash          map[int]*LRUNode
	head, tail    *LRUNode
	len, capacity int
	ttl           int64 // 过期时间
}

// 如果存在返回value并将该结点移动到头部，否则返回-1
func (this *LRUCache) Get(key int) int {
	if _, ok := this.hash[key]; ok {
		node := this.hash[key]

		if node.expire_at < time.Now().Unix() {
			this.Remove(node)
		} else {
			this.MoveHead(node)
		}
		return node.value
	}
	return -1
}

func Construct(capacity int, ttl int64) LRUCache {
	head, tail := &LRUNode{}, &LRUNode{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		head:     head,
		tail:     tail,
		ttl:      ttl,
		hash:     make(map[int]*LRUNode, capacity+1),
		len:      0,
		capacity: capacity,
	}
}

// 如果存在，修改值，并移动结点到头部；如果不存在，添加到头部，添加后如果超出容量，删除尾部结点
func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.hash[key]; ok {
		node := this.hash[key]

		if node.expire_at < time.Now().Unix() {
			this.Remove(node)
		} else {
			node.value = value
			this.MoveHead(node)
		}
	} else {
		node := &LRUNode{
			key:       key,
			value:     value,
			expire_at: time.Now().Unix() + this.ttl,
		}
		this.AddHead(node)

		if this.len > this.capacity {
			this.Remove(this.tail.prev)
		}
	}
}

func (this *LRUCache) Remove(node *LRUNode) {
	this.len--
	delete(this.hash, node.key)

	prev := node.prev
	nxt := node.next

	node.prev, node.next = nil, nil
	prev.next = nxt
	nxt.prev = prev
}

func (this *LRUCache) AddHead(node *LRUNode) {
	this.len++
	this.hash[node.key] = node

	nxt := this.head.next
	this.head.next = node
	node.prev = this.head
	node.next = nxt
	nxt.prev = node
}

func (this *LRUCache) MoveHead(node *LRUNode) {
	this.Remove(node)
	this.AddHead(node)
}
