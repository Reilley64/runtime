package types

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

const DateFormat = "2006-01-02"

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(DateFormat))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

func (d Date) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return encoder.EncodeElement(d.Time.Format(DateFormat), start)
}

func (d *Date) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var dateStr string
	if err := decoder.DecodeElement(&dateStr, &start); err != nil {
		return err
	}
	parsed, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

func (d Date) String() string {
	return d.Time.Format(DateFormat)
}

func (d *Date) UnmarshalText(data []byte) error {
	parsed, err := time.Parse(DateFormat, string(data))
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}
