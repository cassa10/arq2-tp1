package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FindSellerHandler
// @Summary      Endpoint find seller
// @Description  find seller
// @Param sellerId path int true "Seller ID"
// @Tags         Seller
// @Produce json
// @Success 200
// @Success 400
// @Failure 404
// @Router       /api/v1/seller/{sellerId} [get]
func FindSellerHandler(log model.Logger, findSellerByIdQuery *query.FindSellerById) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "sellerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessage(c, http.StatusBadRequest, err)
			return
		}
		seller, err := findSellerByIdQuery.Do(c.Request.Context(), id)
		if err != nil {
			switch err.(type) {
			case exception.SellerNotFoundErr:
				writeJsonErrorMessage(c, http.StatusNotFound, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when find seller", err)
			}
			return
		}
		c.JSON(http.StatusOK, seller)
	}
}
