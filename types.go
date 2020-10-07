package appstore_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const CustomTimestampFormatDefault = "2006-01-02 15:04:05"
const CustomDateFormatDefault = "2006-01-02"
const CustomDateFormatSlash = "01/02/2006"

type CustomInteger struct {
	Integer int
}

func (ci *CustomInteger) Value() int {
	return ci.Integer
}

func (ci *CustomInteger) MarshalJSON() ([]byte, error) {
	jsonData, err := json.Marshal(ci.Integer)
	if err != nil {
		return nil, errors.New("CustomInteger@MarshalJSON: " + err.Error())
	}
	return jsonData, err
}

func (ci *CustomInteger) UnmarshalCSV(csv string) error {
	if csv != "" {
		var err error
		ci.Integer, err = strconv.Atoi(csv)
		if err != nil {
			return fmt.Errorf("CustomInteger@UnmarshalCSV Parse int: %v", err)
		}
	}
	return nil
}

type CustomFloat64 struct {
	Float64 float64
}

func (cf *CustomFloat64) Value() float64 {
	return cf.Float64
}

func (cf *CustomFloat64) UnmarshalCSV(csv string) error {
	if csv != "" {
		var err error
		cf.Float64, err = strconv.ParseFloat(csv, 32)
		if err != nil {
			return fmt.Errorf("CustomFloat64@UnmarshalCSV Parse float: %v", err)
		}
	}
	return nil
}

func (cf *CustomFloat64) MarshalJSON() ([]byte, error) {
	jsonData, err := json.Marshal(cf.Float64)
	if err != nil {
		return nil, errors.New("CustomFloat64@MarshalJSON: " + err.Error())
	}
	return jsonData, err
}

type CustomTimestamp struct {
	Timestamp time.Time
}

func (ct *CustomTimestamp) Value() time.Time {
	return ct.Timestamp
}

func (ct *CustomTimestamp) UnmarshalCSV(csv string) error {
	if csv != "" {
		var err error
		ct.Timestamp, err = time.Parse(CustomTimestampFormatDefault, csv)
		if err != nil {
			return fmt.Errorf("CustomTimestamp@UnmarshalJSON ParseTime: %v", err)
		}
	}
	return nil
}

func (ct *CustomTimestamp) MarshalJSON() ([]byte, error) {
	if ct.Timestamp.IsZero() {
		return []byte(""), nil
	}
	formatted := ct.Timestamp.Format(CustomTimestampFormatDefault)
	jsonData, err := json.Marshal(formatted)
	if err != nil {
		return nil, errors.New("CustomTimestamp@MarshalJSON: " + err.Error())
	}
	return jsonData, err
}

type CustomDate struct {
	Date time.Time
}

func (ct *CustomDate) Value() time.Time {
	return ct.Date
}

func (ct *CustomDate) UnmarshalCSV(csv string) error {
	if csv == "" {
		return nil
	}

	var format string
	if strings.Contains(csv, `/`) {
		format = CustomDateFormatSlash
	} else {
		format = CustomDateFormatDefault
	}
	var err error
	ct.Date, err = time.Parse(format, csv)
	if err != nil {
		return fmt.Errorf("CustomDate@UnmarshalJSON ParseTime: %v", err)
	}
	return nil
}

func (ct *CustomDate) MarshalJSON() ([]byte, error) {
	if ct.Date.IsZero() {
		return []byte(""), nil
	}
	jsonData, err := json.Marshal(ct.Date.Format(CustomDateFormatDefault))
	if err != nil {
		return nil, errors.New("CustomDate@MarshalJSON: " + err.Error())
	}
	return jsonData, err
}

type CustomBoolean struct {
	Boolean bool
}

func (cb *CustomBoolean) Value() bool {
	return cb.Boolean
}

func (cb *CustomBoolean) UnmarshalCSV(csv string) error {
	switch strings.ToLower(csv) {
	case "false":
		cb.Boolean = false
	case "true":
		cb.Boolean = true
	}
	return nil
}

func (cb *CustomBoolean) MarshalJSON() ([]byte, error) {
	jsonData, err := json.Marshal(cb.Boolean)
	if err != nil {
		return nil, errors.New("CustomBoolean@MarshalJSON: " + err.Error())
	}
	return jsonData, err
}
