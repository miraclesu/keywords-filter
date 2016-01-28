package filter

import (
	"sync"

	"github.com/miraclesu/keywords-filter/keyword"
	"github.com/miraclesu/keywords-filter/listener"
	"github.com/miraclesu/keywords-filter/loader"
)

// Filter is used to build keyword trees
type Filter struct {
	word      *Word    // keyword
	symb      *symbols // special symbols
	Threshold int
}

// New a Filter
func New(threshold int, load loader.Loader) (f *Filter, err error) {
	kws, sbs, err := load.Load()
	if err != nil {
		return
	}

	f = &Filter{
		word:      &Word{lk: new(sync.RWMutex)},
		symb:      &symbols{lk: new(sync.RWMutex)},
		Threshold: threshold,
	}

	f.AddWords(kws)
	f.AddSymbs(sbs)
	return
}

// Filter the content
func (f *Filter) Filter(content string) *Response {
	req := Request{
		Content: content,
	}
	req.Init(f)
	return req.Scan()
}

// AddWord add a keyword to the filter
func (f *Filter) AddWord(w *keyword.Keyword) {
	if f.word == nil {
		f.word = &Word{
			lk: new(sync.RWMutex),
		}
	}
	f.word.addWord(w)
}

// RemoveWord remove a keyword from the filter
func (f *Filter) RemoveWord(w *keyword.Keyword) {
	if f.word == nil {
		return
	}
	f.word.removeWord(w)
}

// AddWords add keywords to the filter
func (f *Filter) AddWords(kws []*keyword.Keyword) {
	for i, count := 0, len(kws); i < count; i++ {
		f.AddWord(kws[i])
	}
}

// RemoveWords remove keywords from the filter
func (f *Filter) RemoveWords(kws []*keyword.Keyword) {
	for i, count := 0, len(kws); i < count; i++ {
		f.RemoveWord(kws[i])
	}
}

// AddSymb add a symbol to the filter
func (f *Filter) AddSymb(s string) {
	if f.symb == nil {
		f.symb = &symbols{
			lk: new(sync.RWMutex),
		}
	}
	for _, v := range s {
		f.symb.add(v)
	}
}

// RemoveSymb remove a symbol from the filter
func (f *Filter) RemoveSymb(s string) {
	if f.word == nil {
		return
	}
	for _, v := range s {
		f.symb.remove(v)
	}
}

// AddSymbs add symbols to the filter
func (f *Filter) AddSymbs(sbs []string) {
	for i, count := 0, len(sbs); i < count; i++ {
		f.AddSymb(sbs[i])
	}
}

// RemoveSymbs remove symbols from the filter
func (f *Filter) RemoveSymbs(sbs []string) {
	for i, count := 0, len(sbs); i < count; i++ {
		f.AddSymb(sbs[i])
	}
}

// StartListen start a listener to listen add or remove keywords and symbols
func (f *Filter) StartListen(listen listener.Listener) {
	go func() {
		for kws := range listen.AddKeywords() {
			f.AddWords(kws)
		}
	}()
	go func() {
		for kws := range listen.RemoveKeywords() {
			f.RemoveWords(kws)
		}
	}()
	go func() {
		for kws := range listen.AddSymbols() {
			f.AddSymbs(kws)
		}
	}()
	go func() {
		for kws := range listen.RemoveSymbols() {
			f.RemoveSymbs(kws)
		}
	}()
}
