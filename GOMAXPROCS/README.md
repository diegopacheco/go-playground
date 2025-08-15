# GOMAXPROCS

GOMAXPROCS is an environment variable in Go that determines the maximum number of CPU cores that can be executing simultaneously. By default, Go sets this value to the number of available CPU cores, but it can be adjusted to limit the number of cores used by a Go program.

## Run

```bash
./build.sh
./run.sh
```

## Result

```bash
[+] Building 5.8s (20/20) FINISHED                                                                                                                        docker:default
 => [internal] load build definition from Dockerfile                                                                                                                0.0s
 => => transferring dockerfile: 454B                                                                                                                                0.0s
 => [internal] load metadata for docker.io/library/golang:1.25-alpine                                                                                               0.7s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                                    0.7s
 => [auth] library/alpine:pull token for registry-1.docker.io                                                                                                       0.0s
 => [auth] library/golang:pull token for registry-1.docker.io                                                                                                       0.0s
 => [internal] load .dockerignore                                                                                                                                   0.0s
 => => transferring context: 306B                                                                                                                                   0.0s
 => [builder 1/6] FROM docker.io/library/golang:1.25-alpine@sha256:77dd832edf2752dafd030693bef196abb24dcba3a2bc3d7a6227a7a1dae73169                                 0.0s
 => [stage-1 1/6] FROM docker.io/library/alpine:latest@sha256:4bcff63911fcb4448bd4fdacec207030997caf25e9bea4045fa6c8c44de311d1                                      0.0s
 => [internal] load build context                                                                                                                                   0.0s
 => => transferring context: 540B                                                                                                                                   0.0s
 => CACHED [builder 2/6] WORKDIR /app                                                                                                                               0.0s
 => CACHED [builder 3/6] COPY go.mod go.sum* ./                                                                                                                     0.0s
 => CACHED [builder 4/6] RUN go mod download                                                                                                                        0.0s
 => [builder 5/6] COPY . .                                                                                                                                          0.0s
 => [builder 6/6] RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .                                                                             4.9s
 => CACHED [stage-1 2/6] RUN apk --no-cache add ca-certificates                                                                                                     0.0s
 => CACHED [stage-1 3/6] RUN addgroup -g 1001 appgroup &&     adduser -D -u 1001 -G appgroup appuser                                                                0.0s
 => CACHED [stage-1 4/6] WORKDIR /root/                                                                                                                             0.0s
 => CACHED [stage-1 5/6] COPY --from=builder /app/main .                                                                                                            0.0s
 => CACHED [stage-1 6/6] RUN chown appuser:appgroup main                                                                                                            0.0s
 => exporting to image                                                                                                                                              0.0s
 => => exporting layers                                                                                                                                             0.0s
 => => writing image sha256:64569c96bd0fa8ff50bcf45cdb8701eb31a024742448ed38010d4de25e4d4a54                                                                        0.0s
 => => naming to docker.io/library/gomaxprocs-demo                                                                                                                  0.0s
❯ ./run.sh
Running Docker container with resource limits...
=== Container-Aware GOMAXPROCS ===
GOMAXPROCS: 2
NumCPU: 12
Running Docker container without resource limits...
=== Container-Aware GOMAXPROCS ===
GOMAXPROCS: 12
NumCPU: 12
```