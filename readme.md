protoc
protoc --go_out=. src/database/models/member.proto
protoc --go_out=. src/handler/models/reqs/user.proto
protoc --go_out=. src/handler/models/resp/menu.proto
swag init