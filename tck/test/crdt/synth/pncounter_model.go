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

package synth

import (
	"github.com/lightbend/akkaserverless-go-sdk/tck/crdt"
	"github.com/golang/protobuf/proto"
)

func pncounterRequest(messages ...proto.Message) *crdt.PNCounterRequest {
	r := &crdt.PNCounterRequest{
		Actions: make([]*crdt.PNCounterRequestAction, 0, len(messages)),
	}
	for _, i := range messages {
		switch t := i.(type) {
		case *crdt.PNCounterIncrement:
			r.Id = t.Key
			r.Actions = append(r.Actions, &crdt.PNCounterRequestAction{Action: &crdt.PNCounterRequestAction_Increment{Increment: t}})
		case *crdt.PNCounterDecrement:
			r.Id = t.Key
			r.Actions = append(r.Actions, &crdt.PNCounterRequestAction{Action: &crdt.PNCounterRequestAction_Decrement{Decrement: t}})
		case *crdt.Get:
			r.Id = t.Key
			r.Actions = append(r.Actions, &crdt.PNCounterRequestAction{Action: &crdt.PNCounterRequestAction_Get{Get: t}})
		case *crdt.Delete:
			r.Id = t.Key
			r.Actions = append(r.Actions, &crdt.PNCounterRequestAction{Action: &crdt.PNCounterRequestAction_Delete{Delete: t}})
		default:
			panic("no type matched")
		}
	}
	return r
}
