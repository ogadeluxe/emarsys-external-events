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

var h Handler

// InitWishlist -Initialize handler
func InitWishlist(dataStore model.Datastore) Handler {
	if h == nil {
		h = WishlistHandler{Store: dataStore}
	}

	return h
}

// PushWishlist : push all wishlist to emarsys
func (handler WishlistHandler) PushWishlist() {
	payload := library.BuildWLPayload(handler.Store)
	if payload == "" {
		log.Panic("Empty payload")
	}

	newEvent := config.Items.EventID + "/trigger"
	res := library.TriggerEvent("POST", newEvent, payload)
	log.Println("response: ", res.Body)
}
