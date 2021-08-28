package telegram

import "errors"

type Authenticator interface {
	Authenticate(Update) (User, error)
}

func NewAuthenticator(repository UserRepository) Authenticator {
	return &authenticator{repository}
}

type authenticator struct {
	repository UserRepository
}

func (a authenticator) Authenticate(u Update) (User, error) {
	userId := u.Message().UserId()
	if userId == 0 {
		return nil, errors.New("can't identificate user")
	}

	return a.repository.GetById(uint(userId)), nil
}
