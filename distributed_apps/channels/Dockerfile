FROM golang:1.16-alpine as builder

RUN mkdir -p /http_json_rest/

WORKDIR /http_json_rest

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -a -o bin/server_main ./server/server_main.go

RUN CGO_ENABLED=0 go build -a -o bin/client_main ./client/client_main.go

FROM alpine:3.14

WORKDIR /home/app

COPY --from=builder /http_json_rest/bin/server_main .

CMD ["./server_main"]
