package models

import (
	"dga-cp-lab-kafka/app/entities"
	"github.com/gin-gonic/gin"
)

type NotifyRequest struct {
	Message string `json:"message"`
}

func (req *NotifyRequest) Parse(c *gin.Context) (*NotifyRequest, error) {

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func (req *NotifyRequest) ToEntity() *entities.Message {
	return &entities.Message{
		Value: req.Message,
	}
}
