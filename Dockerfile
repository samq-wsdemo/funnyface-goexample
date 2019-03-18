FROM golang:1.12

ENV GO111MODULE=on
WORKDIR /go/src/github.com/funnyface-inc/go-example
COPY . .
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/tockins/realize
CMD ["realize", "start", "--run"]
