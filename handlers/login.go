package handlers

import (
	"fmt"

	"goethe/auth"
	"goethe/data"
	"goethe/db"
	"goethe/util"
	"goethe/views/components"

	"github.com/labstack/echo/v4"
)

func LoginForm(c echo.Context) error {
	ff := components.FormFill{
		Values: components.AccountFormValues{},
		Errors: make(map[string]string),
	}

	return Render(c, components.LoginForm(ff))
}

func Login(c echo.Context) error {
	u, ff := validateLoginForm(c)
	if len(ff.Errors) > 0 {
		return Render(c, components.LoginForm(ff))
	}

	jwt, err := auth.CreateJWT(u)
	if err != nil {
		fmt.Println("Failed to create JWT", err)
	}

	err = util.WriteCookie(c, "JWT", jwt)
	if err != nil {
		fmt.Println("Cookie failed to write")
	}

	// c.Response().Header().Add("Hx-Reswap", "delete")
	c.Response().Header().Add("Hx-Retarget", "#sign-in")
	c.String(200, "Logged in!")
	return nil
	// return Render(c, components.NavigationBar())
}

func validateLoginForm(c echo.Context) (data.User, components.FormFill) {
	un := c.FormValue("username")
	pw := c.FormValue("password")

	ff := components.FormFill{
		Values: components.AccountFormValues{
			Username: un,
			Password: pw,
		},
		Errors: make(map[string]string),
	}

	u, err := db.GetUserAccountByUsername(ff.Values.Username)
	if err != nil || !auth.CheckPassword(u.Password, pw) {
		ff.Errors["INVALID_LOGIN"] = "Incorrect username or password"
	}

	return u, ff
}
