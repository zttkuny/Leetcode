package leetCode

type LFUNode struct {
	key, val   int
	freq       int
	prev, next *LFUNode
}

type LFULinkedList struct { // 存放相同频次的LFUNode，并按照最近访问时间排序
	head, tail *LFUNode
	len        int
}

func NewLFULinkedList() *LFULinkedList {
	head, tail := &LFUNode{}, &LFUNode{}
	head.next = tail
	tail.prev = head
	return &LFULinkedList{
		head: head,
		tail: tail,
		len:  0,
	}
}

func (l *LFULinkedList) Remove(node *LFUNode) {
	l.len--
	nxt := node.next
	node.prev.next = nxt
	nxt.prev = node.next
	node.next = nil
	node.prev = nil
}

func (l *LFULinkedList) AddHead(node *LFUNode) {
	nxt := l.head.next
	l.head.next = node
	nxt.prev = node
	node.prev = l.head
	node.next = nxt
	l.len++
}

func (l *LFULinkedList) MoveToHead(node *LFUNode) {
	l.Remove(node)
	l.AddHead(node)
}

type LFUCache struct {
	size     int                    // 长度
	capacity int                    // 容量
	minFreq  int                    // 频次
	nodes    map[int]*LFUNode       // key与node的映射
	Freqs    map[int]*LFULinkedList // 频次与对应双向链表的映射
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		size:     0,
		capacity: capacity,
		minFreq:  0,
		nodes:    make(map[int]*LFUNode),
		Freqs:    make(map[int]*LFULinkedList),
	}
}

func (cache *LFUCache) Get(key int) int {
	if node, ok := cache.nodes[key]; ok {
		cache.IncreaseFreq(node)
		return node.val
	}
	return -1
}

func (cache *LFUCache) Put(key int, value int) {
	if node, ok := cache.nodes[key]; ok { //如果结点存在
		node.val = value
		cache.IncreaseFreq(node) //提高频次并移动该结点
		return
	}

	if cache.size == cache.capacity {
		// 移除最低频次中最久没有使用的结点
		node := cache.Freqs[cache.minFreq].tail.prev
		delete(cache.nodes, node.key)
		cache.Freqs[cache.minFreq].Remove(node)
		cache.size--
	}

	// 如果结点不存在
	node := &LFUNode{key: key, val: value, freq: 1}
	cache.nodes[key] = node
	cache.size++

	cache.minFreq = 1
	if _, ok := cache.Freqs[1]; !ok {
		cache.Freqs[1] = NewLFULinkedList()
	}
	cache.Freqs[1].AddHead(node)
}

func (cache *LFUCache) IncreaseFreq(node *LFUNode) {
	freq := node.freq //原来的频次
	node.freq++
	lst := cache.Freqs[freq]

	lst.Remove(node)

	if lst.len == 0 && cache.minFreq == freq {
		cache.minFreq++
	}

	if _, ok := cache.Freqs[node.freq]; !ok { // 新频次不存在的话需要添加
		cache.Freqs[node.freq] = NewLFULinkedList()
	}

	cache.Freqs[node.freq].AddHead(node)
}
