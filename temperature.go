package main

import (
	"fmt"
	"time"
	log "github.com/Sirupsen/logrus"
        "sensor"
)

func main() {
    log.Info("Starting temperature acqusition service")
    defer log.Info("Exiting temperature acqusition service")
    for {
	time.Sleep(5 * time.Second)
    }
}

