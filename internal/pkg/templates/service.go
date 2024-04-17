package templates

const (
	ServiceDockerfile = `
FROM ghcr.io/mohsensamiei/gopher:builder-latest as builder

ARG VERSION
ARG GOPROXY

WORKDIR /src
COPY go.mod go.sum ./
COPY vendor* vendor
RUN gopher dep
COPY . .
RUN gopher proto

# GOPHER: Don't remove this line
# {{ .command }}

FROM ghcr.io/mohsensamiei/gopher:server-latest

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
