package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateCustomerHandler
// @Summary      Endpoint update customer
// @Description  update customer
// @Param customerId path int true "Customer ID"
// @Param Customer body model.UpdateCustomer true "It is a customer updatable info."
// @Tags         Customer
// @Produce json
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 406
// @Router       /api/v1/customer/{customerId} [put]
func UpdateCustomerHandler(log logger.Logger, updateCustomerCmd *command.UpdateCustomer) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "customerId")
		if err != nil {
			log.WithFields(logger.Fields{"exception": err}).Error("invalid path param")
			return
		}
		var request model.UpdateCustomer
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, dto.NewErrorMessage("invalid json body customer"))
			return
		}
		err = updateCustomerCmd.Do(c.Request.Context(), id, request)
		if err != nil {
			switch err.(type) {
			case exception.CustomerNotFoundErr:
				writeJsonErrorMessage(c, http.StatusNotFound, err)
			case exception.CustomerCannotUpdate, exception.CustomerAlreadyExistError:
				writeJsonErrorMessage(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught exception when update customer", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
