package leetCode

type Triex struct {
	children [26]*Triex
	isEnd    bool
}

func Constructor() Triex {
	return Triex{}
}

func (this *Triex) Insert(word string) {
	node := this
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &Triex{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Triex) SearchPrefix(prefix string) *Triex {
	node := this
	for _, ch := range prefix {
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Triex) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd == true
}
