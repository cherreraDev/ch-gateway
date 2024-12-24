package dependencycontainer

import "gorm.io/gorm"

type Container struct {
}

func NewContainer(db *gorm.DB) Container {
	return Container{}
}
