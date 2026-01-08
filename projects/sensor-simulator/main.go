package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func startSensor(
	ctx context.Context,
	sensor Sensor,
	interval time.Duration,
	ch chan<- SensorData,
) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ch <- sensor.Read()
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dataChan := make(chan SensorData)

	storage, err := NewCSVStorage("sensor_data.csv")

	if err != nil {
		fmt.Println("Gagal buat storage:", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	temp := TemperatureSensor{Name: "Temperature"}
	hum := HumiditySensor{Name: "Humidity"}

	go startSensor(ctx, temp, 1*time.Second, dataChan)
	go startSensor(ctx, hum, 1*time.Second, dataChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	displayTicker := time.NewTicker(3 * time.Second)
	defer displayTicker.Stop()

	fmt.Println("Sensor simulator berjalan... Ctrl+C untuk berhenti")

	latest := make(map[string]SensorData)

	for {
		select {
		case data := <-dataChan:
			latest[data.Name] = data
			storage.Save(data)
		case <-displayTicker.C:
			fmt.Println("\n=== HASIL PENGUKURAN ===")
			for _, data := range latest {
				fmt.Printf("[%s] %s = %.2f\n",
					data.Timestamp.Format("15:04:05"),
					data.Name,
					data.Value,
				)
				if data.Name == "Temperature" && data.Value > 28 {
					fmt.Println("⚠️ ALERT: Suhu tinggi")
				}
			}
		case <-signalChan:
			fmt.Println("\nProgram dihentikan.")
			return
		}
	}
}
