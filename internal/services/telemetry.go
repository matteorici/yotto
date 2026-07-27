package services

import (
	"uav-telemetry-visualizer/internal/parser"
)

func Data() [][]string {

	return parser.Load()

}
