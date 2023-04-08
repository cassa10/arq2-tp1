package v1

import (
	"fmt"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func defaultInternalServerError(log logger.Logger, ginContext *gin.Context, additionalLogInfo string, err error) {
	log.WithFields(logger.Fields{"error": err}).Error(additionalLogInfo)
	ginContext.JSON(http.StatusInternalServerError, dto.NewErrorMessage("internal server error"))
}

// parsePathParamPositiveIntId writes in c *gin.Context with Bad Request with message error when cannot parse properly 'paramKey' as a positive int64
func parsePathParamPositiveIntId(c *gin.Context, paramKey string) (int64, error) {
	_idParam, _ := c.Params.Get(paramKey)
	id, err := strconv.ParseInt(_idParam, 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, dto.NewErrorMessage(fmt.Sprintf("param '%s' is required as positive int64", paramKey)))
		return 0, fmt.Errorf("invalid path param %s as positive int64", paramKey)
	}
	return id, err
}

func writeJsonErrorMessage(c *gin.Context, status int, err error) {
	c.JSON(status, dto.NewErrorMessage(err.Error()))
}
