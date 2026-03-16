package main

import (
	"database/sql"

	"gorm.io/gorm"
)

type TimeReminder struct {
	gorm.Model
	Author    string
	ChannelID string
	Content   sql.NullString
	Hour      int
	Minute    int
}
