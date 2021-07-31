package main

import (
	"strings"
	"time"
)

type Member struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Nickname  string `json:"nickname"`
	Birthday  Date   `json:"birthday"`
}

type Date time.Time

func (date *Date) UnmarshalJSON(b []byte) error {
	v := strings.Trim(string(b), `"`)
	if v == "" || v == "null" {
		return nil
	}
	d, err := time.Parse("2006-01-02", v)
	if err != nil {
		return err
	}
	*date = Date(d)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}
