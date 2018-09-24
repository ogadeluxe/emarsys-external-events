package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ogadeluxe/emarsys-external-events/handler"
	"github.com/ogadeluxe/emarsys-external-events/model"
	"github.com/ogadeluxe/emarsys-external-events/scheduler"
)

func main() {
	// init database model
	db, err := model.NewDB()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	handler := handler.InitWishlist(db)
	scheduler.ScheduleWL(handler)
	scheduler.Execute()
}
