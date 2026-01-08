package main

import "time"

type SensorData struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

type Sensor interface {
	Read() SensorData
}
