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
	v := components.AccountFormValues{}
    e := make(map[string]string)

	return Render(c, components.LoginForm(v, e))
}

func Login(c echo.Context) error {
	u, v, e := validateLoginForm(c)
	if len(e) > 0 {
		fmt.Println(v, e)
		return Render(c, components.LoginForm(v, e))
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
