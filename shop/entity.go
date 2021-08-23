package shop

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
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
	switch c := cmd.(type) {
	case *GetShopCart:
		return s.GetCart(ctx, c)
	case *RemoveLineItem:
		return s.RemoveItem(ctx, c)
	case *AddLineItem:
		return s.AddItem(ctx, c)
	default:
		return nil, nil
	}
}


// AddItem implements the AddItem command handling of the shopping cart service.
// tag::add-item[]
func (s *ShopCart) AddItem(ctx *eventsourced.Context, li *AddLineItem) (*empty.Empty, error) {
	if li.GetQuantity() <= 0 {
		return nil, fmt.Errorf("cannot add negative quantity of to item %q", li.GetProductId())
	}
	ctx.Emit(&persistence.ItemAdded{
		Item: &persistence.LineItem{
			ProductId: li.ProductId,
			Name:      li.Name,
			Quantity:  li.Quantity,
		}})
	return &empty.Empty{}, nil
}

// end::add-item[]

// RemoveItem implements the RemoveItem command handling of the shopping cart service.
func (s *ShopCart) RemoveItem(ctx *eventsourced.Context, li *RemoveLineItem) (*empty.Empty, error) {
	if item, _ := s.find(li.GetProductId()); item == nil {
		return nil, fmt.Errorf("cannot remove item %s because it is not in the cart", li.GetProductId())
	}
	ctx.Emit(&persistence.ItemRemoved{ProductId: li.ProductId})
	return &empty.Empty{}, nil
}


// find finds a product in the shopping cart by productId and returns it as a LineItem.
func (s *ShopCart) find(productID string) (item *persistence.LineItem, index int) {
	for i, item := range s.cart {
		if productID == item.ProductId {
			return item, i
		}
	}
	return nil, 0
}

// GetCart implements the GetCart command handling of the shopping cart service.
// tag::get-cart[]
func (s *ShopCart) GetCart(*eventsourced.Context, *GetShopCart) (*Cart, error) {
	cart := &Cart{}
	for _, item := range s.cart {
		cart.Items = append(cart.Items, &LineItem{
			ProductId: item.ProductId,
			Name:      item.Name,
			Quantity:  item.Quantity,
		})
	}
	return cart, nil
}

func (s ShopCart) HandleEvent(ctx *eventsourced.Context, event interface{}) error {
	switch e := event.(type) {
	case *persistence.ItemAdded:
		return s.ItemAdded(e)
	case *persistence.ItemRemoved:
		return s.ItemRemoved(e)
	default:
		return nil
	}
}


// ItemAdded is a event handler function for the ItemAdded event.
// tag::item-added[]
func (s *ShopCart) ItemAdded(added *persistence.ItemAdded) error {
	if added.Item.GetName() == "FAIL" {
		return errors.New("boom: forced an unexpected error")
	}
	if item, _ := s.find(added.Item.ProductId); item != nil {
		item.Quantity += added.Item.Quantity
		return nil
	}
	s.cart = append(s.cart, &persistence.LineItem{
		ProductId: added.Item.ProductId,
		Name:      added.Item.Name,
		Quantity:  added.Item.Quantity,
	})
	return nil
}

// end::item-added[]

// ItemRemoved is a event handler function for the ItemRemoved event.
func (s *ShopCart) ItemRemoved(removed *persistence.ItemRemoved) error {
	if !s.remove(removed.ProductId) {
		return errors.New("unable to remove product")
	}
	return nil
}

// remove removes a product from the shopping cart.
// An ok flag is returned to indicate that the product was present and removed.
func (s *ShopCart) remove(productID string) (ok bool) {
	if item, i := s.find(productID); item != nil {
		// remove and re-slice
		copy(s.cart[i:], s.cart[i+1:])
		s.cart = s.cart[:len(s.cart)-1]
		return true
	}
	return false
}


func (s ShopCart)Snapshot(ctx *eventsourced.Context) (snapshot interface{}, err error){
	return &persistence.Cart{
		Items: append(make([]*persistence.LineItem, 0, len(s.cart)), s.cart...),
	}, nil
}

func (s ShopCart) HandleSnapshot(ctx *eventsourced.Context, snapshot interface{}) error{
	switch value := snapshot.(type) {
	case *persistence.Cart:
		s.cart = append(s.cart[:0], value.Items...)
		return nil
	default:
		return fmt.Errorf("unknown snapshot type: %v", value)
	}
}
