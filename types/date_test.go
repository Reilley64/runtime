package types

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate_MarshalJSON(t *testing.T) {
	testDate := time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)
	b := struct {
		DateField Date `json:"date"`
	}{
		DateField: Date{testDate},
	}
	jsonBytes, err := json.Marshal(b)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"date":"2019-04-01"}`, string(jsonBytes))
}

func TestDate_UnmarshalJSON(t *testing.T) {
	testDate := time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)
	jsonStr := `{"date":"2019-04-01"}`
	b := struct {
		DateField Date `json:"date"`
	}{}
	err := json.Unmarshal([]byte(jsonStr), &b)
	assert.NoError(t, err)
	assert.Equal(t, testDate, b.DateField.Time)
}

func TestDate_MarshalXML(t *testing.T) {
	testDate := time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)
	b := struct {
		DateField Date `xml:"date"`
	}{
		DateField: Date{testDate},
	}
	xmlBytes, err := xml.Marshal(b)
	assert.NoError(t, err)
	assert.Equal(t, `<date>2019-04-01</date>`, string(xmlBytes))
}

func TestDate_UnmarshalXML(t *testing.T) {
	testDate := time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)
	xmlStr := `<date>2019-04-01</date>`
	b := struct {
		DateField Date `xml:"date"`
	}{}
	err := xml.Unmarshal([]byte(xmlStr), &b)
	assert.NoError(t, err)
	assert.Equal(t, testDate, b.DateField.Time)
}

func TestDate_Stringer(t *testing.T) {
	t.Run("nil date", func(t *testing.T) {
		var d *Date
		assert.Equal(t, "<nil>", fmt.Sprintf("%v", d))
	})

	t.Run("ptr date", func(t *testing.T) {
		d := &Date{
			Time: time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC),
		}
		assert.Equal(t, "2019-04-01", fmt.Sprintf("%v", d))
	})

	t.Run("value date", func(t *testing.T) {
		d := Date{
			Time: time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC),
		}
		assert.Equal(t, "2019-04-01", fmt.Sprintf("%v", d))
	})
}

func TestDate_UnmarshalText(t *testing.T) {
	testDate := time.Date(2022, 6, 14, 0, 0, 0, 0, time.UTC)
	value := []byte("2022-06-14")

	date := Date{}
	err := date.UnmarshalText(value)

	assert.NoError(t, err)
	assert.Equal(t, testDate, date.Time)
}
