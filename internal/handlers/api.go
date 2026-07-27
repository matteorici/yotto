package handlers

import (
	"encoding/json"
	"net/http"

	"uav-telemetry-visualizer/internal/services"
)

func Telemetry(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(

		services.Data(),

	)

}
