package handlers

import (
	"net/http"
	"strings"

	"goethe/db"
	"goethe/views/profile"

	"github.com/labstack/echo/v4"
)

func ProfileBase(c echo.Context) error {
	p := c.Param("username")

	if un := strings.TrimSuffix(p, ".json"); un != p {
		user, err := db.GetUserAccountByUsername(un)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusOK, user)
	}

	user, err := db.GetUserAccountByUsername(p)
	if err != nil {
        // TODO handle error page
	}
	return Render(c, profile.Index(user))
}
