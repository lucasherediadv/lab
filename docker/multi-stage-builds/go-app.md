# Tiny go-application container

The image has both the compiler and the compiled binary, which is too much. We only need the binary to run our application:

go.mod:

```
module example.com/mod
```

hello.go:

```
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

Dockerfile:

```
FROM golang:1.19 as builder
WORKDIR /app
COPY . /app
RUN go mod download && go mod verify
RUN cd /app && go build -o goapp
ENTRYPOINT ./goapp
```

By utilizing multi-stage build, we can separate the build stage (compiling) from the image we actually want to ship. See the app/ folder.

