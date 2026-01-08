package main

import (
	"math/rand"
	"time"
)

type TemperatureSensor struct {
	Name  string
	value float64
}

func (t TemperatureSensor) Read() SensorData {
	if t.value == 0 {
		t.value = 26 + rand.Float64()*2
	}

	delta := (rand.Float64() - 0.5) * 0.4
	t.value += delta

	if t.value > 40 {
		t.value = 40
	}

	return SensorData{
		Name:      t.Name,
		Value:     t.value,
		Timestamp: time.Now(),
	}
}
