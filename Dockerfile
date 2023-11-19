# Build the manager binary
# FROM golang:1.20 as builder
FROM golang:1.20

# FROM ubuntu:latest

WORKDIR /home/oqs


# Install dependencies
RUN apt -y update && \
    apt install -y build-essential cmake libssl-dev

# Get liboqs
RUN git clone --depth 1 --branch main https://github.com/open-quantum-safe/liboqs

# Install liboqs
RUN cmake -S liboqs -B liboqs/build -DBUILD_SHARED_LIBS=ON && \
    cmake --build liboqs/build --parallel 4 && \
    cmake --build liboqs/build --target install

# Enable a normal user
# RUN useradd -m -c "Open Quantum Safe" oqs
# USER oqs
# WORKDIR /home/oqs

# Get liboqs-go
RUN git clone --depth 1 --branch main https://github.com/open-quantum-safe/liboqs-go.git

# Configure liboqs-go
ENV PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/home/oqs/liboqs-go/.config
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib

# COPY --from=builder /workspace/manager .

# ENTRYPOINT ["./manager"]





ARG TARGETOS
ARG TARGETARCH

# WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd/main.go cmd/main.go
COPY api/ api/
COPY internal/ internal/

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
# RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o manager cmd/main.go
RUN go build -a -o /manager cmd/main.go

# RUN useradd -m -c "Open Quantum Safe" oqs
# USER oqs
USER 65532:65532


ENTRYPOINT ["/manager"]


# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# FROM gcr.io/distroless/static:nonroot
# WORKDIR /
# COPY --from=builder /workspace/manager .
# USER 65532:65532

# ENTRYPOINT ["/manager"]
