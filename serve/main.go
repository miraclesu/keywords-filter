package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/miraclesu/keywords-filter"
	"github.com/miraclesu/keywords-filter/listener/http.listen"
	"github.com/miraclesu/keywords-filter/loader/mgo.load"
)

var (
	Port = flag.String("p", ":7520", "serve's port")

	LoadConf   = flag.String("load", "load.json", "loader config file")
	ListenConf = flag.String("listen", "listen.json", "listener config file")

	Threshold = flag.Int("t", 100, "Threshold of filter")

	Filter *filter.Filter
)

func main() {
	flag.Parse()

	loader, err := load.New(*LoadConf)
	if err != nil {
		log.Printf("new loader err: %s\n", err.Error())
		return
	}

	Filter, err = filter.New(*Threshold, loader)
	if err != nil {
		log.Printf("new filter err: %s\n", err.Error())
		return
	}
	loader.Session.Close()

	Filter.StartListen(listen.DefaultListener)

	router := httprouter.New()
	router.POST("/filter", filterHandler)
	router.POST("/addkws", listen.AddKeywords)
	router.POST("/addsbs", listen.AddSymbols)
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
