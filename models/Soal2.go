package models

type DataSensor struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	RoomArea    string  `json:"roomArea"`
	ID          int64   `json:"id"`
	Timestamp   int64   `json:"timestamp"`
}

type DataArray struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

type DataSensorRespone struct {
	Temperature TempStruct
	Humidity    HumidStruct
}

type TempStruct struct {
	TemperatureMean   float64 `json:"temperatureMean"`
	TemperatureModus  float64 `json:"temperatureModus"`
	TemperatureMedian float64 `json:"temperatureMedian"`
}

type HumidStruct struct {
	HumidityMean   float64 `json:"HumidityMean"`
	HumidityModus  float64 `json:"HumidityModus"`
	HumidityMedian float64 `json:"HumidityMedian"`
}
