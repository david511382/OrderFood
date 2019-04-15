protoc
protoc --go_out=. src/database/models/member.proto
protoc --go_out=. src/handler/models/reqs/user.proto
protoc --go_out=. src/handler/models/resp/menu.proto
protoc --go_out=. src/handler/models/resp/shop.proto
swag init

docker
docker build -t dockerfile .
docker run --rm -p 5487:5487 -i dockerfile /bin/orderfood
https://philipzheng.gitbooks.io/docker_practice/content/dockerfile/instructions.html