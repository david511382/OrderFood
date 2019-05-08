protoc
start protomodel.bat

go run main.go -ip
go build -tags release
./orderfood -ip

docker
docker build -t dockerfile .
docker run --rm -p 5487:5487 -i dockerfile /bin/orderfood
https://philipzheng.gitbooks.io/docker_practice/content/dockerfile/instructions.html