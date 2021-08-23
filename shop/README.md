```shell
protoc --go_out=paths=source_relative:shop \
  --go-grpc_out=paths=source_relative:shop  \
  --proto_path=proto/sdk   \
  --proto_path=shop  shop.proto 
  
protoc --go_out=paths=source_relative:shop/persistance \
  --go-grpc_out=paths=source_relative:shop/persistance  \
  --proto_path=proto/sdk   \
  --proto_path=shop/persistance  shopdomain.proto   
```