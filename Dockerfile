FROM golang:alpine as builder

RUN apk update && apk add git && apk add ca-certificates
RUN adduser -D -g '' appuser
RUN mkdir /hetzner-sb-notifier
COPY . /hetzner-sb-notifier/
WORKDIR /hetzner-sb-notifier

ENV VERSION=1.0
RUN chmod +x build.sh

RUN /hetzner-sb-notifier/build.sh linux

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /hetzner-sb-notifier/build/hetzner-sb-notifier_linux_amd64 /hetzner-sb-notifier
RUN chmod +x hetzner-sb-notifier
USER appuser
ENTRYPOINT ["/hetzner-sb-notifier"]