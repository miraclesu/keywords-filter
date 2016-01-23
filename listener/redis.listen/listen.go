package listen

import (
	"fmt"

	"github.com/garyburd/redigo/redis"

	. "github.com/miraclesu/keywords-filter/keyword"
)

type Listener struct {
	conn redis.Conn

	AddKeywordsChan    chan []*Keyword
	RemoveKeywordsChan chan []*Keyword
	AddSymbolsChan     chan []string
	RemoveSymbolsChan  chan []string
}

func New(path string) (l *Listener, err error) {
	conf, err := NewConf(path)
	if err != nil {
		return
	}

	conn, err := conf.Dail()
	if err != nil {
		return
	}

	l = &Listener{
		conn: conn,

		AddKeywordsChan:    make(chan []*Keyword, 8),
		RemoveKeywordsChan: make(chan []*Keyword, 8),
		AddSymbolsChan:     make(chan []string, 8),
		RemoveSymbolsChan:  make(chan []string, 8),
	}

	go l.Subscribe(conf.Channel)
	return
}

func (this *Listener) Subscribe(channel string) {
	psc := redis.PubSubConn{Conn: this.conn}
	psc.Subscribe(channel)

	for {
		switch n := psc.Receive().(type) {
		case redis.Message:
			_ = this.Send(n.Data)
		default:
			continue
		}
	}
}

func (this *Listener) Send(data []byte) error {
	notify, err := NewNotify(data)
	if err != nil {
		return err
	}

	switch notify.Action {
	case AddKws:
		kws, ok := notify.Data.([]*Keyword)
		fmt.Printf("ok=%v, kws=%+v\n", kws, ok)
	case RmKws:
	case AddSbs:
	case RmSbs:
	default:
		err = fmt.Errorf("invalid notify action:[%d]", notify.Action)
	}

	return err
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
