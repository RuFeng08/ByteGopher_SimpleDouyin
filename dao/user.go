package dao

import (
	"ByteGopher_SimpleDouyin/model"
)

type UserDao interface {
	GetUserModelByID(id int) (*model.UserModel, error)
	GetUserByName(username string) (*model.UserModel, error)
	AddUserModel(m *model.UserModel) error
	GetCommonUserByID(id int64) (*model.User, error)
}

type userDao struct{}

func NewUserDao() UserDao {
	return &userDao{}
}

func (dao *userDao) GetUserModelByID(id int) (*model.UserModel, error) {

	var m model.UserModel
	if err := MysqlDb.Table("user").Where("user_id = ?", id).Find(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *userDao) GetCommonUserByID(id int64) (*model.User, error) {
	Wg.Add(1)
	defer Wg.Wait()

	go CountUsers(id)

	var m model.User
	if err := MysqlDb.Table("user").Where("user_id = ?", id).Find(&m).Error; err != nil {
		return nil, err
	}
	if MysqlDb.Table("user").Where("user_id = ?", id).Find(&m).RowsAffected == 0 {
		return nil, nil
	}
	return &m, nil
}

func (dao *userDao) GetUserByName(username string) (*model.UserModel, error) {

	var m model.UserModel

	if err := MysqlDb.Table("user").Where("user_name = ?", username).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *userDao) AddUserModel(m *model.UserModel) error {
	return MysqlDb.Save(m).Error
}
