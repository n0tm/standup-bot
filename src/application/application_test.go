package application

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	appInstance := Instance()
	assert.Equal(t, appInstance, &application{})
	appInstance.SetContainer(NewContainerMock(t))
	assert.Equal(t, Instance(), appInstance)
}

func Test_application_Container(t *testing.T) {
	container := NewContainerMock(t)
	assert.Equal(t, (&application{container}).Container(), container)
}

func Test_application_SetContainer(t *testing.T) {
	container := NewContainerMock(t)

	actualApp := &application{}
	actualApp.SetContainer(container)

	assert.Equal(t, &application{container}, actualApp)
}
