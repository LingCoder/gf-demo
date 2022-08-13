package packed

import (
	"embed"
)

//go:embed *
var f embed.FS

func GetCharts() ([]byte, error) {
	res, err := f.ReadFile("charts.json")
	return res, err
}

func GetSQL() ([]byte, error) {
	res, err := f.ReadFile("sql/demo.sql")
	return res, err
}
