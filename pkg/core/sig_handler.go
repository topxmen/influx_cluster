package core

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
)

func SignalHandling() {
	signal_chan := make(chan os.Signal, 1)

	signal.Notify(signal_chan, syscall.SIGINT, syscall.SIGTERM)
	exit_chan := make(chan int)
	go func() {
		for {
			<-signal_chan
			exit_chan <- 0
			glog.Info("Exit")
		}
	}()

	code := <-exit_chan
	glog.Flush()
	os.Exit(code)
}
