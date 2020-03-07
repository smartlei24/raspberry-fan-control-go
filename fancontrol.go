package main

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	const turnOnTemp = 60
	const turnOffTemp = 50
	const sleepInterval = 10
	const gpioPin = 17

	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}

	pin := rpio.Pin(gpioPin)
	for true {
		temperature, err := getTemperature()
		if err != nil {
			log.Fatal(err)
		}
		if temperature > turnOnTemp && pin.Read() != rpio.High {
			pin.High()
		} else if temperature < turnOffTemp && pin.Read() != rpio.Low {
			pin.Low()
		}

		time.Sleep(time.Second * sleepInterval)
	}
}
