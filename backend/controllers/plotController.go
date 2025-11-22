package controllers

import (
	"net/http"

	"github.com/Lucas4lves/subplots/models"
	"github.com/Lucas4lves/subplots/services"
	"github.com/gin-gonic/gin"
)

type PlotController struct {
	Service *services.PlotService
}

func NewPlotController(svc *services.PlotService) *PlotController {
	return &PlotController{
		Service: svc,
	}
}

func (pc *PlotController) GetAll(ctx *gin.Context) {

	data, err := pc.Service.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"data":       data,
	})

}

func (pc *PlotController) CreatePlot(ctx *gin.Context) {
	var newPlot *models.Plot

	err := ctx.ShouldBindBodyWithJSON(&newPlot)

	if err != nil {
		code := http.StatusBadRequest
		ctx.JSON(code, gin.H{
			"statusCode": code,
			"message":    err.Error(),
		})

		return
	}

	err = pc.Service.Create(newPlot)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	code := http.StatusCreated

	ctx.JSON(code, gin.H{
		"statusCode": code,
		"message":    "Plot successfully created!",
	})

}
