package dto

// ParameterCreateDTO is a struct for creating a new parameter
type ParameterCreateDTO struct {
	ID         int    `json:"id"`         // Unique identifier
	Name       string `json:"name"`       // Name of the parameter (e.g. "EVAC-6 AIR FLOW")
	HostDevice int    `json:"hostDevice"` // Identifier of the host device (7 digits)
	Device     int    `json:"device"`     // Identifier of the device (7 digits)
	Log        int64  `json:"log"`        // Log number (10 digits)
	Point      string `json:"point"`      // Point (e.g. "AV-8084")
	Unit       string `json:"unit"`       // Unit of measurement (e.g. "CFM")
}

// ParameterUpdateDTO is a struct for updating an existing parameter
type ParameterUpdateDTO struct {
	ID         int    `json:"id"`         // Unique identifier
	Name       string `json:"name"`       // Name of the parameter (e.g. "EVAC-6 AIR FLOW")
	HostDevice int    `json:"hostDevice"` // Identifier of the host device (7 digits)
	Device     int    `json:"device"`     // Identifier of the device (7 digits)
	Log        int64  `json:"log"`        // Log number (10 digits)
	Point      string `json:"point"`      // Point (e.g. "AV-8084")
	Unit       string `json:"unit"`       // Unit of measurement (e.g. "CFM")
}
