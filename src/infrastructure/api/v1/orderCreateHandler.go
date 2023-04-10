package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/domain/usecase"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOrderHandler
// @Summary      Endpoint create order
// @Description  create order
// @Param Order body dto.OrderCreateReq true "It is a order creation request."
// @Tags         Order
// @Produce json
// @Success 200 {object} dto.IdResponse
// @Failure 400
// @Failure 404
// @Failure 406
// @Router       /api/v1/order [post]
func CreateOrderHandler(log model.Logger, createOrder *usecase.CreateOrder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.OrderCreateReq
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, dto.NewErrorMessageComplete("invalid json body order", err.Error()))
			return
		}
		if err := req.Validate(); err != nil {
			writeJsonErrorMessage(c, http.StatusBadRequest, err)
			return
		}
		orderId, err := createOrder.Do(c.Request.Context(), req.CustomerId, req.ProductId, req.DeliveryDate, req.DeliveryAddress)
		if err != nil {
			switch err.(type) {
			case exception.CustomerNotFoundErr, exception.ProductNotFoundErr:
				writeJsonErrorMessage(c, http.StatusNotFound, err)
			case exception.ProductWithNoStock:
				writeJsonErrorMessage(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when create order", err)
			}
			return
		}
		c.JSON(http.StatusOK, dto.NewIdResponse(orderId))
	}
}
