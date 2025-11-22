package repositories

import (
	"database/sql"
	"log"

	"github.com/Lucas4lves/subplots/models"
)

type PlotRepository struct {
	Driver *sql.DB
}

func NewPlotRepository(driver *sql.DB) *PlotRepository {

	if driver == nil {
		log.Fatal("ERROR: unable to create a Plot Repository")
	}

	return &PlotRepository{
		Driver: driver,
	}
}

func (pr *PlotRepository) Insert(plot *models.Plot) error {
	tx, err := pr.Driver.Begin()

	if err != nil {
		log.Println(err)
		return err
	}

	defer tx.Rollback()

	sql := `
		insert into plots (title, story, status) 
		values ($1,$2,$3);
	`

	stmt, err := tx.Prepare(sql)

	if err != nil {

		log.Println(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(plot.Title, plot.Story, plot.Status)

	if err != nil {

		log.Println(err)
		tx.Rollback()

		return err
	}

	tx.Commit()

	return nil
}

func (pr *PlotRepository) SelectAll() ([]*models.Plot, error) {

	var results []*models.Plot

	sql := "select * from plots"
	rows, err := pr.Driver.Query(sql)

	if err != nil {
		log.Printf("ERROR: %s", err.Error())

		return nil, err
	}

	for rows.Next() {
		var plotVessel *models.Plot = &models.Plot{}

		// TODO: Lucas 'I could try to use Plot's instance constructor here instead of calling a validate func
		rows.Scan(&plotVessel.ID, &plotVessel.Title, &plotVessel.Story, &plotVessel.Status, &plotVessel.CreatedAt, &plotVessel.UpdatedAt)

		results = append(results, plotVessel)
	}

	return results, nil

}
