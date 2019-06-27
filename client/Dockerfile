FROM golang:1.12.4
WORKDIR /go/src/app
COPY . .
RUN go get github.com/calvinmclean/sandwiches || true
RUN cd $GOPATH/src/github.com/calvinmclean/sandwiches && git checkout gRPC
RUN cd /go/src/app
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
