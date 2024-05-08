package templates

const (
	GatewayDockerfile = `
FROM nginx:latest

COPY deploy/gateway/default.conf /etc/nginx/conf.d/
`
	GatewayConf = `
server {
    listen 80;

    resolver 127.0.0.11 valid=30s;

	# GOPHER: Don't remove this line
	# {{ .command }}
}
`
	GatewayCommand = `
	# {{ .command }}
	location /{{ .name }} {
        set $upstream {{ .name }};
        proxy_pass http://$upstream:8080;
    }
`
)
