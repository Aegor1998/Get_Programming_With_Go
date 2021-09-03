package main

import (
	"fmt"
	"math/rand"
	"time"
)

type k2 float64

//sensor function type
type sensor func() k2

func realSensor() k2 {
	return 0
}

func calibrate(s sensor, offset k2) sensor {
	return func() k2 {
		return s() + offset
	}
}

func main() {
	sensor := calibrate(fakeSensor, 5)
	fmt.Println(sensor())
}

func fakeSensor() k2 {
	rand.Seed(time.Now().UnixNano())
	return k2(rand.Intn(151) + 150)
}
