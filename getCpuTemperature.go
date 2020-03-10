package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func getTemperature() (float64, error) {
	cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
	stdin, err := cmd.StdinPipe()
	defer stdin.Close()
	if err != nil {
		return 0, err
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	numberStr := strings.TrimSuffix(string(out), "\n")
	temperature, err := strconv.ParseFloat(numberStr, 32)
	return temperature / 1000, err
}
