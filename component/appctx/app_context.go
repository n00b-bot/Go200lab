package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMysqlConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppCtx(db *gorm.DB) *appCtx {
	return &appCtx{
		db: db,
	}
}

func (a appCtx) GetMysqlConnection() *gorm.DB {
	return a.db
}
