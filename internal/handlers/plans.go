package handlers

import (
	"go-rest-api-kafka/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
    var req models.PlanRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request body",
        })
        return
    }

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