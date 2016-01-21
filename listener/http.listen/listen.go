package listen

import . "github.com/miraclesu/keywords-filter/keyword"

var (
	DefaultListener = NewListener()
)

type Listener struct {
	AddKeywordsChan    chan []*Keyword
	RemoveKeywordsChan chan []*Keyword
	AddSymbolsChan     chan []string
	RemoveSymbolsChan  chan []string
}

func NewListener() *Listener {
	l := &Listener{
		AddKeywordsChan:    make(chan []*Keyword, 8),
		RemoveKeywordsChan: make(chan []*Keyword, 8),
		AddSymbolsChan:     make(chan []string, 8),
		RemoveSymbolsChan:  make(chan []string, 8),
	}
	return l
}

func (this *Listener) AddKeywords() <-chan []*Keyword {
	return this.AddKeywordsChan
}

func (this *Listener) RemoveKeywords() <-chan []*Keyword {
	return this.RemoveKeywordsChan
}

func (this *Listener) AddSymbols() <-chan []string {
	return this.AddSymbolsChan
}

func (this *Listener) RemoveSymbols() <-chan []string {
	return this.RemoveSymbolsChan
}
