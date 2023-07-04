FROM golang:1-bullseye AS builder

WORKDIR /workdir/
COPY . /workdir/

RUN apt-get update

RUN update-ca-certificates

RUN make build

FROM debian:bullseye-slim

RUN apt-get update \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workdir/ghgrep ./usr/bin/gh-grep

ENTRYPOINT ["/entrypoint.sh"]

COPY scripts/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
