package controller

import (
	"net/http"
	"orange-site/backend/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username      string
	Authenticated bool
	AuthToken     string `default:""`
	Message       string `default:""`
}

func GetSecret() string {
	secret := os.Getenv("SERVER_SECRET")
	if secret == "" {
		panic("No secret specified")
	}
	return secret
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
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

func Login(c echo.Context) error {
	login := new(LoginInfo)
	err := c.Bind(login)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Find if the username/password match
	user, err := model.FindUser(login.Username)

	// Compare password with hashed version in DB
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{
			Message: "Invalid Username or Password",
		})
	}

	// If authenticated, then need to create an access token
	token, err := createToken(login.Username)
	if err != nil {
		return c.String(http.StatusInternalServerError,
			"Unable to make authentication token")
	} else {
		return c.JSON(http.StatusOK, LoginResponse{
			Authenticated: true,
			AuthToken:     token,
		})
	}
}

func Register(c echo.Context) error {
	login := new(LoginInfo)
	err := c.Bind(login)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Find if the username/password match
	_, err = model.FindUser(login.Username)
	if err == nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{
			Message: "Account name already exists",
		})
	}

	hashedPassword, err := hashPassword(login.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{
			Message: "Password not supported",
		})
	}

	_, err = model.AddUser(login.Username, hashedPassword)
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to make account")
	}

	token, err := createToken(login.Username)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to make authentication token")
	} else {
		return c.JSON(http.StatusOK, LoginResponse{
			Authenticated: true,
			AuthToken:     token,
		})
	}
}
