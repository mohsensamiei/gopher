package service

import (
	"fmt"
	"github.com/pinosell/gopher/pkg/netext"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func Start() {
	interrupt := make(chan error)
	for p, s := range serves {
		go func(p netext.Port, s *serve) {
			log.WithFields(log.Fields{
				"port":  p,
				"serve": s.Name,
			}).Info("service start serving")
			interrupt <- s.Listen(p)
		}(p, s)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		interrupt <- fmt.Errorf((<-signalChan).String())
	}()

	if err := <-interrupt; err != nil {
		log.WithError(err).Panic("service interrupted")
	}
	log.Panic("service interrupted")
}
