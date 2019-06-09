protoc
go get -u github.com/golang/protobuf/protoc-gen-go
start protomodel.bat

go run main.go -ip
go build -tags release
./orderfood -ip

go test -v -count=1 ./src/database/mysql

docker
docker-compose -f ./docker/docker-compose.yml up -d --build
https://philipzheng.gitbooks.io/docker_practice/content/dockerfile/instructions.html

go get github.com/akavel/rsrc
rsrc -manifest nac.manifest -o nac.syso

swag init