package dto

// TrendlogsParams represents the request parameters for retrieving trend log data.
type TrendlogsParams struct {
	Bucket       string   `json:"bucket"`       // InfluxDB bucket name
	TimeStart    string   `json:"timeStart"`    // Start time for the query
	TimeStop     string   `json:"timeStop"`     // Stop time for the query
	IdParameters []string `json:"idParameters"` // List of parameter IDs to filter on
}
