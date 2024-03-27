package dto

// CollectData defines the structure for incoming data collection requests.
type CollectData struct {
	DeviceID  string  `json:"device_id"` // Unique identifier for the device
	Timestamp string  `json:"timestamp"` // Timestamp of the data
	Value     float64 `json:"value"`     // Collected value
}
