FROM golang:1.12.4

ARG app=client

COPY . /go/src/sandwiches
WORKDIR /go/src/sandwiches/${app}

RUN go get -v
RUN go build -v

CMD "${app}"
