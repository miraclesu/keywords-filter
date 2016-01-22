package listen

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/miraclesu/keywords-filter"
	. "github.com/miraclesu/keywords-filter/keyword"
)

type Response struct {
	Success bool
	Msg     string
	Result  *filter.Response `json:",omitempty"`
}

func AddKeywords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	kws, resp := make([]*Keyword, 0), new(Response)
	err := json.NewDecoder(r.Body).Decode(&kws)
	if err != nil {
		resp.Msg = err.Error()
		data, _ := json.Marshal(resp)
		w.Write(data)
		return
	}
	DefaultListener.AddKeywordsChan <- kws
	resp.Success = true
	data, _ := json.Marshal(resp)
	w.Write(data)
}

func AddSymbols(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	sbs, resp := make([]string, 0), new(Response)
	err := json.NewDecoder(r.Body).Decode(&sbs)
	if err != nil {
		resp.Msg = err.Error()
		data, _ := json.Marshal(resp)
		w.Write(data)
		return
	}
	DefaultListener.AddSymbolsChan <- sbs
	resp.Success = true
	data, _ := json.Marshal(resp)
	w.Write(data)
}
