package services

import (
	"context"
	"core-api/internal/app/dto"
	"core-api/internal/infrastructure/influxdb"
	"fmt"
	"time"
)

// TrendlogsService is responsible for fetching timeseries data from InfluxDB based on user-defined parameters.
type TrendlogsService struct {
	influxClient influxdb.ClientInfluxDBClient
}

// NewTrendlogsService creates a new instance of TrendlogsService.
func NewTrendlogsService(client influxdb.ClientInfluxDBClient) *TrendlogsService {
	return &TrendlogsService{
		influxClient: client,
	}
}

// DataPoint represents a single point of data in the trendlogs
type DataPoint struct {
	Time        time.Time `json:"time"`
	Value       float64   `json:"value,omitempty"`
	Parameter   string    `json:"parameter,omitempty"`
	Equipment   string    `json:"equipment,omitempty"`
	Table       string    `json:"table,omitempty"`
	Mean        string    `json:"mean,omitempty"`
	Measurement string    `json:"measurement,omitempty"`
	Field       string    `json:"field,omitempty"`
	Start       string    `json:"start,omitempty"`
	Stop        string    `json:"stop,omitempty"`
	Device      string    `json:"device,omitempty"`
	HostDevice  string    `json:"host_device,omitempty"`
	Name        string    `json:"name,omitempty"`
}

// formatArray formats a slice of strings as a Flux array (["value1", "value2"])
func formatArray(array []string) string {
	formatted := "["
	for i, val := range array {
		formatted += fmt.Sprintf("%q", val)
		if i < len(array)-1 {
			formatted += ", "
		}
	}
	formatted += "]"
	return formatted
}

// RetrieveTrendlogs executes a Flux query against InfluxDB using the provided parameters.
func (s *TrendlogsService) RetrieveTrendlogs(ctx context.Context, params dto.TrendlogsParams) ([]DataPoint, error) {
	fluxQuery := fmt.Sprintf(`
 from(bucket: %q)
    |> range(start: %s, stop: %s)
    |> filter(fn: (r) => r["_measurement"] == "hourly_averages")
    |> filter(fn: (r) => r["_field"] == "value")
    |> filter(fn: (r) => contains(value: r["id_parameter"], set: %s))
    |> aggregateWindow(every: 1h, fn: mean, createEmpty: false)
    |> yield(name: "mean")
`, params.Bucket, params.TimeStart, params.TimeStop, formatArray(params.IdParameters))

	queryResult, err := s.influxClient.Query(ctx, fluxQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query InfluxDB: %w", err)
	}
	defer queryResult.Close()

	var datapoints []DataPoint

	for queryResult.Next() {
		record := queryResult.Record()
		value, ok := record.Value().(float64)
		if !ok {
			continue
		}

		parameter, _ := record.ValueByKey("id_parameter").(string)
		equipment, _ := record.ValueByKey("id_equipment").(string)
		table, _ := record.ValueByKey("table").(string)
		mean, _ := record.ValueByKey("mean").(string)
		measurement, _ := record.ValueByKey("_measurement").(string)
		field, _ := record.ValueByKey("_field").(string)
		start, _ := record.ValueByKey("_start").(string)
		stop, _ := record.ValueByKey("_stop").(string)
		device, _ := record.ValueByKey("device").(string)
		hostDevice, _ := record.ValueByKey("host_device").(string)
		name, _ := record.ValueByKey("name").(string)

		datapoints = append(datapoints, DataPoint{
			Time:        record.Time(),
			Value:       value,
			Parameter:   parameter,
			Equipment:   equipment,
			Table:       table,
			Mean:        mean,
			Measurement: measurement,
			Field:       field,
			Start:       start,
			Stop:        stop,
			Device:      device,
			HostDevice:  hostDevice,
			Name:        name,
		})
	}

	if queryResult.Err() != nil {
		return nil, fmt.Errorf("error reading query results: %w", queryResult.Err())
	}

	return datapoints, nil
}
