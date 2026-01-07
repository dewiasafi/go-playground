package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Pakai: go run main.go namafile")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Gagal buka file:", err)
		return
	}
	defer file.Close()

	outputFile, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Gagal bikin result.txt:", err)
		return
	}
	defer outputFile.Close()

	errorByService := make(map[string]map[string]map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, "ERROR") {
			continue
		}

		part := strings.Split(line, "]")

		if len(part) < 3 {
			continue
		}

		service := strings.TrimPrefix(part[0], "[")
		endpoint := strings.TrimPrefix(part[1], "[")

		message := strings.Split(line, "ERROR")[1]

		if errorByService[service] == nil {
			errorByService[service] = make(map[string]map[string]int)
		}

		if errorByService[service][endpoint] == nil {
			errorByService[service][endpoint] = make(map[string]int)
		}

		errorByService[service][endpoint][message]++
	}

	for service, endpoints := range errorByService {
		fmt.Println("SERVICE:", service)
		outputFile.WriteString("\nService: " + service + "\n")
		for endpoint, errors := range endpoints {
			fmt.Println(" Endpoint:", endpoint)
			outputFile.WriteString(" Endpoint: " + endpoint + "\n")
			for msg, total := range errors {
				line := fmt.Sprintf("  - %s : %d\n", msg, total)
				outputFile.WriteString(line)
				fmt.Println("  -", msg, ":", total)
				fmt.Println("----------------------------")
			}
		}
	}
}
