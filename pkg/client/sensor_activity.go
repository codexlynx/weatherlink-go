package client

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type SensorActivity struct {
	LsId             int32 `json:"lsid"`
	TimeReceived     int32 `json:"time_received"`
	TimeRecorded     int32 `json:"time_recorded"`
	TimeLoopReceived int32 `json:"time_loop_received"`
}

type SensorActivities struct {
	SensorActivity []SensorActivity `json:"sensor_activity"`
	GeneratedAd    int              `json:"generated_ad"`
}

func (w *weatherlink) SensorActivities() ([]SensorActivity, error) {
	result, err := w.get("sensor-activity", nil, true)
	if err != nil {
		return []SensorActivity{}, err
	}
	var data SensorActivities
	err = json.Unmarshal(result, &data)
	if err != nil {
		return []SensorActivity{}, err
	}
	return data.SensorActivity, nil
}

func (w *weatherlink) SensorActivity(lsId int32) (SensorActivity, error) {
	lsIdString := strconv.Itoa(int(lsId))
	path := fmt.Sprintf("sensor-activity/%s", lsIdString)
	result, err := w.get(path, map[string]string{"sensor-ids": lsIdString}, true)
	if err != nil {
		return SensorActivity{}, err
	}
	var data SensorActivities
	err = json.Unmarshal(result, &data)
	if err != nil {
		return SensorActivity{}, err
	}
	return data.SensorActivity[0], nil
}
