package handlers

import (
	"bytes"
	"fmt"
	"go-rest-api-kafka/internal/models"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlansHandler struct {
	db *gorm.DB
}

func NewPlansHandler(db *gorm.DB) *PlansHandler {
    return &PlansHandler{db: db}
}

func (h *PlansHandler) GetAllPlans(c *gin.Context) {
    var plans []models.Plan
    
    result := h.db.Find(&plans)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to fetch plans",
        })
        return
    }

    log.Println(plans)

    c.JSON(http.StatusOK, plans)
}

func (h *PlansHandler) CreatePlan(c *gin.Context) {
    // Логируем входящий запрос
    body, _ := io.ReadAll(c.Request.Body)
    c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
    log.Printf("Incoming request body: %s", string(body))

    var req models.PlanRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Request validation error: %v", err)
        
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
            "details": err.Error(),
            "expected_format": models.PlanRequest{
                ID:        uuid.UUID{},
                System:    "Track",
                EventDate: time.Now(),
                EventType: "CREATE",
                SuitCode:  "plan",
                SpaceCode: "AAA",
                Payload: models.PlanPayload{
                    Code:            "PLAN-001",
                    Status:          "active",
                    SourceCreatedAt: time.Now(),
                    SourceUpdatedAt: time.Now(),
                },
            },
        })
        return
    }

    fmt.Println(req);

    plan := models.Plan{
        UUID:            req.ID,
        Code:            req.Payload.Code,
        SuitCode:        req.SuitCode,
        System:          req.System,
        Status:          req.Payload.Status,
        SourceCreatedAt: req.Payload.SourceCreatedAt,
        SourceUpdatedAt: req.Payload.SourceUpdatedAt,
        EventDate:       req.EventDate,
    }

    result := h.db.Create(&plan)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to create plan",
        })
        return
    }

    c.JSON(http.StatusCreated, plan)
}