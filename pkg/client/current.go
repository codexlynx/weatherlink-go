package client

import (
	"encoding/json"
	"fmt"
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
	result, err := w.get(path, map[string]string{"station-id": stationIdString}, false)
	if err != nil {
		return []CurrentSensor{}, err
	}
	var data Current
	err = json.Unmarshal(result, &data)
	if err != nil {
		return []CurrentSensor{}, err
	}
	return data.Sensors, nil
}
