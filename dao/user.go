package dao

import (
	"context"
	"gin_mall_tmp/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(username string) (user *model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=？", username).Find(&user).Error
	if user == nil || err == gorm.ErrRecordNotFound {
		return nil, false, err
	}
	return user, true, nil
}
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}
