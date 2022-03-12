ARG ALPINE_VERSION=3.14
ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as builder

WORKDIR /go_app

COPY . .

ENV GO111MODULE=on

ENV PORT=9001

RUN go mod download

RUN CGO_ENABLED=0 go build -a -o bin/server_main ./server/server_main.go

FROM alpine:${ALPINE_VERSION}

WORKDIR /home/app

COPY --from=builder /go_app/bin/server_main .

EXPOSE 9001

CMD ["./server_main"]
