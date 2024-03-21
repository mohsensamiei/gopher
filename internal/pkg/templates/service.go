package templates

const (
	ServiceDockerfile = `
FROM golang:alpine as builder
RUN apk add --update --no-cache git

ARG VERSION
ARG GOPROXY

WORKDIR /src
COPY go.mod go.sum ./
COPY vendor* vendor
RUN go mod tidy
COPY . .

# GOPHER: Don't remove this line
# {{ .command }}

FROM alpine:latest
RUN apk add --update --no-cache ca-certificates tzdata mailcap

WORKDIR /app
COPY --from=builder /src/build/ ./
COPY ./assets ./assets
`

	ServiceBuild = `
# {{ .command }}
RUN GO111MODULE=on CGO_ENABLED=0 go build -buildvcs=false -a -installsuffix cgo \
    -ldflags "-w -X main.Version=${VERSION}" \
    -o ./build/{{ .name }} ./cmd/{{ .name }}
`
)
