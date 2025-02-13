package middleware

import (
  "net/http"
  "strings"
  "github.com/gin-gonic/gin"
  "github.com/1strewave/cloud-auth/utils"
)

func Authentification() gin.HandlerFunc {
  return func(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "Missing Authentification Token",
      })
      c.Abort()
      return
    }

    tokenParts := strings.Split(tokenString, " ")
    if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "Invalid Authentification Token",
      })
      c.Abort()
      return
    }

    tokenString = tokenParts[1]

    claims, err := utils.VerifyToken(tokenString)
    if err != nil {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "Invalid Authentification Token",
      })
      c.Abort()
      return
    }

    c.Set("user_id", claims["user_id"])
    c.Next()
  }
}