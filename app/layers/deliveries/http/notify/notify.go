package notify

import (
	"dga-cp-lab-kafka/app/layers/deliveries/http/notify/models"
	"dga-cp-lab-kafka/app/utils"
	"github.com/gin-gonic/gin"
)

func (h *handler) Notify(c *gin.Context) {
	req := new(models.NotifyRequest)
	if _, err := req.Parse(c); err != nil {

		utils.JSONErrorResponse(c, err, nil)
		return
	}

	h.NotifyUseCase.ProducerNewMessage(req.ToEntity())

	utils.JSONSuccessResponse(c, nil)
}
