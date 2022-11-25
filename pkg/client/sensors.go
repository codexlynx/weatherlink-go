package client

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type Sensor struct {
	Lsid             int32   `json:"lsid"`
	SensorType       int32   `json:"sensor_type"`
	Category         string  `json:"category"`
	Manufacturer     string  `json:"manufacturer"`
	ProductName      string  `json:"product_name"`
	ProductNumber    string  `json:"product_number"`
	Active           bool    `json:"active"`
	CreatedDate      int32   `json:"created_date"`  //string
	ModifiedDate     int32   `json:"modified_date"` //string
	StationId        int32   `json:"station_id"`
	StationName      string  `json:"station_name"`
	ParentDeviceType string  `json:"parent_device_type"`
	ParentDeviceName string  `json:"parent_device_name"`
	ParentDeviceId   int32   `json:"parent_device_id"`
	ParentDeviceHex  string  `json:"parent_device_hex"`
	PortNumber       int32   `json:"port_number"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Elevation        float64 `json:"elevation"` //int
}

type Sensors struct {
	Sensors     []Sensor `json:"sensors"`
	GeneratedAd int      `json:"generated_ad"`
}

func (w *weatherlink) Sensors() ([]Sensor, error) {
	url, err := w.url("sensors", nil)
	if err != nil {
		return []Sensor{}, err
	}
	resp, err := w.client.Get(url)
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Sensor{}, err
	}
	var data Sensors
	err = json.Unmarshal(result, &data)
	if err != nil {
		return []Sensor{}, err
	}
	return data.Sensors, nil
}

func (w *weatherlink) Sensor(lsId int32) (Sensor, error) {
	lsIdString := strconv.Itoa(int(lsId))
	path := fmt.Sprintf("sensors/%s", lsIdString)
	url, err := w.url(path, map[string]string{"sensor-ids": lsIdString})
	if err != nil {
		return Sensor{}, err
	}
	resp, err := w.client.Get(url)
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return Sensor{}, err
	}
	var data Sensors
	err = json.Unmarshal(result, &data)
	if err != nil {
		return Sensor{}, err
	}
	return data.Sensors[0], nil
}
