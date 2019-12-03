package example

import (
	"github.com/MayCMF/example/controllers"
	"github.com/MayCMF/example/controllers/implement"
	"github.com/MayCMF/example/model"
	imodel "github.com/MayCMF/example/model/impl/gorm/model"
	"go.uber.org/dig"
)

// Inject - injection controllers implementation
func InjectControllers(container *dig.Container) error {
	_ = container.Provide(implement.NewExample)
	_ = container.Provide(func(b *implement.Example) controllers.IExample { return b })
	return nil
}

// Inject - Injection of gorm
func InjectStarage(container *dig.Container) error {
	_ = container.Provide(imodel.NewExample)
	_ = container.Provide(func(m *imodel.Example) model.IExample { return m })
	return nil
}
