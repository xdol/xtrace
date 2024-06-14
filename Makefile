build: front backend

lint: backend-lint

backend:
	go build -ldflags "-s -w" .

.PHONY: front
front:
	mkdir static/
	sass --no-source-map --style=compressed front/style.scss static/style.css

clean:
	rm -r static/
	rm xtrace

backend-lint:
	golangci-lint run .