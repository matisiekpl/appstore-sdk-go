package appstore

import (
	"bytes"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
)

//CSV unmarshaler
type CSV struct {
}

//Unmarshal raw data to Reports structures
func (c *CSV) Unmarshal(in []byte, out interface{}) error {
	r := c.NewCSVReader(bytes.NewReader(in))
	return gocsv.UnmarshalCSV(r, out)
}

//NewCSVReader Create new CSV reader for unmarshaler
func (c *CSV) NewCSVReader(in io.Reader) gocsv.CSVReader {
	r := csv.NewReader(in)
	r.LazyQuotes = true
	r.Comma = '\t'
	return r
}
