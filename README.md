# GoForward

Make requests to locally hosted backend services without worrying about CORS. Use a single endpoint to route requests to different ports

```
http://localhost:1338/{MICROSERVICE_PORT}/}{PATH}
```

For example, if you start this service on 1338, you can query `http://localhost:1338/8090/v1/chat/channels` which will route the request to `http://localhost:8090/v1/chat/channels`

By changing the fist Path variable, you can point to any port on your system.

## Build & Run

1) Install Go
2) Set `$GOPATH`
3) Clone this repo inside `$GOPATH/src/github.com`
4) Run `go build main.go && ./main` which spits out an executable file that you can use forever

You can ovverride the default port by setting an OS Env Variable `PORT` with your override

## BONUS 🔥 - Free CORS Headers

We add the following headers to all your requests

```
	"Access-Control-Allow-Origin", "*"
	"Access-Control-Allow-Methods", "GET, POST, OPTIONS"
	"Access-Control-Allow-Headers", "*"
	"Access-Control-Max-Age", "1728000"
```

*NOTE*: All `OPTIONS` type requests i.e. pre-flight requests are not reverse-proxied and are instead handled by this service and returns with a `204`.