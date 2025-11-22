package container

import (
	"database/sql"

	"github.com/Lucas4lves/subplots/controllers"
	"github.com/Lucas4lves/subplots/repositories"
	"github.com/Lucas4lves/subplots/services"
)

type DependencyContainer struct {
	Driver         *sql.DB
	PlotRepo       *repositories.PlotRepository
	PlotSvc        *services.PlotService
	PlotController *controllers.PlotController
}

func NewDependencyContainer(driver *sql.DB) *DependencyContainer {

	repo := *repositories.NewPlotRepository(driver)
	svc := *services.NewPlotService(&repo)

	ctrl := *controllers.NewPlotController(&svc)

	return &DependencyContainer{
		Driver:         driver,
		PlotRepo:       &repo,
		PlotSvc:        &svc,
		PlotController: &ctrl,
	}
}
