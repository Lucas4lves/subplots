package container

import "database/sql"

type DependencyContainer struct {
	Driver *sql.DB
}

func NewDependencyContainer(driver *sql.DB) *DependencyContainer {
	return &DependencyContainer{
		Driver: driver,
	}
}
