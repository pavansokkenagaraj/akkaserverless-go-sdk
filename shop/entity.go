package shop

import (
	"github.com/golang/protobuf/proto"
	"github.com/lightbend/akkaserverless-go-sdk/akkaserverless/eventsourced"
	"github.com/lightbend/akkaserverless-go-sdk/shop/persistance"
)

func NewShopCart(eventsourced.EntityID) eventsourced.EntityHandler {
	return &ShopCart{
		cart: make([]*persistence.LineItem, 0),
	}
}

type ShopCart struct {
	// our domain object
	cart []*persistence.LineItem
}

func (s ShopCart) HandleCommand(ctx *eventsourced.Context, name string, cmd proto.Message) (reply proto.Message, err error) {
	panic("implement me")
}

func (s ShopCart) HandleEvent(ctx *eventsourced.Context, event interface{}) error {
	panic("implement me")
}

func (s ShopCart)Snapshot(ctx *eventsourced.Context) (snapshot interface{}, err error){
	panic("implement me")
}

func (s ShopCart) HandleSnapshot(ctx *eventsourced.Context, snapshot interface{}) error{
	panic("implement me")
}
