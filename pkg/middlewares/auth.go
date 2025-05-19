package middlewares

import (
	"demo/dto/responses"
	"demo/internal/configs"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey string

func init() {
	configs.LoadEnv()

	secretKey = configs.GetEnv("JWT_SIGNER")
}

func AuthenticationMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.IndentedJSON(http.StatusUnauthorized, responses.ErrorResponse[string](errors.New("you need to login")))
		c.Abort()
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	result, err := VerifyToken(token)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, responses.ErrorResponse[string](err))
		c.Abort()
		return
	}

	expire, err := result.Claims.GetExpirationTime()

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, responses.ErrorResponse[string](err))
		c.Abort()
		return
	}

	if expire.After(time.Now()) {
		c.IndentedJSON(http.StatusUnauthorized, responses.ErrorResponse[string](errors.New("token expired")))
		c.Abort()
		return
	}

	c.Next()
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Return the verified token
	return token, nil
}
