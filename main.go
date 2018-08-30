package main

import (
	"net/http"

	"github.com/xuyz/stock/scron"

	"github.com/xuyz/stock/sconfig"

	"github.com/gorilla/mux"
)

func main() {
	scron.Start()

	r := mux.NewRouter()
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("I'm a superman!"))
	})

	http.ListenAndServe(":"+sconfig.HttpPort(), r)
}
