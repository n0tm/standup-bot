package telegram

import "standup-bot/src/application"

type UserRepository interface {
	GetById(uint) User
}

func NewUserRepository(storage application.Storage) UserRepository {
	return &userRepository{storage}
}

type userRepository struct {
	storage application.Storage
}

func (u userRepository) GetById(i uint) User {
	var user UserModel

	u.storage.First(&user, i)

	if user.GetId() == 0 {
		return nil
	}

	return &user
}
