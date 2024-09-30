package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/paghapour/golang-clean-web-api/api/handlers"
	"github.com/paghapour/golang-clean-web-api/config"
)

func User(router *gin.RouterGroup, cfg *config.Config){
	h := handlers.NewUsersHandlers(cfg)

	router.POST("/send-otp", h.SendOtp)
}