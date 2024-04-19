package healthcheck

import (
	"dga-cp-lab-kafka/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *handler) Health(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, utils.NewSuccessResponse(nil))
}
