package application

var instance *application

func Instance() Application {
	if instance == nil {
		instance = &application{}
	}

	return instance
}

type Application interface {
	Container() Container
	SetContainer(Container)
}

type application struct {
	container Container
}

func (a *application) Container() Container {
	return a.container
}

func (a *application) SetContainer(c Container) {
	a.container = c
}
