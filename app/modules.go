package app

import managerUC "github.com/saaitt/gomem/manager/usecase"

func (a *App) initManager() {
	a.mUc = managerUC.New()
}
