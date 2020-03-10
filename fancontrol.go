package main

import (
	"fmt"
	"os"
	"time"

	"github.com/wfd3/go-rpio"
)

func main() {
	const turnOnTemp = 55
	const turnOffTemp = 45
	const sleepInterval = 10
	const gpioPin = 17

	err := rpio.Open()
	if err != nil {
		fmt.Println("open gpio failed.", err)
		os.Exit(-1)
	}
	defer rpio.Close()

	pin := rpio.Pin(gpioPin)
	pin.Mode(rpio.Output)

	for true {
		temperature, err := getTemperature()
		if err != nil {
			fmt.Println("get temperature failed", err)
			os.Exit(-1)
		}
		if temperature > turnOnTemp && pin.Read() != rpio.High {
			fmt.Printf("now temperature is %.1f, turning up the cpu fan.\n", temperature)
			pin.High()
		} else if temperature < turnOffTemp && pin.Read() != rpio.Low {
			fmt.Printf("now temperature is %.1f, turning off the cpu fan.\n", temperature)
			pin.Low()
		}

		time.Sleep(time.Second * sleepInterval)
	}
}
