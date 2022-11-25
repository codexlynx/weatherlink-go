package client

import (
	"encoding/json"
	"fmt"
	"io"
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
	url, err := w.url("sensor-activity", nil)
	if err != nil {
		return []SensorActivity{}, err
	}
	resp, err := w.client.Get(url)
	result, err := io.ReadAll(resp.Body)
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
	url, err := w.url(path, map[string]string{"sensor-ids": lsIdString})
	if err != nil {
		return SensorActivity{}, err
	}
	resp, err := w.client.Get(url)
	result, err := io.ReadAll(resp.Body)
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
