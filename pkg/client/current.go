package client

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type CurrentSensor struct {
	LsId              int32                    `json:"lsid"`
	SensorType        int32                    `json:"sensor_type"`
	DataStructureType int32                    `json:"data_structure_type"`
	Data              []map[string]interface{} `json:"data"` // This field is very variable, impossible to type.
}

type Current struct {
	StationId   int32
	Sensors     []CurrentSensor `json:"sensors"`
	GeneratedAd int             `json:"generated_ad"`
}

func (w *weatherlink) Current(stationId int32) ([]CurrentSensor, error) {
	stationIdString := strconv.Itoa(int(stationId))
	path := fmt.Sprintf("current/%s", stationIdString)
	url, err := w.url(path, map[string]string{"station-id": stationIdString}, false)
	if err != nil {
		return []CurrentSensor{}, err
	}
	resp, err := w.client.Get(url)
	result, err := io.ReadAll(resp.Body)
	var data Current
	err = json.Unmarshal(result, &data)
	if err != nil {
		return []CurrentSensor{}, err
	}
	return data.Sensors, nil
}
