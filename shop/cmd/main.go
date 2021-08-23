package main

import (
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/eventsourced"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/protocol"
	"github.com/lightbend/akkaserverless-go-sdk/shop"
	"log"
	"os"
)

func main() {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "8080")
	server, err := akkaserverless.New(protocol.Config{
		ServiceName:    "cloudstate.tck.model.EventSourcedTckModel", // the servicename the proxy gets to know about
		ServiceVersion: "0.2.0",
	})
	if err != nil {
		log.Fatalf("akkaserverless.New failed: %s", err)
	}

	err = server.RegisterEventSourced(&eventsourced.Entity{
		ServiceName:   "com.example.shop.ShopService",
		PersistenceID: "ShopCart",
		EntityFunc:    shop.NewShopCart,
	}, protocol.DescriptorConfig{
		Service: "shop.proto",
	}.AddDomainDescriptor("shopdomain.proto"))
	// end::register[]
	if err != nil {
		log.Fatalf("Akkaserverless failed to register entity: %s", err)
	}
	err = server.Run()
	if err != nil {
		log.Fatalf("Akkaserverless failed to run: %v", err)
	}
}