FROM golang:1.20.0

WORKDIR /usr/src/app

EXPOSE 8080

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy

ENTRYPOINT ["go", "run", "./src/main.go"]

