package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goparking/internals/libs/logger"
	"net/http"
	"strings"

	"goparking/pkgs/jwt"
	"goparking/pkgs/redis"
)

func JWTAuth(cache redis.IRedis) gin.HandlerFunc {
	return JWT(jwt.AccessTokenType, cache)
}

func JWTRefresh(cache redis.IRedis) gin.HandlerFunc {
	return JWT(jwt.RefreshTokenType, cache)
}

func JWT(tokenType string, cache redis.IRedis) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		// Lấy dữ liệu từ Redis
		var rawValue string
		err := cache.Get(fmt.Sprintf("blacklist:%s", strings.ReplaceAll(token, " ", "_")), &rawValue)
		if err != nil {
			logger.Error("Failed to get value from Redis:", err)
		}

		var value map[string]string
		err = json.Unmarshal([]byte(rawValue), &value)
		if err != nil {
			logger.Error("Failed to unmarshal JSON:", err)
		}

		if value["status"] == "blacklisted" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
			c.Abort()
			return
		}

		payload, err := jwt.ValidateToken(token)
		if err != nil || payload == nil || payload["type"] != tokenType {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		c.Set("userId", payload["id"])
		//c.Set("role", payload["role"])
		c.Set("token", token)
		c.Next()
	}
}
