package filter

import (
	"github.com/miraclesu/keywords-filter/keyword"
)

// Request is used to build a request for a content test
type Request struct {
	filter  *Filter
	content []rune
	Content string

	*Response
}

// Init will init the request
func (r *Request) Init(filter *Filter) {
	r.filter, r.content = filter, []rune(r.Content)
	r.Response = &Response{
		Threshold: filter.Threshold,
	}
}

// Scan the content and return the filter result
func (r *Request) Scan() *Response {
	for i := 0; i < len(r.content); i++ {
		if node := r.trigger(r.content[i]); node != nil {
			if r.search(node, i) {
				break
			}
		}
	}
	return r.Response
}

func (r *Request) trigger(data rune) *Word {
	return r.filter.word.search(data)
}

func (r *Request) search(node *Word, index int) bool {
	//only one word
	if r.check(node, index, 0) {
		return true
	}
	for i := 1; node != nil && i < len(r.content)-index; i++ {
		c := r.content[index+i]
		tmpNode := node.search(c)
		if tmpNode == nil {
			//filter special characters
			if b, _ := r.filter.symb.search(c); b {
				continue
			}
			//is not keyword, break
			break
		}

		tmpNode.lk.RLock()
		ok := r.check(tmpNode, index, i)
		tmpNode.lk.RUnlock()
		if ok {
			return true
		}
		node = tmpNode
	}
	return false
}

func (r *Request) check(node *Word, index, i int) bool {
	if node == nil || !node.isLeaf {
		return false
	}

	r.Rate += node.rate
	r.Keywords = append(r.Keywords, keyword.Keyword{
		Rate:  node.rate,
		Index: index,
		Kind:  node.kind,
		Word:  string(r.content[index : index+i+1]),
	})
	return r.Rate >= r.filter.Threshold
}
