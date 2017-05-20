FROM golang:1.8

ADD . /go/src
WORKDIR /go/src

RUN go get github.com/bitly/go-simplejson

EXPOSE 8080

CMD ["go", "run", "server.go"]
