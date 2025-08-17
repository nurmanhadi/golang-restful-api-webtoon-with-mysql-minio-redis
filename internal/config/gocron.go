package config

import (
	"log"

	"github.com/go-co-op/gocron/v2"
)

func NewGocron() gocron.Scheduler {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}
	return s
}
