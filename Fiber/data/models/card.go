package models

import "database/sql"

type Card struct {
	ID            string
	Title         string
	Description   string
	HasDate       bool
	Date          string
	Frequency     int
	FrequencyUnit string
	TagID         sql.NullString
	ParentID      sql.NullString
}
