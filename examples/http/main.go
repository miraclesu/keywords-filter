package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/miraclesu/keywords-filter"
	"github.com/miraclesu/keywords-filter/listener/http.listen"
	"github.com/miraclesu/keywords-filter/loader/http.load"
)

var (
	Port      = flag.String("p", ":7520", "serve's port")
	Threshold = flag.Int("t", 100, "Threshold of filter")

	Filter *filter.Filter
)

func main() {
	flag.Parse()

	var err error
	Filter, err = filter.New(*Threshold, &load.Loader{})
	if err != nil {
		log.Println(err.Error())
		return
	}

	Filter.StartListen(listen.DefaultListener)

	router := httprouter.New()
	router.POST("/filter", filterHandler)
	router.POST("/addkws", listen.AddKeywords)
	router.POST("/addsbs", listen.AddSymbols)
	router.POST("/rmkws", listen.RemoveKeywords)
	router.POST("/rmsbs", listen.RemoveSymbols)
	log.Println("serve listen on", *Port)
	http.ListenAndServe(*Port, router)
}

func filterHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	req, resp := new(filter.Request), new(listen.Response)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Msg = err.Error()
		data, _ := json.Marshal(resp)
		w.Write(data)
		return
	}

	req.Init(Filter)
	resp.Result, resp.Success = req.Scan(), true
	data, _ := json.Marshal(resp)
	w.Write(data)
}
