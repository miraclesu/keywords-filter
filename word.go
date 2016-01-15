package filter

import (
	"sort"

	. "github.com/miraclesu/keywords-filter/keyword"
)

const (
	Leaf = -1
)

type Word struct {
	nodes []*Word
	data  rune
	rate  int16
	tag   int8
}

func (this *Word) search(data rune) *Word {
	if len(this.nodes) == 0 {
		return nil
	}
	index := sort.Search(len(this.nodes)-1, func(i int) bool { return this.nodes[i].data >= data })
	if this.nodes[index].data == data {
		return this.nodes[index]
	}
	return nil
}

func (this *Word) isLeaf() bool {
	return this.tag == Leaf
}

func (this *Word) addNode(word *Word) *Word {
	if len(this.nodes) == 0 {
		this.nodes = append(this.nodes, word)
		return word
	}
	for k, v := range this.nodes {
		if v.data == word.data {
			return v
		} else if v.data > word.data {
			this.nodes = append(this.nodes[:k], append([]*Word{word}, this.nodes[k:]...)...)
			return word
		}
	}
	this.nodes = append(this.nodes, word)
	return word
}

func (this *Word) removeNode(word *Word) (*Word, bool) {
	if len(this.nodes) == 0 {
		return this, true
	}
	for k, v := range this.nodes {
		if v.data == word.data {
			if len(v.nodes) <= 1 {
				this.nodes = append(this.nodes[:k], this.nodes[k+1:]...)
				return this, true
			} else {
				return v, false
			}
		}
	}
	return this, true
}

func (this *Word) addWord(keyword *Keyword) {
	for _, v := range keyword.Word {
		this = this.addNode(&Word{data: v})
	}
	this.rate, this.tag = keyword.Rate, Leaf
}

func (this *Word) removeWord(keyword *Keyword) {
	ok, word := false, this
	for _, v := range keyword.Word {
		if word, ok = word.removeNode(&Word{data: v}); ok {
			break
		}
	}
}
