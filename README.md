# unstable
Make unstable http/gRPC server

## What's unstable
I think there are multiple types of instability.

This package support three unstable.
1. Return error etc. stochastically
1. Slow response
1. Return server error(http: Status Code 500, gRPC: codes.Internal)

## Configure
You can configure "unstable" by `unstable.json`.
Like, 
```
{
  "interval": 1,
  "slow_response_option": {
    "enable": true,
    "time": 500
  },
  "server_error_option": {
    "enable": true
  }
}
```

- `interval`(seconds): When the unix time(sencods) is be divisable by `interval`, then return `slow response` or `server error`. When `0`, then disable this package.
- `slow_response_option`: Control repsonse time.
  - `enable`: When true, then response time will take more than `time`
  - `time`(milliseconds): Response time will take more than `time`
- `server_error_option`: Control server error.
  - `enable`: When true, then return server error
  
Pseudocode
```
if interval != 0 && UnixTime%interval {
  if slow_response_option.enable {
    sleep(slow_response_option.time)
  }
  if server_error_option.enable {
    raise_server_error()
  }
}
```

## Installation
This package support only `http` and `gRPC` server.
See example https://github.com/rerost/unstable/tree/master/example .

### http
```
$ go get github.com/rerost/unstable/uhttp
```

`server.go`

```
package main

import (
	"fmt"
	"net/http"

	"github.com/rerost/unstable/uhttp"
)

func main() {
	http.Handle("/", uhttp.WithUnstable(handler))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sample"))
}
```

### grpc
```
$ go get github.com/rerost/unstable/ugrpc
```

`server.go`

```
package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	api_pb "github.com/rerost/unstable/example/grpc/server/api"
	"github.com/rerost/unstable/ugrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetSample(ctx context.Context, req *empty.Empty) (*api_pb.GetSampleResponse, error) {
	return &api_pb.GetSampleResponse{Message: "Sample"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(ugrpc.UnstableUnaryServerInterceptor()),
	)
	api_pb.RegisterSampleServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```
