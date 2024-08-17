package handlers

import (
	"fmt"
	"time"

	"goethe/auth"
	"goethe/data"
	"goethe/db"
	"goethe/util"
	"goethe/views/components"

	"github.com/labstack/echo/v4"
)

func RegisterForm(c echo.Context) error {
	v := components.AccountFormValues{}
    e := make(map[string]string)

	return Render(c, components.RegisterForm(v, e))
}

func Register(c echo.Context) error {
	v, e := validateRegisterForm(c)
	if len(e) > 0 {
		return Render(c, components.RegisterForm(v, e))
	}

	u := data.User{
		Username:  v.Username,
		Email:     "nomail",
		Password:  v.Password,
		CreatedAt: time.Now(),
	}

	_, err := db.InsertUserAccount(&u)
	if err != nil {
		fmt.Println("err in insetion:", err)
	}
	fmt.Println(u.ID)

	jwt, err := auth.CreateJWT(u)
	if err != nil {
		fmt.Println("Failed to create JWT", err)
	}

	err = util.WriteCookie(c, "JWT", jwt)
	if err != nil {
		fmt.Println("Cookie failed to write")
	}

    // c.Response().Header().Add("Hx-Reswap", "outerHTML")
    c.Response().Header().Add("Hx-Retarget", "#sign-in")
	c.String(200, "Logged in!")
    return nil
	// return Render(c, components.RegisterForm(v, e))
}

func validateRegisterForm(c echo.Context) (components.AccountFormValues, map[string]string) {
	un := c.FormValue("username")
	pw := c.FormValue("password")
	pwc := c.FormValue("confirmation")

	v := components.AccountFormValues{
		Username:     un,
		Password:     pw,
		Confirmation: pwc,
	}

	e := make(map[string]string)
	if len(pw) < 5 {
		e["PW_LEN"] = "Password length must be at least 5"
	}
	if pw != pwc {
		e["PW_CONF"] = "Confirmation does not match password"
	}

	exists, _ := db.UserAccountExists(un)
	if exists {
		e["USER_EXISTS"] = "Username already taken"
	}

	return v, e
}
