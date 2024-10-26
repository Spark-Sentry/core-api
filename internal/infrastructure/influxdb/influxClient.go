package influxdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// ClientInfluxDBClient defines the interface for interacting with InfluxDB
type ClientInfluxDBClient interface {
	WritePoint(measurement string, tags map[string]string, fields map[string]interface{}, timestamp time.Time) error
	Query(query string) (*api.QueryTableResult, error)
	Close()
}

// Client implements the ClientInfluxDBClient interface
type Client struct {
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI
	org      string
	bucket   string
}

// NewClient initializes a new InfluxDB2 client with default settings
func NewClient() *Client {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := os.Getenv("INFLUXDB_URL")
	org := os.Getenv("INFLUXDB_ORG")
	bucket := os.Getenv("INFLUXDB_BUCKET")

	if token == "" || org == "" || bucket == "" {
		log.Fatal("Environment variables INFLUXDB_TOKEN, INFLUXDB_ORG, and INFLUXDB_BUCKET must be set")
	}
	if url == "" {
		url = "http://localhost:8086"
	}

	client := influxdb2.NewClient(url, token)
	return &Client{
		client:   client,
		writeAPI: client.WriteAPIBlocking(org, bucket),
		queryAPI: client.QueryAPI(org),
		org:      org,
		bucket:   bucket,
	}
}

// WritePoint writes a data point to InfluxDB
func (c *Client) WritePoint(measurement string, tags map[string]string, fields map[string]interface{}, timestamp time.Time) error {
	p := influxdb2.NewPoint(measurement, tags, fields, timestamp)
	return c.writeAPI.WritePoint(context.Background(), p)
}

// Query executes a Flux query and returns the result
func (c *Client) Query(query string) (*api.QueryTableResult, error) {
	result, err := c.queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	return result, nil
}

// Close shuts down the InfluxDB client
func (c *Client) Close() {
	c.client.Close()
}

// ExampleWrite demonstrates how to use the WritePoint method
func ExampleWrite() {
	client := NewClient()
	defer client.Close()

	tags := map[string]string{
		"host": "server01",
	}
	fields := map[string]interface{}{
		"cpu_usage": 0.65,
	}

	err := client.WritePoint("system", tags, fields, time.Now())
	if err != nil {
		log.Fatalf("Error writing point: %v", err)
	}
}

// ExampleQuery demonstrates how to use the Query method
func ExampleQuery() {
	client := NewClient()
	defer client.Close()

	query := `from(bucket:"my-bucket") |> range(start: -1h) |> filter(fn: (r) => r._measurement == "system")`
	result, err := client.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	// Process query result (this is just an example)
	for result.Next() {
		fmt.Printf("Time: %v, Value: %v\n", result.Record().Time(), result.Record().Value())
	}
	if result.Err() != nil {
		log.Fatalf("Query parsing error: %v", result.Err())
	}
}
