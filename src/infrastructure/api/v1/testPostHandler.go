package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TestPostHandler
// @Summary      Endpoint test post
// @Description  test post endpoint with a product
// @Param Product body model.Product true "Representa un producto."
// @Tags         Test handler
// @Produce json
// @Success 200
// @Failure 400
// @Router       /api/v1/test [post]
func TestPostHandler(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			log.WithFields(logger.Fields{"error": err.Error()}).Error("couldn't bind body")
			c.Status(http.StatusBadRequest)
			return
		}

		log.Infof("succesfull unmarshall %s", req)
		c.JSON(http.StatusOK, req)
	}
}
