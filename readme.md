protoc
start protomodel.bat

swag init

docker
docker build -t dockerfile .
docker run --rm -p 5487:5487 -i dockerfile /bin/orderfood
https://philipzheng.gitbooks.io/docker_practice/content/dockerfile/instructions.html