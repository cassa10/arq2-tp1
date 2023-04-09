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

// UpdateSellerHandler
// @Summary      Endpoint update seller
// @Description  update seller
// @Param sellerId path int true "Seller ID"
// @Param Seller body model.UpdateSeller true "It is a seller updatable info."
// @Tags         Seller
// @Produce json
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 406
// @Router       /api/v1/seller/{sellerId} [put]
func UpdateSellerHandler(log logger.Logger, updateSellerCmd *command.UpdateSeller) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parsePathParamPositiveIntId(c, "sellerId")
		if err != nil {
			log.WithFields(logger.Fields{"exception": err}).Error("invalid path param")
			return
		}
		var request model.UpdateSeller
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, dto.NewErrorMessage("invalid json body seller"))
			return
		}
		err = updateSellerCmd.Do(c.Request.Context(), id, request)
		if err != nil {
			switch err.(type) {
			case exception.SellerNotFoundErr:
				writeJsonErrorMessage(c, http.StatusNotFound, err)
			case exception.SellerCannotUpdate, exception.SellerAlreadyExistError:
				writeJsonErrorMessage(c, http.StatusNotAcceptable, err)
			default:
				defaultInternalServerError(log, c, "uncaught exception when update seller", err)
			}
			return
		}
		c.Status(http.StatusNoContent)
	}
}
