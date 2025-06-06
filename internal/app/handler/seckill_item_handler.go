package handler

import (
	"OrderEZ/internal/app/service"
)

type SeckillGoodHandler struct {
	seckillGoodService *service.SeckillGoodService
}

func NewSeckillGoodHandler(seckillGoodService *service.SeckillGoodService) *SeckillGoodHandler {
	return &SeckillGoodHandler{seckillGoodService: seckillGoodService}
}
