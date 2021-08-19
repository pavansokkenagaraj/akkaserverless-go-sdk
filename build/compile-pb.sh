#!/usr/bin/env bash

set -o nounset
set -o errexit
set -o pipefail

# protoc akkaserverless-proxy-protocol
protoc --go_out=paths=source_relative:akkaserverless/protocol \
  --go-grpc_out=paths=source_relative:akkaserverless/protocol \
  --proto_path=proto/akkaserverless/protocol discovery.proto \
  --proto_path=proto/akkaserverless/component component.proto

protoc --go_out=paths=source_relative:akkaserverless/action \
  --go-grpc_out=paths=source_relative:akkaserverless/action \
  --proto_path=proto \
  --proto_path=proto/akkaserverless/component/action action.proto

protoc --go_out=paths=source_relative:akkaserverless/protocol \
  --go-grpc_out=paths=source_relative:akkaserverless/protocol \
  --proto_path=proto \
  --proto_path=proto/akkaserverless/component/entity entity.proto

protoc --go_out=paths=source_relative:akkaserverless/entity \
  --go-grpc_out=paths=source_relative:akkaserverless/entity \
  --proto_path=proto \
  --proto_path=proto/akkaserverless/component/eventsourcedentity event_sourced_entity.proto

protoc --go_out=paths=source_relative:akkaserverless/entity \
  --go-grpc_out=paths=source_relative:akkaserverless/entity \
  --proto_path=proto \
  --proto_path=proto/akkaserverless/component/replicatedentity replicated_entity.proto

protoc --go_out=paths=source_relative:akkaserverless/entity \
  --go-grpc_out=paths=source_relative:akkaserverless/entity \
  --proto_path=proto \
  --proto_path=proto/akkaserverless/component/valueentity value_entity.proto

protoc --go_out=paths=source_relative:akkaserverless/view \
  --go-grpc_out=paths=source_relative:akkaserverless/view \
  --proto_path=proto \
  --proto_path=proto/akkaserverless/component/view view.proto


# protoc akkaserverless-sdk-protocol
protoc --go_out=paths=source_relative:akkaserverless \
  --go-grpc_out=paths=source_relative:akkaserverless \
  --proto_path=proto/akkaserverless component.proto \
  --proto_path=proto/akkaserverless eventing.proto \
  --proto_path=proto/akkaserverless legacy_entity_key.proto \
  --proto_path=proto/akkaserverless views.proto

protoc --go_out=paths=source_relative:akkaserverless \
  --go-grpc_out=paths=source_relative:akkaserverless \
  --proto_path=proto \
  --proto_path=proto/akkaserverless annotations.proto


# TODO: WIP - fix later PA1
## protoc akkaserverless-tck-protocol
#protoc --go_out=paths=source_relative:tck/action \
#  --go-grpc_out=paths=source_relative:tck/action \
#  --proto_path=proto \
#  --proto_path=proto/akkaserverless/tck/model/action action.proto
#
#protoc --go_out=paths=source_relative:tck/eventing \
#  --go-grpc_out=paths=source_relative:tck/eventing \
#  --proto_path=proto \
#  --proto_path=proto/akkaserverless/tck/model/eventing local_persistence_eventing.proto
#
#protoc --go_out=paths=source_relative:tck/eventsourcedentity \
#  --go-grpc_out=paths=source_relative:tck/eventsourcedentity \
#  --proto_path=proto \
#  --proto_path=proto/akkaserverless/tck/model/eventsourcedentity event_sourced_entity.proto
#
#protoc --go_out=paths=source_relative:tck/replicatedentity \
#  --go-grpc_out=paths=source_relative:tck/replicatedentity \
#  --proto_path=proto \
#  --proto_path=proto/akkaserverless/tck/model/replicatedentity replicated_entity.proto
#
#protoc --go_out=paths=source_relative:tck/valueentity \
#  --go-grpc_out=paths=source_relative:tck/valueentity \
#  --proto_path=proto \
#  --proto_path=proto/akkaserverless/tck/model/valueentity value_entity.proto
#
#protoc --go_out=paths=source_relative:tck/view \
#  --go-grpc_out=paths=source_relative:tck/view \
#  --proto_path=proto \
#  --proto_path=proto/akkaserverless/tck/model/view view.proto