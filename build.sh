GO111MODULE=on
go get github.com/mitchellh/gox
go get ./...
gox -output="build/hetzner-sb-notifier_{{.OS}}_{{.Arch}}_${VERSION}" -osarch=$1/amd64