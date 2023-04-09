package v1

import (
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProductHandler
// @Summary      Endpoint create product
// @Description  create product
// @Param sellerId path int true "Seller ID"
// @Param Product body dto.ProductCreateReq true "It is a product creation request."
// @Tags         Product
// @Produce json
// @Success 200 {object} dto.IdResponse
// @Failure 400
// @Failure 404
// @Router       /api/v1/seller/{sellerId}/product [post]
func CreateProductHandler(log logger.Logger, createProductCmd *command.CreateProduct) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := parsePathParamPositiveIntId(c, "sellerId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessage(c, http.StatusBadRequest, err)
			return
		}
		var request dto.ProductCreateReq
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, dto.NewErrorMessageComplete("invalid json body product", err.Error()))
			return
		}
		productId, err := createProductCmd.Do(c.Request.Context(), request.MapToModel(sellerId))
		if err != nil {
			switch err.(type) {
			case exception.SellerNotFoundErr:
				writeJsonErrorMessage(c, http.StatusNotFound, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when create product", err)
			}
			return
		}
		c.JSON(http.StatusOK, dto.NewIdResponse(productId))
	}
}
