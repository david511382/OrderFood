protoc
protoc --go_out=. src/database/models/member.proto
protoc --go_out=. src/handler/user/reqs/models.proto
swag init