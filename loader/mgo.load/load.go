// load the keywords from mongodb
package load

import (
	"encoding/json"
	"os"

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

	loader.Session, err = mgo.DialWithInfo(loader.DialInfo)
	return
}

func (this *Loader) Load() (kws []*keyword.Keyword, sbs []*keyword.Keyword, err error) {
	db := this.DB(this.Database)
	if err = db.C(this.KeywordsColl).Find(bson.M{}).All(&kws); err != nil {
		return
	}

	err = db.C(this.SymbolsColl).Find(bson.M{}).All(&sbs)
	return
}
