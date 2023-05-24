package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	gojwt "github.com/golang-jwt/jwt/v4"
	responseErr "github.com/gigihprasetyo/backend-standard-code/internal/error"
	"github.com/spf13/viper"
)

var (
	ErrMalformedJWT = errors.New("karedensial tidak dapat diketahui")
	ErrUnauthorized = errors.New("kamu tidak memiliki akses")
	ErrDefaultError = errors.New("server sedang sibuk coba beberapa saat lagi")
)

// Protected protect routes
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(viper.GetString("auth.jwt_secret")),
		ErrorHandler: jwtError,
	})
}

// middleware for Guest
func Guest() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return GuestUser(c)
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return responseErr.Response(c, responseErr.New(fiber.StatusUnauthorized, responseErr.WithMessage(ErrMalformedJWT.Error())))

	} else {
		c.Status(fiber.StatusUnauthorized)
		return responseErr.Response(c, responseErr.New(fiber.StatusUnauthorized, responseErr.WithMessage(ErrUnauthorized.Error())))
	}
}

type User struct {
	UUID         string
	Role         []string
	IsFacebook   bool
	IsGooglePlus bool
	Phone        *string
	Email        *string
	Username     *string
}

func ExternalUser(c context.Context) *User {
	if c.Value("user") == nil {
		return nil
	}
	user := c.Value("user").(*gojwt.Token)
	claims := user.Claims.(gojwt.MapClaims)
	var (
		role     []string
		phone    *string
		email    *string
		username *string
	)
	if claims["phone"] != nil {
		getPhone := claims["phone"].(string)
		phone = &getPhone
	}
	if claims["email"] != nil {
		getEmail := claims["email"].(string)
		email = &getEmail
	}
	if len(claims["role"].([]interface{})) > 0 {
		getRole := claims["role"].([]interface{})
		for _, v := range getRole {
			role = append(role, v.(string))
		}
	}
	if claims["username"] != nil {
		getUsername := claims["username"].(string)
		username = &getUsername
	}
	return &User{
		UUID:         claims["uuid"].(string),
		IsFacebook:   claims["is_facebook"].(bool),
		IsGooglePlus: claims["is_google"].(bool),
		Phone:        phone,
		Email:        email,
		Role:         role,
		Username:     username,
	}
}

func (u *User) IsCustomer() bool {
	for _, b := range u.Role {
		if b == "customer" {
			return true
		}
	}
	return false
}

func (u *User) IsOperator() bool {
	for _, b := range u.Role {
		if b == "operator" {
			return true
		}
	}
	return false
}

func GuestUser(c *fiber.Ctx) error {
	authorization := c.Request().Header.Peek("Authorization")
	if authorization != nil {
		splitToken := strings.Split(string(authorization), "Bearer ")
		token, err := gojwt.Parse(splitToken[1], func(token *gojwt.Token) (interface{}, error) {
			return []byte(viper.GetString("auth.jwt_secret")), nil
		})
		if err != nil {
			return responseErr.Response(c, responseErr.New(fiber.StatusUnauthorized, responseErr.WithMessage(ErrUnauthorized.Error())))
		}
		c.Locals("user", token)
		return c.Next()
	}
	c.Locals("user", nil)
	return c.Next()
}

func ChangeSignatureForTransaction(c context.Context) (string, error) {
	user := c.Value("user").(*gojwt.Token)
	claims := user.Claims.(gojwt.MapClaims)

	tokenClaims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	token, err := tokenClaims.SignedString([]byte(viper.GetString("transaction_service.jwt_access_secret")))
	if err != nil {
		return "", err
	}

	return "Bearer " + token, err
}
