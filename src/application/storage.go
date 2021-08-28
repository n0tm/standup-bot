package application

import "gorm.io/gorm"

type Storage interface {
	First(model interface{}, conditions ...interface{})
	Create(model interface{})
}

func NewStorage(gorm *gorm.DB) Storage {
	return &storage{gorm}
}

type storage struct {
	db *gorm.DB
}

func (s storage) Create(model interface{}) {
	s.db.Create(model)
}

func (s storage) First(model interface{}, conditions ...interface{}) {
	s.db.First(model, conditions)
}
