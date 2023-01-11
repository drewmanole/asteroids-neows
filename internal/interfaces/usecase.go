package interfaces

import "github.com/drewmanole/asteroids-neows/internal/entity"

type Asteroids interface {
	Create(reports []entity.AsteroidsReport) error
	Get(dates []string) ([]entity.AsteroidsReport, error)
	Clear() error
}
