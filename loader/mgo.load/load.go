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
	Collection string
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

func (this *Loader) Load() (kws []*keyword.Keyword, err error) {
	c := this.DB(this.Database).C(this.Collection)
	err = c.Find(bson.M{}).All(&kws)
	return
}
