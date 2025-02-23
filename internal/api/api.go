package api

import (
	"github.com/irbis_helper/internal/files"
	"github.com/labstack/echo/v4"
)

type Application struct {
	e      *echo.Echo
	fls    *files.Files
	passph string
}

func Initiate(passph string, fls *files.Files) *Application {
	e := echo.New()
	e.HideBanner = true

	return &Application{e, fls, passph}
}

func (a *Application) Start(port string) error {
	a.SetHandlers()
	return a.e.Start(":" + port)

}
