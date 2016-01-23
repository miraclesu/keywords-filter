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
	kws := make([]*Keyword, 0)
	if err := Do(w, r, &kws); err != nil {
		return
	}
	DefaultListener.AddKeywordsChan <- kws
}

func RemoveKeywords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	kws := make([]*Keyword, 0)
	if err := Do(w, r, &kws); err != nil {
		return
	}
	DefaultListener.RemoveKeywordsChan <- kws
}

func AddSymbols(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sbs := make([]string, 0)
	if err := Do(w, r, &sbs); err != nil {
		return
	}
	DefaultListener.AddSymbolsChan <- sbs
}

func RemoveSymbols(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sbs := make([]string, 0)
	if err := Do(w, r, &sbs); err != nil {
		return
	}
	DefaultListener.RemoveSymbolsChan <- sbs
}

func Do(w http.ResponseWriter, r *http.Request, body interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json")
	resp := new(Response)
	err = json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		resp.Msg = err.Error()
		data, _ := json.Marshal(resp)
		w.Write(data)
		return
	}

	resp.Success = true
	data, _ := json.Marshal(resp)
	w.Write(data)
	return
}
