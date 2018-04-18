# Consul Plugin for Caddy

*Forked from https://github.com/krishamoud/caddy-consul*

## Installation

### 1. Get the source
```
go get github.com/chenjpu/caddy-consul
```

### 2. Get the Caddy source
```
go get github.com/mholt/caddy/caddy
```

### 3. Add the plugin
```
// github.com/mholt/caddy/caddy/caddymain/run.go
package caddymain

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/xenolf/lego/acme"

	"github.com/mholt/caddy"
	// plug in the HTTP server type
	_ "github.com/mholt/caddy/caddyhttp"

	"github.com/mholt/caddy/caddytls"
	// This is where other plugins get plugged in (imported)
	_ "github.com/chenjpu/caddy-consul" // ADD THIS LINE
)
```

### 4. Build Caddy
```
cd github.com/mholt/caddy/caddy
go get github.com/caddyserver/builds
go run build.go
```

### 5. Set CONSUL ENV variable
This tells consul where to connect
```
export CONSUL=http://10.0.0.1:8500
```

### 6. Set CADDYFILE_PATH env
This tells caddy where to look for the Caddyfile.tmpl template
```
export CADDYFILE_PATH=/path/to/Caddyfile.tmpl
```

### 7. Run Caddy
Either add your newly built binary to your `$PATH` or just run `./caddy` from where you built the application.

## Example Caddyfile.tmpl
```
{{range $domain, $services := .}} http://{{$domain}}.xxx.com {
  root /var/www/dev.xxx.com
  gzip
  proxy /services {
     header_upstream Host {host}
     header_upstream X-Real-IP {remote}
     header_upstream X-Forwarded-For {remote}
     header_upstream X-Forwarded-Proto {scheme}
     max_conns 200
     policy ip_hash
    {{range $key, $service := $services}}upstream {{.Address}}:{{.ServicePort}}
    {{end}}
  }
  rewrite / {
   if {uri} not_starts_with /services
   to {uri} /
  }
}
{{end}}
```

The above will output something like this:

```
 http://test-dev.xxx.com {
  root /var/www/dev.xxx.com
  gzip
  proxy /services {
     header_upstream Host {host}
     header_upstream X-Real-IP {remote}
     header_upstream X-Forwarded-For {remote}
     header_upstream X-Forwarded-Proto {scheme}
     max_conns 200
     policy ip_hash
    upstream 10.110.200.26:23530

  }
  rewrite / {
   if {uri} not_starts_with /services
   to {uri} /
  }
}
```

## TODOS:
1. Tests
2. Refactoring
