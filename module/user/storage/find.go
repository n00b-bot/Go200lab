package userstorage

import (
	"context"
	usermodel "food/module/user/model"
)

func (s *sqlStore) FindUser(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.UserCreate{}.TableName())
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var user usermodel.User
	if err := s.db.Where(conditions).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
