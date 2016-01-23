package listen

import (
	"fmt"
	"log"

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

	go l.Subscribe(conf.KwsChannel, conf.SbsChannel)
	return
}

func (this *Listener) Subscribe(kws, sbs string) {
	psc := redis.PubSubConn{Conn: this.conn}
	psc.Subscribe(kws)
	psc.Subscribe(sbs)

	kind := 0
	for {
		switch n := psc.Receive().(type) {
		case redis.Message:
			if n.Channel == kws {
				kind = KW
			} else {
				kind = SB
			}
			if err := this.Send(kind, n.Data); err != nil {
				log.Printf("invalid notify, channel:[%s], data:[%s]\n", n.Channel, n.Data)
			}
		default:
			continue
		}
	}
}

func (this *Listener) Send(kind int, data []byte) (err error) {
	notify, err := NewNotify(kind, data)
	if err != nil {
		return err
	}

	switch notify.Action {
	case ADD:
		if kind == KW {
			this.AddKeywordsChan <- *notify.Data.(*[]*Keyword)
		} else {
			this.AddSymbolsChan <- *notify.Data.(*[]string)
		}
	case RM:
		if kind == KW {
			this.RemoveKeywordsChan <- *notify.Data.(*[]*Keyword)
		} else {
			this.RemoveSymbolsChan <- *notify.Data.(*[]string)
		}
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
