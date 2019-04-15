for %%f in (src/handler/models/resp/*.proto) do protoc --go_out=. src/handler/models/resp/%%f
for %%f in (src/handler/models/reqs/*.proto) do protoc --go_out=. src/handler/models/reqs/%%f
for %%f in (src/database/models/*.proto) do protoc --go_out=. src/database/models/%%f