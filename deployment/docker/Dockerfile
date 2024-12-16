ARG GO_VERSION=1.23.4-alpine3.20

FROM golang:${GO_VERSION}

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build

EXPOSE 8080

ENTRYPOINT ["/app/simple-golang-api"]

