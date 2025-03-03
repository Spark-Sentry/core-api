package services

import (
	"context"
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/influxdb"
	"core-api/internal/infrastructure/repository"
	"fmt"
	"strconv"
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

// EfficiencyMeasureData represents a single data point returned by the service.
type EfficiencyMeasureData struct {
	Time                time.Time `json:"time"`
	Value               float64   `json:"value"`
	IdEfficiencyMeasure string    `json:"id_efficiency_measure"`
	Mesh                string    `json:"mesh"`
}

// SavingsService is responsible for retrieving savings data from InfluxDB.
type SavingsService struct {
	influxClient   influxdb.ClientInfluxDBClient
	regressionRepo repository.RegressionRepository
}

// RegressionSavingsData represents a single computed data point for the regression savings.
type RegressionSavingsData struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

// NewSavingsService creates and returns a new SavingsService.
func NewSavingsService(client influxdb.ClientInfluxDBClient, regressionRepo repository.RegressionRepository) *SavingsService {
	return &SavingsService{influxClient: client, regressionRepo: regressionRepo}
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

// RetrieveSavingsForMultipleMeasures retrieves savings data for multiple efficiency measures.
func (s *SavingsService) RetrieveSavingsForMultipleMeasures(
	ctx context.Context,
	params dto.EfficiencyMeasureByMeasurementParams,
) ([]EfficiencyMeasureData, error) {
	bucket := "computed_data" // This can be set via configuration if needed.
	fluxQuery := s.buildMultipleMeasuresFluxQuery(bucket, params)

	queryResult, err := s.influxClient.Query(ctx, fluxQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query InfluxDB: %w", err)
	}
	defer queryResult.Close()

	var results []EfficiencyMeasureData
	for queryResult.Next() {
		record := queryResult.Record()
		value, ok := record.Value().(float64)
		if !ok {
			continue
		}
		measureID, _ := record.ValueByKey("id_efficiency_measure").(string)
		results = append(results, EfficiencyMeasureData{
			Time:                record.Time(),
			Value:               value,
			IdEfficiencyMeasure: measureID,
			Mesh:                params.Mesh,
		})
	}
	if queryResult.Err() != nil {
		return nil, fmt.Errorf("error reading query results: %w", queryResult.Err())
	}

	return results, nil
}

// buildMultipleMeasuresFluxQuery builds a Flux query for multiple efficiency measure IDs.
func (s *SavingsService) buildMultipleMeasuresFluxQuery(
	bucket string,
	params dto.EfficiencyMeasureByMeasurementParams,
) string {
	var fluxMeasureIDs []string
	for _, mID := range params.EfficiencyMeasureID {
		fluxMeasureIDs = append(fluxMeasureIDs, fmt.Sprintf("%q", mID))
	}

	queryBuilder := &strings.Builder{}
	fmt.Fprintf(queryBuilder, `
from(bucket: %q)
  |> range(start: %s, stop: %s)
  |> filter(fn: (r) => r["_measurement"] == "hourly_gas_savings")
  |> filter(fn: (r) => r["_field"] == "value")
  |> filter(fn: (r) => contains(value: r["id_efficiency_measure"], set: [%s]))
`, bucket, params.TimeStart, params.TimeStop, strings.Join(fluxMeasureIDs, ", "))

	// Adjust the aggregation based on the mesh parameter.
	switch params.Mesh {
	case "hourly":
		queryBuilder.WriteString(`
  |> aggregateWindow(every: 1h, fn: sum, createEmpty: false)
`)
	case "daily":
		queryBuilder.WriteString(`
  |> aggregateWindow(every: 1d, fn: sum, createEmpty: false)
`)
	case "monthly":
		queryBuilder.WriteString(`
  |> aggregateWindow(every: 1mo, fn: sum, createEmpty: false)
`)
	case "annually":
		queryBuilder.WriteString(`
  |> aggregateWindow(every: 1y, fn: sum, createEmpty: false)
`)
	case "total":
		queryBuilder.WriteString(`
  |> sum()
`)
	}

	queryBuilder.WriteString(`  |> yield(name: "result")`)
	return queryBuilder.String()
}

// RetrieveSavingsByRegression computes the savings for a project based on a regression.
// For each month between start and stop, it calculates the sum over all coefficients:
//
//	savings = sum(coefficient.Value * independentVariableValue)
//
// Independent variable values are retrieved from different sources:
//   - If source == "input", they are obtained from the relational DB.
//   - If source == "computed", they are obtained by calling an external endpoint (simulated here).
//
// The result is aggregated according to the mesh parameter ("monthly", "annually", or "total").
func (s *SavingsService) RetrieveSavingsByRegression(params dto.ProjectByRegressionParams) ([]RegressionSavingsData, error) {
	// Convert RegressionId from string to uint
	regIDUint, err := strconv.ParseUint(params.RegressionId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid regressionId: %w", err)
	}

	// Retrieve regression details from the relational DB.
	// (Assume s.regressionRepo.FindRegressionByID accepts a uint)
	regression, err := s.regressionRepo.FindRegressionByID(uint(regIDUint))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve regression: %w", err)
	}

	// Parse start and stop times.
	startTime, err := time.Parse(time.RFC3339, params.TimeStart)
	if err != nil {
		return nil, fmt.Errorf("invalid start time: %w", err)
	}
	stopTime, err := time.Parse(time.RFC3339, params.TimeStop)
	if err != nil {
		return nil, fmt.Errorf("invalid stop time: %w", err)
	}

	// Generate monthly intervals between start and stop.
	monthlyIntervals := generateMonthlyIntervals(startTime, stopTime)

	// For each month, compute the savings.
	var monthlyResults []RegressionSavingsData
	for _, monthStart := range monthlyIntervals {
		monthEnd := monthStart.AddDate(0, 1, 0)
		if monthEnd.After(stopTime) {
			monthEnd = stopTime
		}

		var monthSum float64
		// Iterate over each coefficient in the regression.
		for _, coef := range regression.Coefficients {
			var indVarValue float64
			// Use pointer to coef: &coef
			if coef.IndependantVariable != nil && coef.IndependantVariable.Source == "input" {
				indVarValue, err = getInputIndependantVariable(&coef, monthStart, monthEnd)
				if err != nil {
					return nil, fmt.Errorf("failed to retrieve input independant variable: %w", err)
				}
			} else {
				indVarValue, err = s.getComputedIndependantVariable(&coef, monthStart, monthEnd)
				if err != nil {
					return nil, fmt.Errorf("failed to retrieve computed independant variable: %w", err)
				}
			}
			monthSum += coef.Value * indVarValue
		}

		monthlyResults = append(monthlyResults, RegressionSavingsData{
			Time:  monthStart,
			Value: monthSum,
		})
	}

	// Aggregate results based on mesh parameter.
	switch strings.ToLower(params.Mesh) {
	case "annually":
		annualResults := aggregateByYear(monthlyResults)
		return annualResults, nil
	case "total":
		var total float64
		for _, r := range monthlyResults {
			total += r.Value
		}
		return []RegressionSavingsData{{Time: startTime, Value: total}}, nil
	default: // "monthly"
		return monthlyResults, nil
	}
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

// generateMonthlyIntervals creates a slice of time.Time marking the start of each month between start and stop.
func generateMonthlyIntervals(start, stop time.Time) []time.Time {
	var intervals []time.Time
	t := time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, start.Location())
	for t.Before(stop) {
		intervals = append(intervals, t)
		t = t.AddDate(0, 1, 0)
	}
	return intervals
}

// getInputIndependantVariable retrieves the independent variable value for an "input" source
// from the relational DB. It assumes that coef.IndependantVariable is preloaded.
func getInputIndependantVariable(coef *entities.Coefficient, start, stop time.Time) (float64, error) {
	iv := coef.IndependantVariable
	if iv == nil {
		return 0, fmt.Errorf("no independent variable linked to coefficient")
	}
	ivStart, err := time.Parse(time.RFC3339, iv.DateStart)
	if err != nil {
		return 0, fmt.Errorf("invalid independent variable start date: %w", err)
	}
	ivStop, err := time.Parse(time.RFC3339, iv.DateStop)
	if err != nil {
		return 0, fmt.Errorf("invalid independent variable stop date: %w", err)
	}
	// Check if the requested period [start, stop] is fully within the independent variable's period.
	if start.Before(ivStart) || stop.After(ivStop) {
		// Selon votre logique métier, vous pouvez retourner 0 ou une erreur.
		return 0, nil
	}
	return iv.Value, nil
}

// getComputedIndependantVariable retrieves the computed independent variable value from InfluxDB.
// It builds a Flux query to aggregate (here, sum) the values from a measurement (e.g., "computed_independent")
// where the tag "id_independent" matches the ID of the independent variable.
func (s *SavingsService) getComputedIndependantVariable(coef *entities.Coefficient, start, stop time.Time) (float64, error) {
	if coef.IndependantVariable == nil {
		return 0, fmt.Errorf("no independent variable linked to coefficient")
	}
	// Build the Flux query.
	fluxQuery := fmt.Sprintf(`
from(bucket: "computed_data")
  |> range(start: %s, stop: %s)
  |> filter(fn: (r) => r["_measurement"] == "computed_independent")
  |> filter(fn: (r) => r["_field"] == "value")
  |> filter(fn: (r) => r["id_independent"] == %q)
  |> sum()
  |> yield(name: "sum")
`, start.Format(time.RFC3339), stop.Format(time.RFC3339), strconv.Itoa(int(coef.IndependantVariable.ID)))

	// Execute the Flux query.
	result, err := s.influxClient.Query(context.Background(), fluxQuery)
	if err != nil {
		return 0, fmt.Errorf("failed to query InfluxDB: %w", err)
	}
	defer result.Close()

	var sumValue float64
	for result.Next() {
		if v, ok := result.Record().Value().(float64); ok {
			sumValue = v
		}
	}
	if result.Err() != nil {
		return 0, fmt.Errorf("error reading query result: %w", result.Err())
	}
	return sumValue, nil
}

// aggregateByYear aggregates monthly savings data into annual totals.
func aggregateByYear(monthlyResults []RegressionSavingsData) []RegressionSavingsData {
	annualMap := make(map[int]float64)
	for _, r := range monthlyResults {
		year := r.Time.Year()
		annualMap[year] += r.Value
	}
	var annualResults []RegressionSavingsData
	for year, value := range annualMap {
		annualResults = append(annualResults, RegressionSavingsData{
			Time:  time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
			Value: value,
		})
	}
	return annualResults
}
