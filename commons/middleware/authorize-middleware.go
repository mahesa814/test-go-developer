package middleware

import (
	"errors"
	"net/http"
	"strings"
	commonResponse "test-go-developer/commons/response"
	"test-go-developer/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IsLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			commonResponse.ResponseFormater(c, http.StatusUnauthorized, "error", "Authorization header is required", struct {
			}{})
			c.Abort()
			return
		}

		// The token usually comes in the format "Bearer <token>"
		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			// Return the secret key
			return []byte(configs.JwtSecret), nil
		})
		if err != nil {
			commonResponse.ResponseFormater(c, http.StatusUnauthorized, "error", "Invalid Token", struct {
			}{})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Extract user ID from token claims
			userID := claims["user_id"].(string)
			// Find the user in the database
			var user entities.Users
			if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					commonResponse.ResponseFormater(c, http.StatusUnauthorized, "error", "User not found", struct {
					}{})
					c.Abort()
					return
				}
				commonResponse.ResponseFormater(c, http.StatusInternalServerError, "error", err.Error(), struct {
				}{})
				c.Abort()
				return
			}
			if user.RememberToken == nil {
				commonResponse.ResponseFormater(c, http.StatusUnauthorized, "error", "User unauthorized", struct {
				}{})
			}
			// Attach user information to the context
			c.Set("userID", userID)
			c.Set("user", &user)
			c.Next()
		} else {
			commonResponse.ResponseFormater(c, http.StatusUnauthorized, "error", "Invalid token claims", struct {
			}{})
			c.Abort()
		}
	}
}
