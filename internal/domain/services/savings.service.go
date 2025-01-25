package services

import (
	"context"
	"core-api/internal/app/dto"
	"core-api/internal/infrastructure/influxdb"
	"fmt"
	"strings"
	"time"
)

// SavingsDataPoint represents a single data point for savings results.
type SavingsDataPoint struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value,omitempty"`
	Unit  string    `json:"unit,omitempty"`
	Mesh  string    `json:"mesh,omitempty"`
	// Add any other fields you want to return
}

// SavingsService is responsible for retrieving savings data from InfluxDB.
type SavingsService struct {
	influxClient influxdb.ClientInfluxDBClient
}

// NewSavingsService creates and returns a new SavingsService.
func NewSavingsService(client influxdb.ClientInfluxDBClient) *SavingsService {
	return &SavingsService{influxClient: client}
}

// RetrieveSavings builds and executes a Flux query against InfluxDB to get savings data.
func (s *SavingsService) RetrieveSavings(ctx context.Context, params dto.SavingsParams) ([]SavingsDataPoint, error) {
	// Validate or sanitize the input parameters if needed.
	// For example: check if "mesh" is valid, check if "units" are valid, etc.
	if err := validateMesh(params.Mesh); err != nil {
		return nil, fmt.Errorf("mesh validation failed: %v", err)
	}
	if err := validateUnits(params.Units); err != nil {
		return nil, fmt.Errorf("units validation failed: %v", err)
	}

	// Build the Flux query dynamically based on the mesh, time range, etc.
	fluxQuery := s.buildSavingsFluxQuery(params)

	// Execute the query
	queryResult, err := s.influxClient.Query(ctx, fluxQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query InfluxDB: %w", err)
	}
	defer queryResult.Close()

	var results []SavingsDataPoint
	for queryResult.Next() {
		record := queryResult.Record()
		value, ok := record.Value().(float64)
		if !ok {
			// If it's not a float64, skip or handle differently
			continue
		}

		// Extract other tag/field values if needed
		unit, _ := record.ValueByKey("unit").(string)

		results = append(results, SavingsDataPoint{
			Time:  record.Time(),
			Value: value,
			Unit:  unit,
			Mesh:  params.Mesh,
		})
	}

	if queryResult.Err() != nil {
		return nil, fmt.Errorf("error reading query results: %w", queryResult.Err())
	}

	return results, nil
}

// buildSavingsFluxQuery constructs a Flux query string based on user parameters.
func (s *SavingsService) buildSavingsFluxQuery(params dto.SavingsParams) string {
	bucket := "computed_data"

	queryBuilder := &strings.Builder{}
	fmt.Fprintf(queryBuilder, `
from(bucket: %q)
  |> range(start: %s, stop: %s)
  |> filter(fn: (r) => r["_measurement"] == "hourly_gas_savings")
  |> filter(fn: (r) => r["_field"] == "value")
  |> filter(fn: (r) => r["id_efficiency_measure"] == %q)
`, bucket, params.TimeStart, params.TimeStop, params.IdEfficiencyMeasure)

	if len(params.Units) > 0 {
		var fluxUnits []string
		for _, u := range params.Units {
			fluxUnits = append(fluxUnits, fmt.Sprintf("%q", u))
		}
		fmt.Fprintf(queryBuilder, `
  |> filter(fn: (r) => contains(value: r["unit"], set: [%s]))
`, strings.Join(fluxUnits, ", "))
	}

	switch params.Mesh {
	case "Hourly":
	case "Daily":
		queryBuilder.WriteString(`
  |> aggregateWindow(every: 1d, fn: sum, createEmpty: false)
`)
	case "Weekly":
		queryBuilder.WriteString(`
  |> aggregateWindow(every: 1w, fn: sum, createEmpty: false)
`)
	case "Monthly":
		queryBuilder.WriteString(`
  |> aggregateWindow(every: 1mo, fn: sum, createEmpty: false)
`)
	case "Annually":
		queryBuilder.WriteString(`
  |> aggregateWindow(every: 1y, fn: sum, createEmpty: false)
`)
	}

	queryBuilder.WriteString(`  |> yield(name: "result")`)
	return queryBuilder.String()
}

// validateUnits checks if the units array contains valid values.
func validateUnits(units []string) error {
	validUnits := map[string]bool{
		"m3":   true,
		"kWh":  true,
		"GJ":   true,
		"tGHH": true,
		"$":    true,
	}
	for _, u := range units {
		if !validUnits[u] {
			return fmt.Errorf("invalid unit: %s", u)
		}
	}
	return nil
}

// validateMesh checks if the mesh value is valid.
func validateMesh(mesh string) error {
	validMesh := map[string]bool{
		"Hourly":   true,
		"Daily":    true,
		"Weekly":   true,
		"Monthly":  true,
		"Annually": true,
	}
	if !validMesh[mesh] {
		return fmt.Errorf("invalid mesh: %s", mesh)
	}
	return nil
}
