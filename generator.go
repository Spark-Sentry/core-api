package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Measurement represents a single measurement entry
type Measurement struct {
	Value     float64 `json:"value"`
	Timestamp string  `json:"timestamp"`
}

// Data holds all measurements and metadata for JSON output
type Data struct {
	Measurements []Measurement `json:"measurements"`
	Name         string        `json:"name"`
	HostDevice   int           `json:"hostDevice"`
	Device       int           `json:"device"`
	Log          float64       `json:"log"`
	Point        string        `json:"point"`
	IDEquipment  int           `json:"id_equipment"`
}

// generateData generates measurement data based on given parameters
func generateData(name string, hostDevice int, device int, log float64, point string, idEquipment int, weeks int, minValue, maxValue float64) Data {
	// Use the current time as the endpoint and generate data going backwards
	endTime := time.Now().UTC()
	totalMeasurements := weeks * 7 * 24 // Calculate total number of measurements for the given weeks

	measurements := make([]Measurement, 0, totalMeasurements)
	for i := 0; i < totalMeasurements; i++ {
		// Generate a random measurement value within the specified range
		value := minValue + rand.Float64()*(maxValue-minValue)

		// Calculate the timestamp for each consecutive hour going backwards
		timestamp := endTime.Add(-time.Duration(i) * time.Hour).Format(time.RFC3339)

		// Append the measurement to the list
		measurements = append(measurements, Measurement{
			Value:     value,
			Timestamp: timestamp,
		})
	}

	// Reverse the slice to have the oldest data first
	for i, j := 0, len(measurements)-1; i < j; i, j = i+1, j-1 {
		measurements[i], measurements[j] = measurements[j], measurements[i]
	}

	// Return the final data structure with measurements and metadata
	return Data{
		Measurements: measurements,
		Name:         name,
		HostDevice:   hostDevice,
		Device:       device,
		Log:          log,
		Point:        point,
		IDEquipment:  idEquipment,
	}
}

func main() {
	// Define input parameters
	name := flag.String("name", "VENT-7 TEMPERATURE", "Parameter name")
	hostDevice := flag.Int("hostDevice", 1002, "Host device ID")
	device := flag.Int("device", 502, "Device ID")
	log := flag.Float64("log", 1.03, "Log value")
	point := flag.String("point", "P2", "Measurement point")
	idEquipment := flag.Int("idEquipment", 43, "Equipment ID")
	weeks := flag.Int("weeks", 3, "Number of weeks of data")
	minValue := flag.Float64("minValue", 21.5, "Minimum value for measurements")
	maxValue := flag.Float64("maxValue", 23.5, "Maximum value for measurements")
	output := flag.String("output", "data.json", "Output JSON file name")

	flag.Parse()

	// Generate data based on input parameters
	data := generateData(*name, *hostDevice, *device, *log, *point, *idEquipment, *weeks, *minValue, *maxValue)

	// Save data to JSON file
	file, err := os.Create(*output)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Printf("JSON file successfully generated: %s\n", *output)
}
