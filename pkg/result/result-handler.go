package result

import (
	"euros-sweepstakes-api/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	S ServiceI
}

// Get @Summary Get result
// @Description Get result
// @Tags result
// @ID get-result
// @Produce json
// @Success 200 {object} Result
// @Failure 500 {object} api.ErrorResponse
// @Router /result [get]
func (h *Handler) Get(c *gin.Context) {
	result, err := h.S.GetResult()
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
