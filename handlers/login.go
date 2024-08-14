package handlers

import (
	"fmt"
	"goethe/auth"
	"goethe/db"
	"goethe/util"
	"goethe/views/components"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginForm(c echo.Context) error {
	sf, err := strconv.ParseBool(c.FormValue("showForm"))
	if err != nil {
		fmt.Println("Failed to parse bool")
	}

	if sf {
		return Render(c, components.NavigationBarWForm(components.LoginFormValues{}, make(map[string]string), true))
	}

	return Render(c, components.NavigationBar(true))
}

func Login(c echo.Context) error {
	v, e := validateLoginForm(c, false)
	if len(e) > 0 {
		fmt.Println(v, e)
		return Render(c, components.NavigationBarWForm(v, e, true))
	}

    un := v.Username // username
    user := db.GetUserAccount(un)

	jwt, err := auth.CreateJWT(user)
	if err != nil {
		fmt.Println("Failed to create JWT", err)
	}

	err = util.WriteCookie(c, "JWT", jwt)
	if err != nil {
		fmt.Println("Cookie failed to write")
	}

	return Render(c, components.NavigationBar(true))
}

func RegisterForm(c echo.Context) error {
	u := c.FormValue("username")
	p := c.FormValue("password")
	pc := c.FormValue("confirmation")

	v := components.LoginFormValues{
		Username:     u,
		Password:     p,
		Confirmation: pc,
	}

    fmt.Println(v)

	return Render(c, components.RegisterForm(v, make(map[string]string)))
}

func Register(c echo.Context) error {
	v, e := validateLoginForm(c, true)
	if len(e) > 0 {
		fmt.Println(v, e)
		return Render(c, components.NavigationBarWForm(v, e, false))
	}

    un := v.Username // username
    user := db.GetUserAccount(un)

	jwt, err := auth.CreateJWT(user)
	if err != nil {
		fmt.Println("Failed to create JWT", err)
	}

	err = util.WriteCookie(c, "JWT", jwt)
	if err != nil {
		fmt.Println("Cookie failed to write")
	}

	return Render(c, components.NavigationBar(true))
}

func validateLoginForm(c echo.Context, checkConfirm bool) (components.LoginFormValues, map[string]string) {
	u := c.FormValue("username")
	p := c.FormValue("password")
	pc := c.FormValue("confirmation")

	v := components.LoginFormValues{
		Username:     u,
		Password:     p,
		Confirmation: pc,
	}

	e := make(map[string]string)
	if len(p) < 5 {
		e["PW_LEN"] = "Password length must be at least 5"
	}
	if checkConfirm && p != pc {
		e["PW_CONF"] = "Confirmation does not match password"
	}

	return v, e
}
