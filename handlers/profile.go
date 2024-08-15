package handlers

import (
	"goethe/db"
	"goethe/views/profile"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func ProfileBase(c echo.Context) error {
	p := c.Param("username")

	if un := strings.TrimSuffix(p, ".json"); un != p {
		user, err := db.GetUserAccount(un)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusOK, user)
	}

    user, _ := db.GetUserAccount(p)
	return Render(c, profile.Index(user))
}

