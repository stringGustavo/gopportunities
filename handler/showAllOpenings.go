package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stringGustavo/gopportunities.git/schemas"
)

// @BasePath /api/v1

// @Summary Show openings
// @Description Show all existing job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ShowAllOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ShowAllOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "show-all-openings", openings)
}
