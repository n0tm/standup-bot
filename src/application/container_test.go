package application

import (
	"github.com/magiconair/properties/assert"
	"go.uber.org/dig"
	"testing"
)

func TestNewContainer(t *testing.T) {
	digContainer := dig.New()
	assert.Equal(t, NewContainer(digContainer), &container{digContainer})
}

func Test_container_Invoke(t *testing.T) {
	t.Skip("Пока не понятно как тестировать завивисомти без интерфейсов")
}

func Test_container_Provide(t *testing.T) {
	t.Skip("Пока не понятно как тестировать завивисомти без интерфейсов")
}
