package telegram

import "gorm.io/gorm"

type User interface {
	GetId() uint
	SetId(uint)
}

type UserModel struct {
	gorm.Model
}

func (u UserModel) TableName() string {
	return "telegram_users"
}

func (u UserModel) GetId() uint {
	return u.ID
}

func (u *UserModel) SetId(id uint) {
	u.ID = id
}
