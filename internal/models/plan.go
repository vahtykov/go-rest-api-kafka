package models

import (
	"gorm.io/gorm"
)

type Plan struct {
    gorm.Model `gorm:"table:track_plans"`
    Name        string `varchar:"name"`
    Description string `text:"description"`
    Data        string `jsonb:"data"`
    CreatedAt   string `timestamptz:"createdAt"`
    UpdatedAt   string `timestamptz:"updatedAt"`
}