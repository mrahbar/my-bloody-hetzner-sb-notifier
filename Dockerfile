FROM centurylink/ca-certs
ENV VERSION=1.0
ADD ./build/hetzner-sb-notifier_linux_amd64_1.0 /
ENTRYPOINT ["/hetzner-sb-notifier_linux_amd64_1.0"]