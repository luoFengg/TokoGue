package middleware

import (
	"net/http"
	"strings"
	"tokogue-api/config"
	"tokogue-api/models/web"
	"tokogue-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. Ambil Header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Success:   false,
				Message: "Header Authorization tidak ditemukan",
				Data:   nil,
			})
			return
		}

		// 2. Cek format Header Authorization
		if !strings.HasPrefix(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Success:   false,
				Message: "Format Header Authorization salah, harus: Bearer <token>",
				Data:   nil,
			})
			return
		}

		// 3. Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. Parse & validate token
		token, err := jwt.ParseWithClaims(tokenString, &utils.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.WebResponse{
				Success:   false,
				Message: "Token tidak valid atau kadaluarsa",
				Data:   nil,
			})
			return
		}

		// 5. Simpan claims ke context
		if claims, ok := token.Claims.(*utils.JWTCustomClaims); ok {
			c.Set("userID", claims.Subject)
			c.Set("role", claims.Role)
		}

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil role dari context
		role, exists := c.Get("role")

		// 2. Cek apakah role adalah admin
		if !exists || role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, web.WebResponse{
				Success:   false,
				Message: "Akses ditolak: hanya admin yang dapat mengakses resource ini",
				Data:   nil,
			})
			return
		}

		// 3. Lanjutkan ke handler berikutnya
		c.Next()

	}
}