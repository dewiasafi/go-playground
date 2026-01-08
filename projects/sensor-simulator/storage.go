package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVStorage struct {
	writer *csv.Writer
}

func formatFloat(v float64) string {
	return fmt.Sprintf("%.2f", v)
}

func NewCSVStorage(filename string) (*CSVStorage, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"timestamp", "sensor", "value"})

	return &CSVStorage{writer: writer}, nil
}

func (c *CSVStorage) Save(data SensorData) {
	c.writer.Write([]string{
		data.Timestamp.Format("2006-01-02 15:04:05"),
		data.Name,
		formatFloat(data.Value),
	})
	c.writer.Flush()
}
