package userbusiness

import (
	"context"
	"food/common"
	"food/component/tokenprovider"
	usermodel "food/module/user/model"
)

type LoginStorage interface {
	FindUser(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	store  LoginStorage
	token  tokenprovider.Provider
	hasher Hasher
	expiry int
}

func NewLoginBusiness(store LoginStorage, token tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		store:  store,
		token:  token,
		hasher: hasher,
		expiry: expiry,
	}
}

func (l *loginBusiness) Login(context context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := l.store.FindUser(context, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, common.ErrLoginFail(err)
	}
	passHashed := l.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, common.ErrLoginFail(err)
	}
	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}
	token, err := l.token.Generate(payload, l.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return token, nil
}
