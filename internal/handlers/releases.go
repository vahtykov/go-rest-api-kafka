package handlers

import (
	"go-rest-api-kafka/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReleasesHandler struct {
	db *gorm.DB
}

func NewReleasesHandler(db *gorm.DB) *ReleasesHandler {
    return &ReleasesHandler{db: db}
}

func (h *ReleasesHandler) GetAllReleases(c *gin.Context) {
    var releases []models.Release
    
    result := h.db.Find(&releases)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to fetch releases",
        })
        return
    }

    log.Println(releases)

    c.JSON(http.StatusOK, releases)
}