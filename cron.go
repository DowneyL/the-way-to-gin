package main

import (
	"github.com/DowneyL/the-way-to-gin/models"
	"github.com/DowneyL/the-way-to-gin/pkg/logging"
	"github.com/DowneyL/the-way-to-gin/pkg/setting"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	log.Println("Starting")

	c := cron.New()
	_ = c.AddFunc("* * * * * *", func() {
		log.Println("Run clean deleted tags")
		models.CleanTag()
	})
	_ = c.AddFunc("* * * * * *", func() {
		log.Println("Run clean deleted article")
		models.CleanArticle()
	})

	c.Start()

	t := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-t.C:
			t.Reset(10 * time.Second)
		}
	}
}
