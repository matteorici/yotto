package server

import (
	"fmt"
	"net/http"

	"uav-telemetry-visualizer/internal/handlers"
)

func Start() {

	http.Handle("/", http.FileServer(http.Dir("./web")))

	http.HandleFunc("/api/telemetry", handlers.Telemetry)

	fmt.Println("Listening on :8080")

	http.ListenAndServe(":8080", nil)

}
