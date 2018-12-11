FROM golang:latest
RUN mkdir $GOPATH/src/challenge
ADD . $GOPATH/src/challenge/
WORKDIR $GOPATH/src/challenge

RUN go get -u github.com/kardianos/govendor
RUN go get -u golang.org/x/sys/unix
RUN govendor sync
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
