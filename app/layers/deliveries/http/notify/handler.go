package notify

import (
	notifyUseCase "dga-cp-lab-kafka/app/layers/usecases/notify"
	"github.com/gin-gonic/gin"
)

type handler struct {
	NotifyUseCase notifyUseCase.UseCase
}

func NewEndpointHTTPHandler(ginEngine *gin.Engine, notifyUseCase notifyUseCase.UseCase) {
	handler := handler{
		NotifyUseCase: notifyUseCase,
	}

	v1 := ginEngine.Group("/notify")
	{
		v1.POST("/new-message", handler.Notify)

	}
}
