package models

type Card struct {
	ID            int
	Title         string
	Description   string
	HasDate       bool
	Date          string
	Frequency     int
	FrequencyUnit string
	TagID         int
	ParentID      int
}
