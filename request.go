package filter

type Request struct {
	filter  *Filter
	content []rune
	rate    int
}

func NewRequest(filter *Filter, content string) *Request {
	return &Request{
		filter:  filter,
		content: []rune(content),
	}
}

func (this *Request) scan() bool {
	for i := 0; i < len(this.content); i++ {
		if node := this.trigger(this.content[i]); node != nil {
			if this.search(node, i) {
				return true
			}
		}
	}
	return false
}

func (this *Request) trigger(data rune) *Word {
	return this.filter.word.search(data)
}

func (this *Request) search(node *Word, index int) bool {
	//only one word
	if this.check(node, index, 0) {
		return true
	}
	for i := 1; node != nil && i < len(this.content)-index; i++ {
		c := this.content[index+i]
		tmpNode := node.search(c)
		if tmpNode == nil {
			//filter special characters
			if b, _ := this.filter.symb.search(c); b {
				continue
			}
			//is not keyword, break
			break
		}
		if this.check(tmpNode, index, i) {
			return true
		}
		node = tmpNode
	}
	return false
}

func (this *Request) check(node *Word, index, i int) bool {
	if node != nil && node.isLeaf {
		this.rate += node.rate
		return this.rate >= this.filter.Threshold
	}
	return false
}
