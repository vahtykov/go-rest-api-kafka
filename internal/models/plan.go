package models

import (
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID              int       `gorm:"primarykey;column:id;not null;<-:false"`
	UUID            uuid.UUID `gorm:"type:uuid;column:uuid;not null"`
	Code            string    `gorm:"column:code;not null"`
	SuitCode        string    `gorm:"column:suit_code;not null"`
	System          string    `gorm:"column:system;not null"`
	Status          string    `gorm:"column:status;not null"`
	SourceCreatedAt time.Time `gorm:"column:source_createdAt;not null"`
	SourceUpdatedAt time.Time `gorm:"column:source_updatedAt;not null"`
	CreatedAt       time.Time `gorm:"column:createdAt;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"column:updatedAt;default:CURRENT_TIMESTAMP"`
	EventDate       time.Time `gorm:"column:event_date;default:CURRENT_TIMESTAMP"`
}

type PlanPayload struct {
    Code            string    `json:"code"`
    Status          string    `json:"status"`
    SourceCreatedAt time.Time `json:"createdAt"`
    SourceUpdatedAt time.Time `json:"updatedAt"`
}

type PlanRequest struct {
    ID        uuid.UUID   `json:"id"`
    System    string      `json:"system"`
    EventDate time.Time   `json:"eventDate"`
    EventType string      `json:"eventType"`
    SuitCode  string      `json:"suitCode"`
    SpaceCode string      `json:"spaceCode"`
    Payload   PlanPayload `json:"payload"`
}

func (Plan) TableName() string {
	return "public.plans"
}