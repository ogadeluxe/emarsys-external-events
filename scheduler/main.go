package scheduler

import (
	"log"

	"github.com/jasonlvhit/gocron"
	"github.com/ogadeluxe/emarsys-external-events/handler"
)

var scheduler *gocron.Scheduler

func init() {
	if scheduler != nil {
		return
	}
	scheduler = gocron.NewScheduler()
}

// ScheduleWL -> Run PushWishlist scheduler
func ScheduleWL(task handler.Handler) {
	log.Println("Running scheduler push wishlist running...")
	scheduler.Every(1).Day().At("13:44").Do(task.PushWishlist)
}

// Execute : running all crons
func Execute() {
	<-scheduler.Start()
}
