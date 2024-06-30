package scheduler

import (
	"context"
	"github.com/go-co-op/gocron/v2"
	"log"
	"online-shop-api/internal/helper"
	"online-shop-api/internal/service"
	"time"
)

type Scheduler struct {
	Scheduler    gocron.Scheduler
	OrderService service.OrderService
}

func NewScheduler(orderService service.OrderService) *Scheduler {
	s, _ := gocron.NewScheduler()
	return &Scheduler{Scheduler: s, OrderService: orderService}
}

func (scheduler *Scheduler) ScheduleCancelOrder() {
	s, err := gocron.NewScheduler()
	helper.PanicIfError(err)

	_, err = s.NewJob(
		gocron.DurationJob(5*time.Minute),
		gocron.NewTask(func() {
			defer s.Shutdown()

			err := scheduler.OrderService.TaskCancelOrder(context.Background())
			if err != nil {
				log.Printf("Error cancel order: %v", err)
			}
			log.Printf(time.Now().String() + " - Cancel order task is running")
		}),
	)

	s.Start()
}
