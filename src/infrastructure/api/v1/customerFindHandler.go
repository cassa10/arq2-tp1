package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FindCustomerHandler
// @Summary      Endpoint find customer
// @Description  find customer
// @Param customerId path int true "Customer ID"
// @Tags         Customer
// @Produce json
// @Success 200
// @Success 400
// @Failure 404
// @Router       /api/v1/customer/{customerId} [get]
func FindCustomerHandler(log logger.Logger, findCustomerQuery *query.FindCustomerById) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "customerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessage(c, http.StatusBadRequest, err)
			return
		}
		customer, err := findCustomerQuery.Do(c.Request.Context(), id)
		if err != nil {
			switch err.(type) {
			case exception.CustomerNotFoundErr:
				writeJsonErrorMessage(c, http.StatusNotFound, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when find customer", err)
			}
			return
		}
		c.JSON(http.StatusOK, customer)
	}
}
