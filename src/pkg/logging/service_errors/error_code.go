package service_errors

const (
	// Token
	UnExpectedError = "expected error"
	ClaimNotFound = "Claims not found"
	
	// OTP
	OtpExists = "Otp exists"
	OtpUsed = "Otp used"
	OtpInvalid = "Otp invalid"

	// User
	EmailExists = "Email exists"
	UsernameExists = "Username exists"
)