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

// CreateCustomerHandler
// @Summary      Endpoint create customer
// @Description  create customer
// @Param Customer body model.Customer true "It is a customer."
// @Tags         Customer
// @Produce json
// @Success 200 {object} dto.IdResponse
// @Failure 400
// @Failure 406
// @Router       /api/v1/customer [post]
func CreateCustomerHandler(log logger.Logger, createCustomerCmd *command.CreateCustomer) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.Customer
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, dto.NewErrorMessageComplete("invalid json body customer", err.Error()))
			return
		}
		customerId, err := createCustomerCmd.Do(c.Request.Context(), request)
		if err != nil {
			switch err.(type) {
			case exception.CustomerAlreadyExistError:
				writeJsonErrorMessage(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught exception when create customer", err)
			}
			return
		}
		c.JSON(http.StatusOK, dto.NewIdResponse(customerId))
	}
}
