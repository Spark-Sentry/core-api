package influxdb

import (
	"context"
	"fmt"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// TrendlogsQueryOptions is an example struct to hold parameter values
// that will be injected into a Flux script.
type TrendlogsQueryOptions struct {
	Bucket       string
	TimeStart    string
	TimeStop     string
	IdParameters []string
}

// ClientInfluxDBClient defines the interface for interacting with InfluxDB
type ClientInfluxDBClient interface {
	WritePoint(measurement string, tags map[string]string, fields map[string]interface{}, timestamp time.Time) error
	Query(ctx context.Context, query string) (*api.QueryTableResult, error)
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
		panic("Environment variables INFLUXDB_TOKEN, INFLUXDB_ORG, and INFLUXDB_BUCKET must be set")
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
func (c *Client) Query(ctx context.Context, query string) (*api.QueryTableResult, error) {
	result, err := c.queryAPI.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	return result, nil
}

// Close shuts down the InfluxDB client
func (c *Client) Close() {
	c.client.Close()
}
