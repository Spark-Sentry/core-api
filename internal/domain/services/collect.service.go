package services

type CollectService struct{}

func NewCollectService() *CollectService {
	return &CollectService{}
}

// CollectData handle the data from Building Management service and save it to influxDB
// We will need to call the infrastructure method of 3rd party library to save the data to influxDB
func (c *CollectService) CollectData(data []byte) error {
	var err error
	// TODO: implement the method
	// 1. parse the data from Building Management service
	// 2. save the data to influxDB
	// 3. return error if any
	return err

}
