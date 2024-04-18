package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaqubmw/web-sales-app-golang/delivery/middleware"
	"github.com/yaqubmw/web-sales-app-golang/model"
	"github.com/yaqubmw/web-sales-app-golang/usecase"
)

type SaleController struct {
	router *gin.Engine
	saleUC usecase.SaleUsecase
}

func (s *SaleController) Input(c *gin.Context) {
	var sale model.Sale
	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := s.saleUC.Create(sale)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	salesRes := map[string]any{
		"id":      sale.Id,
		"user_id": sale.UserId,
		"tanggal_transaksi": sale.TanggalTransaksi,
		"jenis":   sale.Jenis,
		"nominal": sale.Nominal,
	}

	c.JSON(http.StatusCreated, salesRes)
}

func NewSaleController(router *gin.Engine, saleUC usecase.SaleUsecase) *SaleController {
	controller := SaleController{
		router: router,
		saleUC: saleUC,
	}

	rg := router.Group("/api/sales")
	rg.POST("/input", middleware.AuthMiddleware() , controller.Input)

	return &controller
}
