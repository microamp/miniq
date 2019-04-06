FROM golang:1.12-alpine

ENV GOPATH /go
ENV GOBIN $GOPATH/bin

WORKDIR $GOPATH/src/github.com/microamp/miniq

RUN apk add --no-cache git && go get github.com/golang/dep/cmd/dep

COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure -v -vendor-only

COPY . .

CMD ["go", "run", "examples/github-trending/main.go"]
