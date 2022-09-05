package appctx

import (
	"food/component/tokenprovider"
	"food/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMysqlConnection() *gorm.DB
	GetUploadProvider() uploadprovider.Provider
	GetTokenProvider() tokenprovider.Provider
}

type appCtx struct {
	db       *gorm.DB
	provider uploadprovider.Provider
	token    tokenprovider.Provider
}

func NewAppCtx(db *gorm.DB, provider uploadprovider.Provider, token tokenprovider.Provider) *appCtx {
	return &appCtx{
		db:       db,
		provider: provider,
		token:    token,
	}
}

func (a appCtx) GetUploadProvider() uploadprovider.Provider {
	return a.provider
}

func (a appCtx) GetMysqlConnection() *gorm.DB {
	return a.db
}

func (a appCtx) GetTokenProvider() tokenprovider.Provider {
	return a.token
}
