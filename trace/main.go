package trace

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func DisplayTrace(w http.ResponseWriter, req *http.Request) {
	listHeader := []string{"Accept", "Accept-Encoding", "Accept-Language", "User-Agent"}
	var builder strings.Builder
	builder.WriteString("ip: " + req.Header.Get("X-Forwarded-For") + "\n")
	builder.WriteString("proxy: " + req.Header.Get("X-Forwarded-Server") + "\n")
	builder.WriteString("host: " + req.Header.Get("X-Forwarded-Host") + "\n")
	for _, value := range listHeader {
		builder.WriteString(value + ": " + req.Header.Get(value) + "\n")
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := io.WriteString(w, builder.String())
	if err != nil {
		log.Println(err)
	}
}
