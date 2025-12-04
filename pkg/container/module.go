package container

import "gorm.io/gorm"

type Module interface {
	Register(db *gorm.DB, c *AppContainer)
}
