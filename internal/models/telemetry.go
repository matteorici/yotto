package models

type Telemetry struct {

	Time string `json:"time"`

	Latitude float64 `json:"latitude"`

	Longitude float64 `json:"longitude"`

	Speed float64 `json:"speed"`

	Altitude float64 `json:"altitude"`

	Battery int `json:"battery"`
}
