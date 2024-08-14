package handlers

import (
	"fmt"
	"strconv"

	"goethe/auth"
	"goethe/data"
	"goethe/db"
	"goethe/util"
	"goethe/views/components"

	"github.com/labstack/echo/v4"
)

func LoginForm(c echo.Context) error {
	sf, err := strconv.ParseBool(c.FormValue("showForm"))
	if err != nil {
		fmt.Println("Failed to parse bool")
	}

	if sf {
		return Render(c, components.NavigationBarWForm(components.AccountFormValues{}, make(map[string]string), true))
	}

	return Render(c, components.NavigationBar())
}

func Login(c echo.Context) error {
	u, v, e := validateLoginForm(c)
	if len(e) > 0 {
		fmt.Println(v, e)
		return Render(c, components.NavigationBarWForm(v, e, true))
	}

	jwt, err := auth.CreateJWT(u)
	if err != nil {
		fmt.Println("Failed to create JWT", err)
	}

	err = util.WriteCookie(c, "JWT", jwt)
	if err != nil {
		fmt.Println("Cookie failed to write")
	}

	return Render(c, components.NavigationBar())
}

func validateLoginForm(c echo.Context) (data.User, components.AccountFormValues, map[string]string) {
	un := c.FormValue("username")
	pw := c.FormValue("password")

	v := components.AccountFormValues{
		Username: un,
		Password: pw,
	}
	e := make(map[string]string)

	u, err := db.GetUserAccount(v.Username)
	if err != nil || !auth.CheckPassword(u.Password, pw) {
		e["INVALID_LOGIN"] = "Incorrect username or password"
	}

	return u, v, e
}
