package parser

import (
	"encoding/csv"
	"os"
)

func Load() [][]string {

	file, _ := os.Open("data/telemetry.csv")

	defer file.Close()

	reader := csv.NewReader(file)

	rows, _ := reader.ReadAll()

	return rows
}
