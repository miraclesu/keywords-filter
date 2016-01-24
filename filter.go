package filter

import (
	"errors"
	"sync"

	. "github.com/miraclesu/keywords-filter/keyword"
	"github.com/miraclesu/keywords-filter/listener"
	"github.com/miraclesu/keywords-filter/loader"
)

var (
	InvalidFilter = errors.New("Invalid filter")
)

type Filter struct {
	word      *Word    // keyword
	symb      *symbols // special symbols
	Threshold int
}

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

func (this *Filter) Filter(content string) (resp *Response, err error) {
	if this.word == nil || this.symb == nil {
		err = InvalidFilter
		return
	}

	req := Request{
		Content: content,
	}
	req.Init(this)
	return req.Scan(), nil
}

func (this *Filter) AddWord(w *Keyword) {
	if this.word == nil {
		this.word = &Word{
			lk: new(sync.RWMutex),
		}
	}
	this.word.addWord(w)
}

func (this *Filter) RemoveWord(w *Keyword) {
	if this.word == nil {
		return
	}
	this.word.removeWord(w)
}

func (this *Filter) AddWords(kws []*Keyword) {
	for i, count := 0, len(kws); i < count; i++ {
		this.AddWord(kws[i])
	}
}

func (this *Filter) RemoveWords(kws []*Keyword) {
	for i, count := 0, len(kws); i < count; i++ {
		this.RemoveWord(kws[i])
	}
}

func (this *Filter) AddSymb(s string) {
	if this.symb == nil {
		this.symb = &symbols{
			lk: new(sync.RWMutex),
		}
	}
	for _, v := range s {
		this.symb.add(v)
	}
}

func (this *Filter) RemoveSymb(s string) {
	if this.word == nil {
		return
	}
	for _, v := range s {
		this.symb.remove(v)
	}
}

func (this *Filter) AddSymbs(sbs []string) {
	for i, count := 0, len(sbs); i < count; i++ {
		this.AddSymb(sbs[i])
	}
}

func (this *Filter) RemoveSymbs(sbs []string) {
	for i, count := 0, len(sbs); i < count; i++ {
		this.AddSymb(sbs[i])
	}
}

func (this *Filter) StartListen(listen listener.Listener) {
	go func() {
		for kws := range listen.AddKeywords() {
			this.AddWords(kws)
		}
	}()
	go func() {
		for kws := range listen.RemoveKeywords() {
			this.RemoveWords(kws)
		}
	}()
	go func() {
		for kws := range listen.AddSymbols() {
			this.AddSymbs(kws)
		}
	}()
	go func() {
		for kws := range listen.RemoveSymbols() {
			this.RemoveSymbs(kws)
		}
	}()
}
