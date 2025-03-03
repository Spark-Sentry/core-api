package dto

// TrendlogsParams represents the request parameters for retrieving trend log data.
type TrendlogsParams struct {
	Bucket       string   `json:"bucket"`
	TimeStart    string   `json:"timeStart"`
	TimeStop     string   `json:"timeStop"`
	IdParameters []string `json:"idParameters"`
	Mesh         string   `json:"mesh"` // New field: hourly, daily, monthly, or annually
}
