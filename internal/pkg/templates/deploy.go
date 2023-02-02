package templates

const (
	DeployBuild = `
# GOPHER: Don't remove this line
# registry {{ .registry }}
version: "3.9"

services:
  gateway:
    container_name: gateway
    image: {{ .registry }}/gateway:${VERSION:-latest}
    build:
      context: ..
      dockerfile: services/gateway/Dockerfile
      network: host
# GOPHER: Don't remove this line
# {{ .service }}
`

	DeployBuildService = `
# {{ .service }}
  {{ .name }}:
    container_name: {{ .name }}
    image: {{ .registry }}/{{ .name }}:${VERSION:-latest}
    build:
      context: ..
      dockerfile: services/{{ .name }}/Dockerfile
      args:
        - VERSION=${VERSION:-latest}
        - GOPROXY=${GOPROXY}
      network: host
`

	DeployUp = `
# GOPHER: Don't remove this line
# registry {{ .registry }}
version: "3.9"

services:
  gateway:
    container_name: gateway
    image: {{ .registry }}/gateway:${VERSION:-latest}
    ports:
      - "8080:80"
    restart: on-failure
# GOPHER: Don't remove this line
# {{ .command }}
`
	DeployRunService = `
# {{ .command }}
  {{ .name }}:
    container_name: {{ .name }}
    image: {{ .registry }}/{{ .service }}:${VERSION:-latest}
    entrypoint: ./{{ .name }}
    env_file:
      - ${ENV_FILE}
    restart: on-failure
`
)
