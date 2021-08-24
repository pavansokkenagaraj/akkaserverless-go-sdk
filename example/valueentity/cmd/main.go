//
// Copyright 2019 Lightbend Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main implements a value entity shopping cart example
package main

import (
	"log"
	"time"

	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/protocol"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/value"
	"github.com/lightbend/akkaserverless-go-sdk/example/valueentity"
)

// main creates an Akkaserverless instance and registers the ShoppingCart as a value entity.
func main() {
	server, err := akkaserverless.New(protocol.Config{
		ServiceName:    "shopping-cart",
		ServiceVersion: "0.0.1",
	})
	if err != nil {
		log.Fatalf("akkaserverless.New failed: %v", err)
	}

	err = server.RegisterValueEntity(&value.Entity{
		ServiceName:   "com.example.valueentity.shoppingcart.ShoppingCartService",
		PersistenceID: "ShoppingCart",
		EntityFunc:    valueentity.NewShoppingCart,
	}, protocol.DescriptorConfig{
		Service: "value_shoppingcart.proto",
	}, value.WithPassivationStrategyTimeout(10000*time.Millisecond))

	if err != nil {
		log.Fatalf("Akkaserverless failed to register entity: %v", err)
	}
	err = server.Run()
	if err != nil {
		log.Fatalf("Akkaserverless failed to run: %v", err)
	}
}
