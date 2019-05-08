protoc
go get -u github.com/golang/protobuf/protoc-gen-go
start protomodel.bat

go run main.go -ip
go build -tags release
./orderfood -ip

docker
docker build -t dockerfile .
docker run --rm -p 5487:5487 -i dockerfile /bin/orderfood
https://philipzheng.gitbooks.io/docker_practice/content/dockerfile/instructions.html

go get github.com/akavel/rsrc
rsrc -manifest nac.manifest -o nac.syso