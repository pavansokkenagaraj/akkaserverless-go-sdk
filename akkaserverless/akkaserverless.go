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

package akkaserverless

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/action"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/crdt"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/discovery"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/entity"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/eventsourced"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/protocol"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/value"
	"google.golang.org/grpc"
)

// Akkaserverless is an instance of a Akkaserverless User Function.
type Akkaserverless struct {
	grpcServer             *grpc.Server
	entityDiscoveryServer  *discovery.EntityDiscoveryServer
	eventSourcedServer     *eventsourced.Server
	replicatedEntityServer *crdt.Server
	actionServer           *action.Server
	valueServer            *value.Server
}

// New returns a new Akkaserverless instance.
func New(c protocol.Config) (*Akkaserverless, error) {
	cs := &Akkaserverless{
		grpcServer:             grpc.NewServer(),
		entityDiscoveryServer:  discovery.NewServer(c),
		eventSourcedServer:     eventsourced.NewServer(),
		replicatedEntityServer: crdt.NewServer(),
		actionServer:           action.NewServer(),
		valueServer:            value.NewServer(),
	}
	protocol.RegisterDiscoveryServer(cs.grpcServer, cs.entityDiscoveryServer)
	entity.RegisterEventSourcedEntitiesServer(cs.grpcServer, cs.eventSourcedServer)
	entity.RegisterValueEntitiesServer(cs.grpcServer, cs.valueServer)
	action.RegisterActionsServer(cs.grpcServer, cs.actionServer)
	entity.RegisterReplicatedEntitiesServer(cs.grpcServer, cs.replicatedEntityServer)
	return cs, nil
}

// RegisterEventSourced registers an event sourced entity.
func (s *Akkaserverless) RegisterEventSourced(entity *eventsourced.Entity, config protocol.DescriptorConfig, options ...eventsourced.Option) error {
	entity.Options(options...)
	if err := s.eventSourcedServer.Register(entity); err != nil {
		return err
	}
	if err := s.entityDiscoveryServer.RegisterEventSourcedEntity(entity, config); err != nil {
		return err
	}
	return nil
}

// RegisterCRDT registers a CRDT entity.
func (s *Akkaserverless) RegisterCRDT(entity *crdt.Entity, config protocol.DescriptorConfig, options ...crdt.Option) error {
	entity.Options(options...)
	if err := s.replicatedEntityServer.Register(entity); err != nil {
		return err
	}
	if err := s.entityDiscoveryServer.RegisterCRDTEntity(entity, config); err != nil {
		return err
	}
	return nil
}

// RegisterAction registers an action entity.
func (s *Akkaserverless) RegisterAction(entity *action.Entity, config protocol.DescriptorConfig) error {
	if err := s.actionServer.Register(entity); err != nil {
		return err
	}
	if err := s.entityDiscoveryServer.RegisterActionEntity(entity, config); err != nil {
		return err
	}
	return nil
}

// RegisterValueEntity registers a Value entity.
func (s *Akkaserverless) RegisterValueEntity(entity *value.Entity, config protocol.DescriptorConfig, options ...value.Option) error {
	entity.Options(options...)
	if err := s.valueServer.Register(entity); err != nil {
		return err
	}
	if err := s.entityDiscoveryServer.RegisterValueEntity(entity, config); err != nil {
		return err
	}
	return nil
}

// Run runs the Akkaserverless instance on the interface and port defined by
// the HOST and PORT environment variable.
func (s *Akkaserverless) Run() error {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		return errors.New("unable to get environment variable \"HOST\"")
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return errors.New("unable to get environment variable \"PORT\"")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	if err := s.RunWithListener(lis); err != nil {
		return fmt.Errorf("failed to RunWithListener for: %v with: %w", lis, err)
	}
	return nil
}

// Run runs the Akkaserverless instance with a listener provided.
func (s *Akkaserverless) RunWithListener(lis net.Listener) error {
	return s.grpcServer.Serve(lis)
}

// Stop gracefully stops the Akkaserverless instance.
func (s *Akkaserverless) Stop() {
	s.grpcServer.GracefulStop()
	log.Println("Akkaserverless stopped")
}
