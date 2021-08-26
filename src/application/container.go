package application

import (
	"go.uber.org/dig"
)

type Container interface {
	Invoke(function interface{}) error
	Provide(constructor interface{}) error
}

func NewContainer(di *dig.Container) Container {
	return &container{di}
}

type container struct {
	dig *dig.Container
}

func (c *container) Invoke(function interface{}) error {
	return c.dig.Invoke(function)
}

func (c *container) Provide(constructor interface{}) error {
	return c.dig.Provide(constructor)
}
