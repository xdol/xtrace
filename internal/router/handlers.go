package router

import (
	"log"
	"net/http"

	"git.xdol.org/xdol/xtrace/internal/constant"
	"git.xdol.org/xdol/xtrace/internal/trace"
)

func HTTPHandle(listenAddress string) {
	http.HandleFunc("/trace", trace.DisplayTrace)
	log.Println("Listen on", listenAddress)

	server := &http.Server{
		Addr:        listenAddress,
		ReadTimeout: constant.ReadTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
