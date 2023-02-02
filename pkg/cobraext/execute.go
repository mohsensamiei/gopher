package cobraext

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

func Execute(root *cobra.Command, commanders []CommanderRegister) {
	for _, commander := range commanders {
		commander.RegisterCommander(root)
	}
	start := time.Now()
	if err := root.Execute(); err != nil {
		log.WithField("duration", fmt.Sprint(time.Since(start))).
			WithError(err).Error("cli exec failed")
	} else {
		log.WithField("duration", fmt.Sprint(time.Since(start))).
			Info("cli exec succeeded")
	}
}
