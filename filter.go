package filter

import (
	"errors"

	. "github.com/miraclesu/keywords-filter/keyword"
	"github.com/miraclesu/keywords-filter/loader"
)

var (
	InvalidFilter = errors.New("Invalid filter")
)

type Filter struct {
	word      *Word
	Threshold int16
}

func New(threshold int16, load loader.Loader) (f *Filter, err error) {
	kws, err := load.Load()
	if err != nil {
		return
	}

	f = &Filter{
		word:      new(Word),
		Threshold: threshold,
	}

	f.AddWords(kws)
	return
}

func (this *Filter) Filter(content string) (b bool, err error) {
	if this.word == nil {
		err = InvalidFilter
		return
	}
	return NewResponse(this, content).scan(), nil
}

func (this *Filter) AddWord(w *Keyword) {
	if this.word == nil {
		this.word = new(Word)
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
