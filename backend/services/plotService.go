package services

import (
	"log"

	"github.com/Lucas4lves/subplots/models"
	"github.com/Lucas4lves/subplots/repositories"
)

type PlotService struct {
	Repository *repositories.PlotRepository
}

func NewPlotService(repo *repositories.PlotRepository) *PlotService {

	if repo == nil {
		log.Fatal("ERROR: unable to create a Plot Repository")
	}

	return &PlotService{
		Repository: repo,
	}
}

func (ps *PlotService) Create(plot *models.Plot) error {
	err := ps.Repository.Insert(plot)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (ps *PlotService) FindAll() ([]*models.Plot, error) {
	data, err := ps.Repository.SelectAll()

	if err != nil {
		log.Println("ERROR", err.Error())

		return nil, err
	}

	return data, nil

}
