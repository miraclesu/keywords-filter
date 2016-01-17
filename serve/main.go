package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	Filter.StartListen(listen.NewListener())

	r := mux.NewRouter()
	r.HandleFunc("/filter", filterHandler).
		Methods("POST")
	log.Println("serve listen on", *Port)
	http.ListenAndServe(*Port, r)
}

func filterHandler(w http.ResponseWriter, r *http.Request) {

}
