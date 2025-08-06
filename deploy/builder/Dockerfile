FROM golang:alpine as builder

RUN apk add --update --no-cache git protoc build-base

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/favadi/protoc-go-inject-tag@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
RUN go install github.com/mohsensamiei/gopher/v3/cmd/gopher@latest