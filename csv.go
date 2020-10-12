package appstore_sdk

import (
	"bytes"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
)

type CSV struct {
}

func (c *CSV) Unmarshal(in []byte, out interface{}) error {
	r := c.NewCSVReader(bytes.NewReader(in))
	return gocsv.UnmarshalCSV(r, out)
}

func (c *CSV) NewCSVReader(in io.Reader) gocsv.CSVReader {
	r := csv.NewReader(in)
	r.LazyQuotes = true
	r.Comma = '\t'
	return r
}
