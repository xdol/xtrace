package errors

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Code    int
	Title   string
	Message string
}

func (c Config) Get(w http.ResponseWriter, r *http.Request) {
	code, err := strconv.Atoi(r.PathValue("code"))
	if err != nil {
		log.Println(err)
	}
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Println(err)
	}

	s, err := findResponse(c.Errors, code)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(code)
	err = t.Execute(w, s)
	if err != nil {
		log.Println(err)
	}
}

func findResponse(c []Response, code int) (Response, error) {
	for _, v := range c {
		if v.Code == code {
			return v, nil
		}
	}

	return Response{}, errors.New("error" + strconv.FormatInt(int64(code), 10) + "not found")
}
