package listen

import (
	"encoding/json"
)

const (
	AddKws = iota + 1
	RmKws
	AddSbs
	RmSbs
)

type Notify struct {
	Action int
	Data   interface{}
}

func NewNotify(data []byte) (n *Notify, err error) {
	n = new(Notify)
	err = json.Unmarshal(data, n)
	return
}
