ARG GO_VERSION=1.19.1

FROM golang:${GO_VERSION}-alpine3.16 as builder

ENV GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates
RUN apk update

WORKDIR /go/src

COPY ["./go.mod", "./go.sum","./vendor/", "./"]

RUN go mod download

COPY [".", "."]

RUN CGO_ENABLED=0 go build \
    -o /go/bin/rest-go-ws

CMD [ "/go/bin/rest-go-ws" ]