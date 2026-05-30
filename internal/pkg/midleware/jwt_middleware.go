package midleware

import (
	"apotek-pos-go/internal/pkg/response"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.Response[any]{
				Code:    401,
				Message: "Akses ditolak: kamu tidak membawa token",
				Data:    nil,
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, response.Response[any]{
				Code:    401,
				Message: "Akses ditolak: Guanakan format Bearer <token>",
				Data:    nil,
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, response.Response[any]{
				Code:    401,
				Message: "Akses ditolak: token tidak valid atau sudah kadaluarsa",
				Data:    nil,
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("user_id",claims["id"])
			c.Set("role", claims["role"])
			c.Set("username", claims["username"])
		}

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exist := c.Get("role")

		if !exist || role != "admin" {
			c.JSON(http.StatusUnauthorized, response.Response[any]{
				Code:    401,
				Message: "Akses ditolak: Menu ini khusus untuk admin",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func KasirMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exist := c.Get("role")

		if !exist || role != "kasir" {
			c.JSON(http.StatusUnauthorized, response.Response[any]{
				Code:    401,
				Message: "Akses ditolak: Menu ini khusus untuk kasir",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func OwnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exist := c.Get("role")

		if !exist || role != "owner" {
			c.JSON(http.StatusUnauthorized, response.Response[any]{
				Code:    401,
				Message: "Akses ditolak: Menu ini khusus untuk owner",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
