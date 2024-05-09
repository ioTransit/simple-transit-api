package jobs

import (
	"fmt"
	"os"

	"github.com/robfig/cron/v3"
)

func UpdateGtfs() *cron.Cron {
	updateInterval := os.Getenv("UPDATE_INTERVAL")
	fmt.Printf("interval %s", updateInterval)
	if updateInterval == "" {
		return nil
	}
	c := cron.New()
	c.AddFunc(updateInterval, func() {
		Load(true)
	})
	c.Start()
	return c
}
