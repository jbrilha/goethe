package util

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func WriteCookie(c echo.Context, name, value string) error {
	cookie := new(http.Cookie)
	cookie.Name = name
    cookie.Path = "/"
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = true
    cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)
	// return c.String(http.StatusOK, "write a cookie")
	return nil
}

func ReadCookie(c echo.Context, cookieName string) (*http.Cookie, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
        // fmt.Println("erred in read cookie")
		return nil, err
	}

	return cookie, nil
}

func ReadAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	// return c.String(http.StatusOK, "read all the cookies")
	return nil
}
