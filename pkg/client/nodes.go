package client

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Node struct {
	NodeId           int32   `json:"node_id"`
	NodeName         string  `json:"node_name"`
	RegistrationDate int32   `json:"registration_date"`
	DeviceId         int32   `json:"device_id"`
	DeviceIdHex      string  `json:"device_id_hex"`
	FirmwareVersion  int     `json:"firmware_version"`
	Active           bool    `json:"active"`
	StationId        int32   `json:"station_id"`
	StationName      string  `json:"station_name"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Elevation        float64 `json:"elevation"` //int
}

type Nodes struct {
	Nodes       []Node `json:"nodes"` //Stations
	GeneratedAd int    `json:"generated_ad"`
}

func (w *weatherlink) Nodes() ([]Node, error) {
	result, err := w.get("nodes", nil, true)
	if err != nil {
		return []Node{}, err
	}
	var data Nodes
	err = json.Unmarshal(result, &data)
	if err != nil {
		return []Node{}, err
	}
	return data.Nodes, nil
}

func (w *weatherlink) Node(nodeId int32) (Node, error) {
	nodeIdString := strconv.Itoa(int(nodeId))
	path := fmt.Sprintf("nodes/%s", nodeIdString)
	result, err := w.get(path, map[string]string{"node-ids": nodeIdString}, true)
	if err != nil {
		return Node{}, err
	}
	var data Nodes
	err = json.Unmarshal(result, &data)
	if err != nil {
		return Node{}, err
	}
	return data.Nodes[0], nil
}
