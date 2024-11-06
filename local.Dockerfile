FROM golang:1.23-alpine3.20 AS builder

WORKDIR /github.com/Chigvero/source

COPY . /github.com/Chigvero/source

RUN go mod tidy
RUN go mod download
RUN go build -o chatService cmd/chat-server/main.go

FROM alpine:3.20

WORKDIR /root/

COPY --from=builder /github.com/Chigvero/source/chatService /root
COPY --from=builder /github.com/Chigvero/source/local.env .



ENTRYPOINT ["./chatService","-config","local.env"]