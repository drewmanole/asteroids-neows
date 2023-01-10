package interfaces

import "github.com/drewmanole/asteroids-neows/internal/entity"

type Asteroids interface {
	Get(dates ...string) ([]entity.AsteroidsReport, error)
	Create(reports ...entity.AsteroidsReport) error
	// TODO: Нужно ли делать метод Update?
}
