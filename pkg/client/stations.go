package client

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Station struct {
	StationId           int32   `json:"station_id"`
	StationName         string  `json:"station_name"`
	GatewayId           int32   `json:"gateway_id"`
	GatewayIdHex        string  `json:"gateway_id_hex"`
	ProductNumber       string  `json:"product_number"`
	Username            string  `json:"username"`
	UserEmail           string  `json:"user_email"`
	CompanyName         string  `json:"company_name"`
	Active              bool    `json:"active"`
	Private             bool    `json:"private"`
	RecordingInterval   int32   `json:"recording_interval"` //UploadInterval
	FirmwareVersion     string  `json:"firmware_version"`   //int32
	Imei                string  `json:"imei"`
	Meid                string  `json:"meid"`
	RegisteredDate      int32   `json:"registered_date"`
	SubscriptionEndDate int32   `json:"subscription_end_date"`
	TimeZone            string  `json:"time_zone"`
	City                string  `json:"city"`
	Region              string  `json:"region"`
	Country             string  `json:"country"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Elevation           float64 `json:"elevation"` //int
}

type Stations struct {
	Stations    []Station `json:"stations"`
	GeneratedAd int       `json:"generated_ad"`
}

func (w *weatherlink) Stations() ([]Station, error) {
	result, err := w.get("stations", nil, true)
	if err != nil {
		return []Station{}, err
	}
	var data Stations
	err = json.Unmarshal(result, &data)
	if err != nil {
		return []Station{}, err
	}
	return data.Stations, nil
}

func (w *weatherlink) Station(stationId int32) (Station, error) {
	stationIdString := strconv.Itoa(int(stationId))
	path := fmt.Sprintf("stations/%s", stationIdString)
	result, err := w.get(path, map[string]string{"station-ids": stationIdString}, true)
	if err != nil {
		return Station{}, err
	}
	var data Stations
	err = json.Unmarshal(result, &data)
	if err != nil {
		return Station{}, err
	}
	return data.Stations[0], nil
}
