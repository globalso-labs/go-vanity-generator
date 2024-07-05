# syntax = docker/dockerfile:1.2

# get modules, if they don't change the cache can be used for faster builds
FROM golang:1.22 AS base
ARG ACCESS_TOKEN

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /src
COPY go.* .

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# build th application
FROM base AS build
# temp mount all files instead of loading into image with COPY
# temp mount module cache
# temp mount go build cache
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-w -s" -o /app/main ./cmd/*.go

# Load the busybox image
FROM busybox:1.34.1-uclibc as busybox

# Import the binary from build stage
FROM gcr.io/distroless/static:nonroot as release
COPY --from=build --chown=65534:65534 --chmod=755 /app/main /app/main
# this is the numeric version of user nobody:nobody to check runAsNonRoot in kubernetes
USER 65534:65534
WORKDIR /app
CMD ["/app/main", "--help"]

FROM release as debug
COPY --from=busybox /bin/sh /bin/sh
COPY --from=busybox /bin/ls /bin/ls
COPY --from=busybox /bin/cat /bin/cat
COPY --from=busybox /bin/vi /bin/vi
COPY --from=busybox /bin/id /bin/id

# Run the server
CMD ["/bin/sh", "-c", "/app/main --help"]
