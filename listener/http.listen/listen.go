package listen

import . "github.com/miraclesu/keywords-filter/keyword"

type Listener struct {
	AddKeywordsChan    chan []*Keyword
	RemoveKeywordsChan chan []*Keyword
	AddSymbolsChan     chan []*Keyword
	RemoveSymbolsChan  chan []*Keyword
}

func NewListener() *Listener {
	l := &Listener{
		AddKeywordsChan:    make(chan []*Keyword, 8),
		RemoveKeywordsChan: make(chan []*Keyword, 8),
		AddSymbolsChan:     make(chan []*Keyword, 8),
		RemoveSymbolsChan:  make(chan []*Keyword, 8),
	}
	return l
}

func (this *Listener) AddKeywords() <-chan []*Keyword {
	return this.AddKeywordsChan
}

func (this *Listener) RemoveKeywords() <-chan []*Keyword {
	return this.RemoveKeywordsChan
}

func (this *Listener) AddSymbols() <-chan []*Keyword {
	return this.AddSymbolsChan
}

func (this *Listener) RemoveSymbols() <-chan []*Keyword {
	return this.RemoveSymbolsChan
}
