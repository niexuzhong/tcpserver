FROM golang:latest

WORKDIR $GOPATH/src/tcpserver
COPY . $GOPATH/src/tcpserver
RUN go build .

EXPOSE 9001
ENTRYPOINT ["./tcpserver"]
