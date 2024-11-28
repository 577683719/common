package initialization

import (
	"ecs_cloud/common/util/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitDeefault() {
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		sig := <-sigChan
		logger.LogrusObj.Errorf("程序被关闭: %s at %s", sig, time.Now().Format(time.RFC3339))
		os.Exit(0)
	}()
}
