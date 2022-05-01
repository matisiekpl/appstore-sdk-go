package appstore

import (
	"bytes"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
)

//UnmarshalCSV raw data to structures
func UnmarshalCSV(in []byte, out interface{}) error {
	r := NewCSVReader(bytes.NewReader(in))
	return gocsv.UnmarshalCSV(r, out)
}

//NewCSVReader Create new CSV reader for unmarshaler
func NewCSVReader(in io.Reader) gocsv.CSVReader {
	r := csv.NewReader(in)
	r.LazyQuotes = true
	r.Comma = '\t'
	return r
}

//func NewLineSkipDecoder(r io.Reader, LinesToSkip int) (gocsv.SimpleDecoder, error) {
//	reader := csv.NewReader(r)
//	reader.LazyQuotes = true
//	reader.Comma = '\t'
//	reader.FieldsPerRecord = -1
//	for i := 0; i < LinesToSkip; i++ {
//		if _, err := reader.Read(); err != nil {
//			return nil, err
//		}
//	}
//	reader.FieldsPerRecord = 0
//	return gocsv.NewSimpleDecoderFromCSVReader(reader), nil
//}
