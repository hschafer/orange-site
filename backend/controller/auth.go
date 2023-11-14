package controller

import (
	"net/http"
	"orange-site/backend/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetSecret() string {
	secret := os.Getenv("SERVER_SECRET")
	if secret == "" {
		panic("No secret specified")
	}
	return secret
}

func createToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := GetSecret()
	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	} else {
		return tokenStr, nil
	}

}

type LoginResponse struct {
	Username      string
	Authenticated bool
	AuthToken     string `default:""`
	Message       string `default:""`
}

func Login(c echo.Context) error {
	login := new(LoginInfo)
	err := c.Bind(login)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Find if the username/password match
	authenticated, err := model.AuthenticateUser(login.Username, login.Password)

	if authenticated {
		// If authenticated, then need to create an access token
		token, err := createToken(login.Username)

		if err != nil {
			return c.String(http.StatusInternalServerError, "Unable to make authentication token")
		} else {
			return c.JSON(http.StatusOK, LoginResponse{
				Authenticated: true,
				AuthToken:     token,
			})
		}
	} else {
		return c.JSON(http.StatusBadRequest, LoginResponse{
			Authenticated: false,
			Message:       "Invalid Username or Password",
		})
	}
}
