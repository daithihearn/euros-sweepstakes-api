package score

import (
	"euros-sweepstakes-api/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	S ServiceI
}

// Get @Summary Get scores
// @Description Get all scores
// @Tags scores
// @ID get-scores
// @Produce json
// @Success 200 {array} Score
// @Failure 500 {object} api.ErrorResponse
// @Router /scores [get]
func (h *Handler) Get(c *gin.Context) {
	scores, err := h.S.GetScores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, scores)
}
