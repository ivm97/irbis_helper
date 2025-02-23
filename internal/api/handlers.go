package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *Application) GetLastVersion(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{"version": a.fls.ClientVersion()})
}

func (a *Application) GetFiles(c echo.Context) error {
	file, err := a.fls.GetSourceFiles()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Attachment(*file, a.fls.ClientVersion()+".zip")
}
