package filter

import (
	"sort"

	. "github.com/miraclesu/keywords-filter/keyword"
)

type Word struct {
	nodes  []*Word
	data   rune
	rate   int
	kind   string
	isLeaf bool
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

func (this *Word) removeNode(word *Word, isSuffix bool) (*Word, bool) {
	if len(this.nodes) == 0 {
		return this, true
	}

	for k, v := range this.nodes {
		// find the delete word
		if v.data == word.data {
			// end of the word
			if isSuffix {
				if len(v.nodes) == 0 {
					this.nodes = append(this.nodes[:k], this.nodes[k+1:]...)
					return v, true
				} else if v.isLeaf {
					v.isLeaf = false
					return v, true
				}
			} else {
				return v, false
			}
		}
	}
	return this, true
}

func (this *Word) addWord(keyword *Keyword) {
	w := this
	for _, v := range keyword.Word {
		w = w.addNode(&Word{data: v})
	}
	w.rate, w.isLeaf = keyword.Rate, true
}

func (this *Word) removeWord(keyword *Keyword) {
	ok, word := false, this
	runes := []rune(keyword.Word)
	count := len(runes) - 1
	for k, v := range keyword.Word {
		if word, ok = word.removeNode(&Word{data: v}, k == count); ok {
			break
		}
	}
}
