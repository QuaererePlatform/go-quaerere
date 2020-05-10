ARG ALPINE_VERSION=3.11
ARG GOLANG_VERSION=1.14

# Root certificate bundle
# -----------------------
FROM alpine:${ALPINE_VERSION} as certs
RUN apk --update add ca-certificates

# Build kootenay server
# ---------------------
FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} as builder
WORKDIR /tmp/build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build_columbia

# Package kootenay
# ----------------
FROM alpine:${ALPINE_VERSION}
WORKDIR /opt/columbia
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /tmp/build/columbia .
COPY ./build/docker/prod/entrypoint.sh /usr/local/bin/entrypoint.sh
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
