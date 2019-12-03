# Example entity
Example entity for May CMF aims to show how to create custom modules and entity for MaySMF. 

# Usage

### Get code

```bash
$ go get github.com/MayCMF/example
```

in `app.go` file:
```go
import (
    ...
	"github.com/MayCMF/core/src/example"
	...
)

// BuildContainer Create a dependency injection container
func BuildContainer() (*dig.Container, func()) {
    ...
	err = example.InjectControllers(container)
    handleError(err)
    ...
}
```

in `migrate.go` file:
```go
import (
	...
    example "github.com/MayCMF/example/model/impl/gorm/entity"
    ...
)

// AutoMigrate - Automatic mapping data table
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
        ...
        new(example.Example),
        ...
	).Error
}
```

in `storage.go` file:
```go
import (
	...
    exampleIject "github.com/MayCMF/example"
    ...
)

// InitStore - Initialize storage
func InitStore(container *dig.Container) (func(), error) {
    ...
    i18nIject.InjectStarage(container)
    // Inject Example datatables to storage
    exampleIject.InjectStarage(container)
    ...
}
```

in `web.go` file:
```go
import (
	...
	exampleApi "github.com/MayCMF/cexample/routers/api"
    ...
)

// InitWeb - Initialize the web engine
func InitWeb(container *dig.Container) *gin.Engine {
    ...
	// Registration Example /api routing
	_ = exampleApi.RegisterRouter(app, container)
    ...
}
```

In the future versions of CMF will be combine in one file