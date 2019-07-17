package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func StartServer(){
	r:=Routers()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
	}
	srv.ListenAndServe()
}

func Routers() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {


	response, _ := http.Get("https://google.com/")
	fmt.Println(response)
	body, _ := ioutil.ReadAll(response.Body)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}