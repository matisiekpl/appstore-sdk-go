package appstore

import (
	"bytes"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
	"strings"
)

//UnmarshalCSV unmarshal raw data to structures
func UnmarshalCSV(in []byte, out interface{}) error {
	r := NewCSVReader(bytes.NewReader(in))
	return gocsv.UnmarshalCSV(r, out)
}

//UnmarshalCSVWithFilterLines unmarshal raw data to structures with filter lines
func UnmarshalCSVWithFilterLines(in []byte, out interface{}) error {
	decoder, err := NewLineSkipDecoder(bytes.NewReader(in))
	if err != nil {
		return err
	}
	return gocsv.UnmarshalDecoder(decoder, out)
}

//NewCSVReader Create new CSV reader for unmarshaler
func NewCSVReader(in io.Reader) gocsv.CSVReader {
	r := csv.NewReader(in)
	r.LazyQuotes = true
	r.Comma = '\t'
	return r
}

//NewLineSkipDecoder
func NewLineSkipDecoder(r io.Reader) (gocsv.SimpleDecoder, error) {
	reader := csv.NewReader(r)
	reader.LazyQuotes = true
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	writer.Comma = reader.Comma

	for {
		row, err := reader.Read()
		if err != nil {
			return nil, err
		}
		if err == io.EOF || strings.Contains(row[0], "Total_") {
			break
		}
		err = writer.Write(row)
		if err != nil {
			return nil, err
		}
	}
	reader.FieldsPerRecord = 0

	writer.Flush()

	rf := bytes.NewReader(buf.Bytes())
	readerFiltered := csv.NewReader(rf)
	readerFiltered.LazyQuotes = reader.LazyQuotes
	readerFiltered.Comma = reader.Comma
	return gocsv.NewSimpleDecoderFromCSVReader(readerFiltered), nil
}
