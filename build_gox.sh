VERSION=1.0
GO111MODULE=on

echo "Fetching dependencies"
go get ./...
go get github.com/mitchellh/gox
echo "Building project"
mkdir -p builds
gox -output="builds/hetzner-sb-notifier_{{.OS}}_{{.Arch}}_$VERSION"
echo "Done"
