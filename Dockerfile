FROM hub.skyinno.com/common/alpine:latest
MAINTAINER FAE Config Server "fae@fiberhome.com"
ADD caddy /usr/local/bin/
ADD Caddyfile.tmpl /etc/caddy/

ENV CONSUL=http://127.0.0.1:8500
ENV CADDYFILE_PATH=/etc/caddy/Caddyfile.tmpl

ENTRYPOINT ["/usr/local/bin/caddy"]