package client

import (
	"encoding/json"
)

type SensorType struct {
	SensorType    int32  `json:"sensor_type"`
	Manufacturer  string `json:"manufacturer"`
	ProductName   string `json:"product_name"`
	ProductNumber string `json:"product_number"`
	Category      string `json:"category"`
}

type SensorCatalog struct {
	SensorTypes []SensorType `json:"sensor_types"`
}

func (w *weatherlink) SensorCatalog() ([]SensorType, error) {
	result, err := w.get("sensor-catalog", nil, true)
	if err != nil {
		return []SensorType{}, err
	}
	var data SensorCatalog
	err = json.Unmarshal(result, &data)
	if err != nil {
		return []SensorType{}, err
	}
	return data.SensorTypes, nil
}
