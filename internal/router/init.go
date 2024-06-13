package router

import (
	"log"
	"net/http"

	"git.xdol.org/xdol/xtrace/internal/constant"
	"git.xdol.org/xdol/xtrace/internal/router/trace"
)

func (c Config) Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /trace", trace.Get)
	log.Println("Listen on", c.ListenAddress)

	server := &http.Server{
		Addr:        c.ListenAddress,
		ReadTimeout: constant.ReadTimeout,
		Handler:     mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
