package listen

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	. "github.com/miraclesu/keywords-filter/keyword"
)

func AddKeywords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	kws := make([]*Keyword, 0)
	err := json.NewDecoder(r.Body).Decode(&kws)
	if err != nil {
		return
	}
	DefaultListener.AddKeywordsChan <- kws
}
