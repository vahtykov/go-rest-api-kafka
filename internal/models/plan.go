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
    Code            string    `json:"code" binding:"required"`
    Status          string    `json:"status" binding:"required"`
    SourceCreatedAt time.Time `json:"createdAt" binding:"required"`
    SourceUpdatedAt time.Time `json:"updatedAt" binding:"required"`
}

type PlanRequest struct {
    ID        uuid.UUID   `json:"id" binding:"required"`
    System    string      `json:"system" binding:"required"`
    EventDate time.Time   `json:"eventDate" binding:"required"`
    EventType string      `json:"eventType" binding:"required,oneof=CREATE UPDATE DELETE"`
    SuitCode  string      `json:"suitCode" binding:"required"`
    SpaceCode string      `json:"spaceCode" binding:"required"`
    Payload   PlanPayload `json:"payload" binding:"required"`
}

func (Plan) TableName() string {
	return "public.plans"
}