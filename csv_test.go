package main

import (
	"encoding/csv"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"os"
	"testing"
)

func TestCSV(t *testing.T) {
	f, err := os.Create(export.GetExcelFullPath() + "test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	data := [][]string{
		{"1", "test1", "test1-1"},
		{"2", "test2", "test2-1"},
		{"3", "test3", "test3-1"},
	}

	w.WriteAll(data)
}
