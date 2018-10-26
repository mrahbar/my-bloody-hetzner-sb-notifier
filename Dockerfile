FROM golang:1.11-stretch
ENV VERSION=1.0
ADD ./builds/hetzner-sb-notifier_linux_amd64_$VERSION /root/hetzner-sb-notifier
ENTRYPOINT ["/root/hetzner-sb-notifier"]