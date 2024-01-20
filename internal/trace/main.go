package trace

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func DisplayTrace(writer http.ResponseWriter, request *http.Request) {
	listHeader := []string{"Accept", "Accept-Encoding", "Accept-Language", "User-Agent"}
	var builder strings.Builder
	builder.WriteString("ip: " + request.Header.Get("X-Forwarded-For") + "\n")
	builder.WriteString("proxy: " + request.Header.Get("X-Forwarded-Server") + "\n")
	builder.WriteString("host: " + request.Header.Get("X-Forwarded-Host") + "\n")
	for _, value := range listHeader {
		builder.WriteString(value + ": " + request.Header.Get(value) + "\n")
	}

	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := io.WriteString(writer, builder.String())
	if err != nil {
		log.Println(err)
	}
}
