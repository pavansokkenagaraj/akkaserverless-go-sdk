# code gen 
```shell
```shell
protoc --go_out=paths=source_relative:example/shoppingcart \
  --go-grpc_out=paths=source_relative:example/shoppingcart  \
  --proto_path=proto/sdk   \
  --proto_path=example/shoppingcart shoppingcart.proto 
  
protoc --go_out=paths=source_relative:example/shoppingcart/persistence \
  --go-grpc_out=paths=source_relative:example/shoppingcart/persistence  \
  --proto_path=proto/sdk   \
  --proto_path=example/shoppingcart/persistence domain.proto   
```
```