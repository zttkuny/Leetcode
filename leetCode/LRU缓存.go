package leetCode

type LRUCache struct {
	size       int
	capacity   int
	hash       map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, val   int
	prev, next *DLinkedNode
}
