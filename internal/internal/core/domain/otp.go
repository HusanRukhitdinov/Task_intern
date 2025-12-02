package domain

import "errors"

type OTPCreateRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type OTPConfirmRequest struct {
	OTPID string `json:"otp_id" binding:"required"`
	Code  string `json:"code" binding:"required,len=6"`
}

type OTPResponse struct {
	OTPID string `json:"otp_id"`
}
type OTPConfirmationTokenResponse struct {
	Token string `json:"token"`
}

var (
	ErrOTPAlreadyExists            = errors.New("otp already exists")
	ErrIncorrectOTP                = errors.New("incorrect otp code")
	ErrOTPConfirmationTokenExpired = errors.New("otp confirmation token expired")
	ErrOTPNotFound                 = errors.New("otp not found")
	ErrUserAlreadyExists           = errors.New("user already exists")
)
