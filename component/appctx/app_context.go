package appctx

import (
	"food/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMysqlConnection() *gorm.DB
	GetUploadProvider() uploadprovider.Provider
}

type appCtx struct {
	db       *gorm.DB
	provider uploadprovider.Provider
}

func NewAppCtx(db *gorm.DB, provider uploadprovider.Provider) *appCtx {
	return &appCtx{
		db:       db,
		provider: provider,
	}
}

func (a appCtx) GetUploadProvider() uploadprovider.Provider {
	return a.provider
}

func (a appCtx) GetMysqlConnection() *gorm.DB {
	return a.db
}
