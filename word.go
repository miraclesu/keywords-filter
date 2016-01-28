package filter

import (
	"sort"
	"sync"

	"github.com/miraclesu/keywords-filter/keyword"
)

// Word is used to build a keyword's tree
type Word struct {
	lk     *sync.RWMutex
	nodes  []*Word
	data   rune
	rate   int
	kind   string
	isLeaf bool
}

func (w *Word) search(data rune) *Word {
	if len(w.nodes) == 0 {
		return nil
	}
	index := sort.Search(len(w.nodes)-1, func(i int) bool { return w.nodes[i].data >= data })
	if w.nodes[index].data == data {
		return w.nodes[index]
	}
	return nil
}

func (w *Word) addNode(word *Word) *Word {
	w.lk.Lock()
	defer w.lk.Unlock()

	if len(w.nodes) == 0 {
		w.nodes = append(w.nodes, word)
		return word
	}
	for k, v := range w.nodes {
		if v.data == word.data {
			return v
		} else if v.data > word.data {
			w.nodes = append(w.nodes[:k], append([]*Word{word}, w.nodes[k:]...)...)
			return word
		}
	}
	w.nodes = append(w.nodes, word)
	return word
}

func (w *Word) removeNode(word *Word, isSuffix bool) (*Word, bool) {
	w.lk.Lock()
	defer w.lk.Unlock()

	if len(w.nodes) == 0 {
		return w, true
	}

	for k, v := range w.nodes {
		// find the delete word
		if v.data == word.data {
			// end of the word
			if isSuffix {
				if len(v.nodes) == 0 {
					w.nodes = append(w.nodes[:k], w.nodes[k+1:]...)
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
	return w, true
}

func (w *Word) addWord(kw *keyword.Keyword) {
	word := w
	for _, v := range kw.Word {
		word = word.addNode(&Word{lk: new(sync.RWMutex), data: v})
	}

	word.lk.Lock()
	word.rate, word.kind, word.isLeaf = kw.Rate, kw.Kind, true
	word.lk.Unlock()
}

func (w *Word) removeWord(kw *keyword.Keyword) {
	ok, word := false, w
	runes := []rune(kw.Word)
	count := len(runes) - 1
	for k, v := range kw.Word {
		if word, ok = word.removeNode(&Word{data: v}, k == count); ok {
			break
		}
	}
}
