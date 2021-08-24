# code gen 

### shoppingcart
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


### valueentity
```shell
protoc --go_out=paths=source_relative:example/valueentity \
  --go-grpc_out=paths=source_relative:example/valueentity  \
  --proto_path=proto/sdk   \
  --proto_path=example/valueentity value_shoppingcart.proto 
  
protoc --go_out=paths=source_relative:example/valueentity/persistence \
  --go-grpc_out=paths=source_relative:example/valueentity/persistence  \
  --proto_path=proto/sdk   \
  --proto_path=example/valueentity/persistence value_domain.proto   
```
