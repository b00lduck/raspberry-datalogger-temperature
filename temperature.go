package main

import (
	"time"
	log "github.com/Sirupsen/logrus"
	"github.com/b00lduck/raspberry-datalogger-temperature/sensor"
)

func main() {
    log.Info("Starting temperature acqusition service")
    defer log.Info("Exiting temperature acqusition service")

    var therm1 = sensor.NewThermometer("temperature_boiler_room", "28-031655cbb5ff", 0.2)
    var therm2 = sensor.NewThermometer("temperature_hallway_1", "28-041658b369ff", 0.2)
    var therm3 = sensor.NewThermometer("temperature_hallway_2", "28-041658b3cbff", 0.2)

    for {
        therm1.Process()
        therm2.Process()
        therm3.Process()
        time.Sleep(5 * time.Second)
    }
}


