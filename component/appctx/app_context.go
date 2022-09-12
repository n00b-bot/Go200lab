package appctx

import (
	"food/component/tokenprovider"
	"food/component/uploadprovider"
	"food/pubsub"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMysqlConnection() *gorm.DB
	GetUploadProvider() uploadprovider.Provider
	GetTokenProvider() tokenprovider.Provider
	GetPubSub() pubsub.PubSub
}

type appCtx struct {
	db       *gorm.DB
	provider uploadprovider.Provider
	token    tokenprovider.Provider
	pb       pubsub.PubSub
}

func NewAppCtx(db *gorm.DB, provider uploadprovider.Provider, token tokenprovider.Provider, pb pubsub.PubSub) *appCtx {
	return &appCtx{
		db:       db,
		provider: provider,
		token:    token,
		pb:       pb,
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

func (a appCtx) GetPubSub() pubsub.PubSub {
	return a.pb
}
