package facades

import "gorm.io/gorm"

var orm *gorm.DB

func MakeOrm(dbInstance *gorm.DB) {
	orm = dbInstance
}

func Orm() *gorm.DB {
	return orm
}
