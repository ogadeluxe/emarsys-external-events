package main

import (
	"log"

	"github.com/jasonlvhit/gocron"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ogadeluxe/emarsys-external-events/handler"
	"github.com/ogadeluxe/emarsys-external-events/model"
)

func main() {
	// init database model
	db, err := model.NewDB()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	handler := handler.WishlistHandler{Store: db}
	gocron.Every(1).Day().At("21:21").Do(handler.PushWishlist)
	<-gocron.Start()
}
