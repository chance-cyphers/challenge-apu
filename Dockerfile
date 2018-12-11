FROM golang:latest
RUN mkdir $GOPATH/src/app
ADD . $GOPATH/src/app/
ADD proto $GOPATH/src/app/proto
WORKDIR $GOPATH/src/app

RUN go get -u github.com/kardianos/govendor
RUN go get -u golang.org/x/sys/unix
RUN govendor sync
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
