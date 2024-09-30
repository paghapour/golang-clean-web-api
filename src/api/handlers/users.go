package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paghapour/golang-clean-web-api/api/dto"
	"github.com/paghapour/golang-clean-web-api/api/helper"
	"github.com/paghapour/golang-clean-web-api/config"
	"github.com/paghapour/golang-clean-web-api/services"
)

type UsersHandlers struct{
	service *services.UserService
}


func NewUsersHandlers (cfg *config.Config) *UsersHandlers{
	service := services.NewUserService(cfg)
	return&UsersHandlers{service: service}
}


// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UsersHandlers) SendOtp(c *gin.Context){
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = h.service.SendOtp(req)
	if err != nil{
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	// call internal sms service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}