// Package jobs
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-04-18 15:52
package jobs

import (
	"github.com/robfig/cron"
)

func Start() {
	// Start scheduled tasks

	c := cron.New()

	c.Start()
}
