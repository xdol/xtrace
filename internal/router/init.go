package router

import (
	"log"
	"net/http"

	"git.xdol.org/xdol/xtrace/internal/constant"
	"git.xdol.org/xdol/xtrace/internal/router/trace"
)

func (c Config) Server() {
	//fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /trace", trace.Get)
	mux.HandleFunc("GET /errors/{code}", c.Error.Get)
	http.Handle("/static/", http.FileServer(http.Dir("./static")))

	server := &http.Server{
		Addr:        c.ListenAddress,
		ReadTimeout: constant.ReadTimeout,
		Handler:     mux,
	}

	log.Println("Listen on", c.ListenAddress)

	log.Fatal(server.ListenAndServe())
}
