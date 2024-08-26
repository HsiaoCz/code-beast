package main

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	c := cron.New(cron.WithSeconds())
	spec := "*/5*****"
	c.AddFunc(spec, func() {
		logrus.WithFields(logrus.Fields{
			"method": "log",
		}).Info("some really bad thing hapend")
	})
	c.Start()
}
