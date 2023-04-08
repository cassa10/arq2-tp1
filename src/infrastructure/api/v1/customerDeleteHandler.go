package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteCustomerHandler
// @Summary      Endpoint delete customer
// @Description  delete customer by id
// @Param customerId path int true "Customer ID"
// @Tags         Customer
// @Produce json
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 406
// @Router       /api/v1/customer/{customerId} [delete]
func DeleteCustomerHandler(log logger.Logger, deleteCustomerCmd *command.DeleteCustomer) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "customerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			return
		}
		err = deleteCustomerCmd.Do(c.Request.Context(), id)
		if err != nil {
			switch err.(type) {
			case model.CustomerNotFoundErr:
				writeJsonErrorMessage(c, http.StatusNotFound, err)
			case model.CustomerCannotDelete:
				writeJsonErrorMessage(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when delete customer", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
