package middleware

import (
	"log"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				log.Println("JWT Missing Bearer prefix")
				return echo.ErrUnauthorized
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.ErrUnauthorized
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				log.Printf("JWT Parsing error: %v\n", err)
				return echo.ErrUnauthorized
			}

			// set token ke context jika ingin akses di handler berikutnya
			c.Set("user", token)

			return next(c)
		}
	}
}
