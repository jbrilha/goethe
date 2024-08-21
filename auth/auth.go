package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"goethe/data"
	"goethe/env"
	"goethe/util"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Checks for JWT validity, if OK returns the fn handlerFunc; otherwise returns the altfn
// TODO this can probably be improved but I'm not sure how to do error handling with HTMX yet
func WithJWT(fn echo.HandlerFunc, altfn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := util.ReadCookie(c, "JWT")
		if err != nil {
			fmt.Println("err reading cookie in jwt middleware", err)
			if c.Request().Header.Get("HX-Request") != "" {
				c.Response().Header().Add("HX-Retarget", "#base")
				c.Response().Header().Add("HX-Reswap", "innerHTML")
				return altfn(c)
			}
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		} else {
			token, err := validateJWT(cookie.Value)
			if err != nil {
				fmt.Println(err)
				if c.Request().Header.Get("HX-Request") != "" {
					c.Response().Header().Add("HX-Retarget", "#base")
					c.Response().Header().Add("HX-Reswap", "innerHTML")
					return altfn(c)
				}
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}

			claims := token.Claims.(jwt.MapClaims)

			fmt.Println("CLAIMS:", claims)
			fmt.Println("TOKEN:", token.Raw)

			// c.Set("JWT", token.Raw)

		}
		return fn(c)
	}
}

func CreateJWT(u data.User, remember bool) (string, error) {
    var expiresAt *jwt.NumericDate
    if remember {
        expiresAt = jwt.NewNumericDate(time.Now().Add(720 * time.Hour)) // 30 days
    } else {
        expiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
    }

	claims := jwt.RegisteredClaims{
		ExpiresAt: expiresAt,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "goethe",
		Subject:   u.Username,
		ID:        string(u.ID),
		// Audience:  []string{"somebody_else"},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := env.JWTSecret()

	return token.SignedString([]byte(secret))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := env.JWTSecret()

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func IsAuthenticated(c context.Context) bool {
	jwtValue := c.Value("JWT")
	if jwtValue == nil {
		fmt.Println("JWT token is missing")
		return false
	}

	jwtString, ok := jwtValue.(string)
	if !ok {
		fmt.Println("JWT token is not a string") // this should never happen but alas
		return false
	}
	_, err := validateJWT(jwtString)
	if err != nil {
		return false
	}

	return true
}

func CheckPassword(encryptedPassword, password string) bool {
	return encryptedPassword == password
}

// func CheckForJWT(fn echo.HandlerFunc, altfn echo.HandlerFunc) echo.MiddlewareFunc {
// 	return func(hf echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			cook, err := util.ReadCookie(c, "JWT")
// 			if err != nil {
// 				fmt.Println("err jwt middleware", err)
// 				if c.Request().Header.Get("HX-Request") != "" {
// 					return altfn(c)
// 				}
// 				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
// 			} else {
// 				_, err = ValidateJWT(cook.Value)
// 				if err != nil {
// 					fmt.Println(err)
// 					if c.Request().Header.Get("HX-Request") != "" {
// 						return altfn(c)
// 					}
// 					return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
// 				}
//
// 			}
// 			return hf(c)
// 		}
// 	}
// }
