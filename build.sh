VERSION=1.0
GO111MODULE=on
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64

echo "Fetching dependencies"
go get ./...
echo "Building project"
mkdir builds
go build -a -installsuffix cgo -o ./builds/hetzner-sb-notifier_${GOOS}_${GOARCH}_${VERSION} .
echo "Done"
