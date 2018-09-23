package handler

import (
	"log"

	"github.com/ogadeluxe/emarsys-external-events/config"
	"github.com/ogadeluxe/emarsys-external-events/library"
	"github.com/ogadeluxe/emarsys-external-events/model"
)

// WishlistHandler : handle pushing wishlist to Emarsys
type WishlistHandler struct {
	Store model.Datastore
}

// InitHandler -Initialize handler
func InitHandler(dataStore model.Datastore) *WishlistHandler {
	handler := new(WishlistHandler)
	handler.Store = dataStore
	return handler
}

// PushWishlist : push all wishlist to emarsys
func (handler *WishlistHandler) PushWishlist() {
	payload := library.BuildWLPayload(handler.Store)
	if payload == "" {
		log.Panic("Empty payload")
	}

	newEvent := config.Items.EventID + "/trigger"
	res := library.TriggerEvent("POST", newEvent, payload)
	log.Println("response: ", res.Body)
}
