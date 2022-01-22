package services

import (
	"net/http"
	"time"

	"github.com/ademolguner/gotodo/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type AuthRequest struct {
	UserName string `json:"admin"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"Token"`
}

func HelloAuth(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from auth!")
}

func Login(c echo.Context) error {
	u := &AuthRequest{}

	if err := c.Bind(u); err != nil {
		return err
	}

	// Throws unauthorized error if given username and password incorrect
	if u.UserName != "admin" || u.Password != "password" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Adem OLGUNER"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JwtTokenSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
