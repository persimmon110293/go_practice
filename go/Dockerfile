FROM golang:1.21-alpine

RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY ./go/src /go/src/app/

RUN apk update && apk add git

RUN go mod tidy \
  && go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

EXPOSE 8080

CMD ["go", "run", "server.go"]