package service

import (
	"context"
	"gin_mall_tmp/dao"
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/e"
	"gin_mall_tmp/pkg/util"
	"gin_mall_tmp/serializer"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` // 前端验证
}

func (service UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if service.Key == "" || len(service.Key) != 16 { //密钥长度必须为16
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}
	// 10000---> 密文存储 对称加密操作
	util.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Status:   model.Active,
		Avatar:   "avatar.jpg",
		Money:    util.Encrypt.AesEncoding("10000"), //初始金额的加密
	}
	//密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	if err = userDao.CreateUser(&user); err != nil {
		code = e.Error
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
