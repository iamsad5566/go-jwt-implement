package controller

import (
	"go-jwt-implement/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(context *gin.Context) {
	usr := config.User{Account: "test", Password: "5678"}
	s, _ := config.GenerateToken(usr)

	_, str, _ := config.ParseToken(s)
	context.JSON(http.StatusOK, gin.H{"res": s, "parse": str})
}
