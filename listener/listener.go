package listener

import (
	. "github.com/miraclesu/keywords-filter/keyword"
)

type Listener interface {
	AddKeywords() <-chan []*Keyword
	RemoveKeywords() <-chan []*Keyword
	AddSymbols() <-chan []string
	RemoveSymbols() <-chan []string
}
