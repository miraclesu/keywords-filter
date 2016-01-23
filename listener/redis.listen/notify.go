package listen

import (
	"encoding/json"

	. "github.com/miraclesu/keywords-filter/keyword"
)

const (
	ADD = iota + 1
	RM
)

const (
	KW = iota + 1
	SB
)

type Notify struct {
	Action int
	Data   interface{}
}

func NewNotify(kind int, data []byte) (n *Notify, err error) {
	n = new(Notify)
	if kind == KW {
		n.Data = new([]*Keyword)
	} else {
		n.Data = new([]string)
	}

	err = json.Unmarshal(data, n)
	return
}
