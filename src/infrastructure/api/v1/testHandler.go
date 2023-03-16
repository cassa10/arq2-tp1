package v1

import (
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TestHandler
// @Summary      Endpoint test always return ok
// @Description  test endpoint
// @Tags         Test handler
// @Produce json
// @Success 200
// @Router       /api/v1/test [get]
func TestHandler(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Infof("succesfull test")
		c.String(http.StatusOK, "test ok")
	}
}
