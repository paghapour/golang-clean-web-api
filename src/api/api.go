package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/paghapour/golang-clean-web-api/api/middlewares"
	"github.com/paghapour/golang-clean-web-api/api/routers"
	validation "github.com/paghapour/golang-clean-web-api/api/validations"
	"github.com/paghapour/golang-clean-web-api/config"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidator()

	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), /*middlewares.TestMiddleware()*/ middlewares.LimitByRequest())

	RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")

		routers.Health(health)
	}
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}
