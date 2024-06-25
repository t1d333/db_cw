package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t1d333/log_collector/internal/delivery/models"
	"github.com/t1d333/log_collector/internal/httputil"
	"github.com/t1d333/log_collector/internal/service"
	"github.com/t1d333/log_collector/pkg/logger"
)

type Delivery struct {
	logger  logger.Logger
	mux     *gin.RouterGroup
	service service.Service
}

func NewDelivery(mux *gin.RouterGroup, logger logger.Logger, serv service.Service) *Delivery {
	return &Delivery{
		mux:     mux,
		logger:  logger,
		service: serv,
	}
}

func (d *Delivery) RegisterHandlers() {
	d.mux.POST("/group", d.createGroup)
	d.mux.DELETE("/group/:name", d.deleteGroup)
}
