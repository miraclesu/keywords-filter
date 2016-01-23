package listen

import (
	"encoding/json"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Conf struct {
	NetWork string
	Address string
	Timeout int

	KwsChannel string
	SbsChannel string
}

func NewConf(path string) (conf *Conf, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	conf = new(Conf)
	err = json.NewDecoder(file).Decode(conf)
	return
}

func (this *Conf) Dail() (conn redis.Conn, err error) {
	option := redis.DialConnectTimeout(time.Duration(this.Timeout) * time.Second)
	conn, err = redis.Dial(this.NetWork, this.Address, option)
	return
}
