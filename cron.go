package main

import (
	"github.com/DowneyL/the-way-to-gin/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
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
