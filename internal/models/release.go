package models

import "time"

type Release struct {
  ID          int `gorm:"primarykey;column:id;not null;<-:false"`
  Name        string `gorm:"varchar(255);column:name"`
  Description string `gorm:"text;column:description"`
	Data        string `gorm:"jsonb;column:data"`
	CreatedAt   time.Time `gorm:"timestamptz;column:createdAt;<-:false"`
	UpdatedAt   time.Time `gorm:"timestamptz;column:updatedAt;<-:false"`
}

func (Release) TableName() string {
	return "public.track_releases"
}