#!/usr/bin/env bash

set -o nounset
set -o errexit
set -o pipefail


# protoc akkaserverless-sdk-protocol
protoc --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  --proto_path=proto/sdk akkaserverless/component.proto \
  --proto_path=proto/sdk akkaserverless/eventing.proto \
  --proto_path=proto/sdk akkaserverless/legacy_entity_key.proto \
  --proto_path=proto/sdk akkaserverless/views.proto

protoc --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  --proto_path=proto/sdk \
  --proto_path=proto/sdk akkaserverless/annotations.proto


# protoc akkaserverless-proxy-protocol
protoc --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  --proto_path=proto/proxy akkaserverless/protocol/discovery.proto

protoc --go_out=paths=source_relative:akkaserverless/protocol \
  --go-grpc_out=paths=source_relative:akkaserverless/protocol \
  --proto_path=proto/proxy/akkaserverless/component component.proto

protoc --go_out=paths=source_relative:akkaserverless/action \
  --go-grpc_out=paths=source_relative:akkaserverless/action \
  --proto_path=proto/proxy \
  --proto_path=proto/proxy/akkaserverless/component/action action.proto

protoc --go_out=paths=source_relative:akkaserverless/protocol \
  --go-grpc_out=paths=source_relative:akkaserverless/protocol \
  --proto_path=proto/proxy \
  --proto_path=proto/proxy/akkaserverless/component/entity entity.proto

protoc --go_out=paths=source_relative:akkaserverless/entity \
  --go-grpc_out=paths=source_relative:akkaserverless/entity \
  --proto_path=proto/proxy \
  --proto_path=proto/proxy/akkaserverless/component/eventsourcedentity event_sourced_entity.proto

protoc --go_out=paths=source_relative:akkaserverless/entity \
  --go-grpc_out=paths=source_relative:akkaserverless/entity \
  --proto_path=proto/proxy \
  --proto_path=proto/proxy/akkaserverless/component/replicatedentity replicated_entity.proto

protoc --go_out=paths=source_relative:akkaserverless/entity \
  --go-grpc_out=paths=source_relative:akkaserverless/entity \
  --proto_path=proto/proxy \
  --proto_path=proto/proxy/akkaserverless/component/valueentity value_entity.proto

protoc --go_out=paths=source_relative:akkaserverless/view \
  --go-grpc_out=paths=source_relative:akkaserverless/view \
  --proto_path=proto/proxy \
  --proto_path=proto/proxy/akkaserverless/component/view view.proto

## protoc akkaserverless-tck-protocol
protoc --go_out=paths=source_relative:tck \
  --go-grpc_out=paths=source_relative:tck \
  --proto_path=proto/sdk \
  --proto_path=proto/tck/akkaserverless/tck/model action/action.proto \
  --proto_path=proto/tck/akkaserverless/tck/model eventing/local_persistence_eventing.proto \
  --proto_path=proto/tck/akkaserverless/tck/model eventsourcedentity/event_sourced_entity.proto \
  --proto_path=proto/tck/akkaserverless/tck/model replicatedentity/replicated_entity.proto \
  --proto_path=proto/tck/akkaserverless/tck/model valueentity/value_entity.proto \
  --proto_path=proto/tck/akkaserverless/tck/model view/view.proto