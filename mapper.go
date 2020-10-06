package appstore_sdk

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
)

type Mapper struct {
}

func (rp *Mapper) Init() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = '\t'
		return r
	})
}
