ARG ALPINE_VERSION=3.11
ARG GOLANG_VERSION=1.14

# Generate self signed cert
# -------------------------
FROM alpine:${ALPINE_VERSION} as certs
RUN apk --no-cache --update add openssl
WORKDIR /tmp/certs
RUN /usr/bin/openssl req -x509 \
                     -newkey rsa:2048 \
                     -nodes \
                     -out cert.pem \
                     -keyout key.pem \
                     -days 365 \
                     -subj "/C=US/ST=OR/L=FooBar/O=SpamEggs"

# Run kootenay via air
# --------------------
FROM cosmtrek/air:latest
COPY --from=certs /tmp/certs /tls
ENTRYPOINT ["/go/bin/air"]
