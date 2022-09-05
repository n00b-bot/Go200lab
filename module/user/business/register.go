package userbusiness

import (
	"context"
	"food/common"
	usermodel "food/module/user/model"
)

type RegisterUserStorage interface {
	Create(context context.Context, data *usermodel.UserCreate) error
	FindUser(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type registerUserBusiness struct {
	store  RegisterUserStorage
	hasher Hasher
}

func NewRegisterUserBusiness(store RegisterUserStorage, hasher Hasher) *registerUserBusiness {
	return &registerUserBusiness{
		store:  store,
		hasher: hasher,
	}
}

func (r *registerUserBusiness) Register(context context.Context, data *usermodel.UserCreate) error {
	user, _ := r.store.FindUser(context, map[string]interface{}{"email": data.Email})
	if user != nil {
		return common.ErrEmailExist(nil)
	}
	salt := common.GenSalt(50)
	data.Salt = salt
	data.Password = r.hasher.Hash(data.Password + salt)
	data.Role = "user"
	if err := r.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
