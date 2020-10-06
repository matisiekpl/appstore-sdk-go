package appstore_sdk

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Types_CustomFloat64_MarshalJSONSuccess(t *testing.T) {
	c := CustomFloat64{}
	c.Float64 = 10.10
	result, err := c.MarshalJSON()
	expected, _ := json.Marshal(c.Float64)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_Types_CustomFloat64_MarshalJSONEmpty(t *testing.T) {
	c := CustomFloat64{}
	result, err := c.MarshalJSON()
	expected, _ := json.Marshal(c.Float64)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_Types_CustomTimestamp_UnmarshalCSVFilled(t *testing.T) {
	c := CustomTimestamp{}
	str := "2020-09-10 15:15:15"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, "2020-09-10 15:15:15", c.Value().Format(CustomTimestampFormatDefault))
}

func Test_Types_CustomTimestamp_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomTimestamp{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.True(t, c.Value().IsZero())
}

func Test_Types_CustomTimestamp_MarshalJSONSuccess(t *testing.T) {
	c := CustomTimestamp{}
	c.Timestamp = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`"2020-01-20 00:00:00"`), result)
	assert.Nil(t, err)
}

func Test_Types_CustomTimestamp_MarshalJSONEmpty(t *testing.T) {
	c := CustomTimestamp{}
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(``), result)
	assert.Nil(t, err)
}

func Test_Types_CustomDate_UnmarshalCSVFilled(t *testing.T) {
	c := CustomDate{}
	str := "2020-09-10"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, "2020-09-10", c.Value().Format(CustomDateFormatDefault))
}

func Test_Types_CustomDate_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomDate{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.True(t, c.Value().IsZero())
}

func Test_Types_CustomDate_MarshalJSONSuccess(t *testing.T) {
	c := CustomDate{}
	c.Date = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`"2020-01-20"`), result)
	assert.Nil(t, err)
}

func Test_Types_CustomDate_MarshalJSONEmpty(t *testing.T) {
	c := CustomDate{}
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(``), result)
	assert.Nil(t, err)
}

func Test_Types_CustomBoolean_UnmarshalCSVTrue(t *testing.T) {
	c := CustomBoolean{}
	str := "true"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, true, c.Value())
}

func Test_Types_CustomBoolean_UnmarshalCSVFalse(t *testing.T) {
	c := CustomBoolean{}
	str := "false"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, false, c.Value())
}

func Test_Types_CustomBoolean_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomBoolean{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, false, c.Value())
}

func Test_Types_CustomBoolean_MarshalJSONSuccess(t *testing.T) {
	c := CustomBoolean{}
	c.Boolean = true
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`true`), result)
	assert.Nil(t, err)
}

func Test_Types_CustomBoolean_MarshalJSONEmpty(t *testing.T) {
	c := CustomBoolean{}
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`false`), result)
	assert.Nil(t, err)
}
