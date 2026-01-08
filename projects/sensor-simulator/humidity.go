package main

import (
	"math/rand"
	"time"
)

type HumiditySensor struct {
	Name  string
	value float64
}

func (h HumiditySensor) Read() SensorData {
	if h.value == 0 {
		h.value = 50 + rand.Float64()*10
	}

	delta := (rand.Float64() - 0.5) * 2 // ±1
	h.value += delta

	if h.value < 40 {
		h.value = 40
	}

	if h.value > 80 {
		h.value = 80
	}

	return SensorData{
		Name:      h.Name,
		Value:     h.value,
		Timestamp: time.Now(),
	}
}
