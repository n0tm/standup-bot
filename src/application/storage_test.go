package application

import (
	"github.com/magiconair/properties/assert"
	"gorm.io/gorm"
	"testing"
)

func TestNewStorage(t *testing.T) {
	db := &gorm.DB{}
	assert.Equal(t, NewStorage(db), &storage{db})
}

func Test_storage_Create(t *testing.T) {
	t.Skip("Пока не понятно как тестировать завивисомти без интерфейсов")
}

func Test_storage_First(t *testing.T) {
	t.Skip("Пока не понятно как тестировать завивисомти без интерфейсов")
}
