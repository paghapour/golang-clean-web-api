package dto

type GetOtpRequest struct{
	MobileNumber string `json:"MobileNumber" binding:"required,mobile,min=11,max=11"` 
}