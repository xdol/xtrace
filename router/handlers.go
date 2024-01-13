package router

import (
	"git.xdol.org/xdol/xtrace/trace"
	"log"
	"net/http"
)

func HTTPHandle(listenAddress string) {
	http.HandleFunc("/trace", trace.DisplayTrace)
	log.Println("Listen on", listenAddress)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
