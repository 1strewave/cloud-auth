package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/1strewave/cloud-auth/models"
    "github.com/1strewave/cloud-auth/utils"
)

func Login(c *gin.Context) {
  var user models.User

  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Invalid data",
    })
    return
  }
}