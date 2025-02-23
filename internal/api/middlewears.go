package api

/*
func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}

		token = strings.ReplaceAll(token, "Bearer ", "")

		vault.rw.RLock()
		trueToken := vault.key
		vault.rw.RUnlock()

		if trueToken != token {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}
		return next(c)
	}
}
*/
