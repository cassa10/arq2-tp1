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

// UpdateProductHandler
// @Summary      Endpoint update product
// @Description  update product
// @Param productId path int true "Product ID"
// @Param Product body model.UpdateProduct true "It is a product updatable info."
// @Tags         Product
// @Produce json
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 406
// @Router       /api/v1/seller/product/{productId} [put]
func UpdateProductHandler(log logger.Logger, updateProductCmd *command.UpdateProduct) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "productId")
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("invalid path param")
			writeJsonErrorMessage(c, http.StatusBadRequest, err)
			return
		}
		var request model.UpdateProduct
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, dto.NewErrorMessage("invalid json body product"))
			return
		}
		err = updateProductCmd.Do(c.Request.Context(), id, request)
		if err != nil {
			switch err.(type) {
			case exception.ProductNotFoundErr:
				writeJsonErrorMessage(c, http.StatusNotFound, err)
			case exception.ProductCannotUpdate:
				writeJsonErrorMessage(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught error when update product", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
