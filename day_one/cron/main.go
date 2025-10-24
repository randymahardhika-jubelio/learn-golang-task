package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds()) // Enable seconds field

	// Add a job to run every 5 seconds
	_, err := c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("Job executed at", time.Now())
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	c.Start()

	// Keep the program running
	select {}
}