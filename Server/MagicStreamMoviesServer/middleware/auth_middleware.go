package middleware

import (
	"fmt"
	"net/http"

	"github.com/Tejas-Musale/Magic-Stream-Movies/Server/MagicStreamMoviesServer/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {

        fmt.Println("=== AUTH MIDDLEWARE ===")

        token, err := utils.GetAccessToken(c)

        fmt.Println("TOKEN FOUND:", token != "")
        fmt.Println("COOKIE ERROR:", err)

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        claims, err := utils.ValidateToken(token)

        fmt.Println("VALIDATE ERROR:", err)

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        fmt.Println("USER ID:", claims.UserID)

        c.Set("userId", claims.UserID)
        c.Set("role", claims.Role)

        c.Next()
    }
}
