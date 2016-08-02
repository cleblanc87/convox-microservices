FROM golang:1.6.2-alpine

RUN apk update
RUN apk add git

ENV GOPATH /go
ENV PATH $GOPATH/BIN:$PATH
ENV GO15VENDOREXPERIMENT 1

RUN go get github.com/ddollar/init
RUN go get github.com/cleblanc87/rerun

WORKDIR /go/src/github.com/cleblanc87/convox-microservices/api1/app
COPY . /go/src/github.com/cleblanc87/convox-microservices/api1/app
RUN go install ./...

ENTRYPOINT ["/go/bin/init"]
CMD ["bin/web"]
