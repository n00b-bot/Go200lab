package userstorage

import (
	"context"
	usermodel "food/module/user/model"
)

func (s *sqlStore) Create(context context.Context, data *usermodel.UserCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
