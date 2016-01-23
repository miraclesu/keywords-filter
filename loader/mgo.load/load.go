// load the keywords from mongodb
package load

import (
	"encoding/json"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/miraclesu/keywords-filter/keyword"
)

type Loader struct {
	KeywordsColl string
	SymbolsColl  string
	*mgo.DialInfo
	*mgo.Session
}

func New(path string) (loader *Loader, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	loader = &Loader{
		DialInfo: &mgo.DialInfo{},
	}
	if err = json.NewDecoder(file).Decode(loader); err != nil {
		return
	}
	loader.Timeout *= time.Second

	loader.Session, err = mgo.DialWithInfo(loader.DialInfo)
	return
}

func (this *Loader) Load() (kws []*keyword.Keyword, sbs []string, err error) {
	db := this.DB(this.Database)
	if err = db.C(this.KeywordsColl).Find(bson.M{}).All(&kws); err != nil {
		return
	}

	sbws := make([]keyword.Keyword, 0)
	if err = db.C(this.SymbolsColl).Find(bson.M{}).All(&sbws); err != nil {
		return
	}

	count := len(sbws)
	if count == 0 {
		return
	}

	sbs = make([]string, 0, count)
	for i := 0; i < count; i++ {
		if len(sbws[i].Word) == 0 {
			continue
		}

		sbs = append(sbs, sbws[i].Word)
	}
	return
}
