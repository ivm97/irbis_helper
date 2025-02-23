package api

func (a *Application) SetHandlers() {
	a.e.GET("/irb/get", a.GetFiles)
	a.e.GET("/irb/ver", a.GetLastVersion)
}
