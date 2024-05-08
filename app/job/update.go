package jobs

import "github.com/robfig/cron/v3"

func updateGtfs() {
	c := cron.New()
	c.AddFunc("0 3 * * *", func() {
		Load()
	})
}
