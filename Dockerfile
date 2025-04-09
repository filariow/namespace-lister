# Build the manager binary
FROM registry.access.redhat.com/ubi9/go-toolset@sha256:d5f8518b974240a8894a0c042fb2af00f22ec1309f6c545d6436a36fd9105969 AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /namespace-lister

# Copy the Go Modules manifests
COPY go.mod go.sum ./
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . .

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -trimpath -a -o /tmp/server .

FROM registry.access.redhat.com/ubi9/ubi-micro@sha256:dfd6e38dfecbb306899a548be0509c2ea16ff1cf843845682948e755de1ae948
WORKDIR /
COPY --from=builder /tmp/server .
USER 65532:65532

ENTRYPOINT ["/server"]
