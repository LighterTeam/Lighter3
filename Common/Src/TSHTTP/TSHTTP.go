package TSHTTP

import (
	"net/http"
	"log"
)

type ReceiveBuffer func(w http.ResponseWriter, r *http.Request)

func CreateHTTPServer (sWebPath string, rb ReceiveBuffer) {
	http.HandleFunc("/", rb)
	log.Fatal(http.ListenAndServe(sWebPath, nil));
}
