package app

import md "github.com/saaitt/gomem/manager/domain"

type App struct {
	mUc md.UseCase
}

func (a *App) InitModules() {
	a.initManager()
}

func New() *App {
	return &App{}
}
