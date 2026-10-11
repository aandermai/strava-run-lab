package activity

import "time"

type Activity struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Distance float64   `json:"distance"`
	Duration int       `json:"duration"`
	Date     time.Time `json:"date"`
}
