package main

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"stravo/repository"
	"time"
)

func main() {
	if err := repository.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	if err := repository.RunMigrations(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	user := e.Group("/user")

	user.GET("", getUser, auth)
	user.POST("", createUser)
	user.POST("/login", login)

	e.Logger.Fatal(e.Start(":8000"))
}

const (
	jwtSigningKey = "secret"
	keyUserId     = "userId"
)

var (
	auth = echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtSigningKey),
		SuccessHandler: func(c echo.Context) {
			j := c.Get("user").(*jwt.Token) // parse token success so the value has to be in the context
			claims := j.Claims.(jwt.MapClaims)
			c.Set(keyUserId, claims["userId"])
		},
	})
)

type jwtClaims struct {
	UserID string              `json:"userId"`
	Role   repository.Userrole `json:"role"`
	jwt.RegisteredClaims
}

func login(c echo.Context) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request

	err := c.Bind(&req)
	if err != nil {
		return err
	}

	user, err := repository.DB.GetUserByEmail(c.Request().Context(), sql.NullString{
		String: req.Email,
		Valid:  true,
	})
	if err != nil {
		return err
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return err
	}
	if hashedPassword != user.Hashedpassword.String {
		// TODO 200 or 500?
		return errors.New("password is incorrect")
	}

	// Set custom claims
	claims := &jwtClaims{
		UserID: user.ID.String(),
		Role:   user.Role.Userrole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSigningKey))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func getUser(c echo.Context) error {
	userID, _ := c.Get(keyUserId).(string)
	return c.JSON(http.StatusOK, echo.Map{
		"userId": userID,
	})
}

func createUser(c echo.Context) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	// save user
	user, err := repository.DB.CreateUser(c.Request().Context(), repository.CreateUserParams{
		Email:          sql.NullString{String: req.Email, Valid: true},
		Hashedpassword: sql.NullString{String: hashedPassword, Valid: true},
		Role:           repository.NullUserrole{Userrole: repository.UserroleUser, Valid: true},
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"userId": user.ID.String(),
	})
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
