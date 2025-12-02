package domain

type TokenRequest struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}
type TokenResponse struct {
	AccessToken        string  `json:"access_token"`
	RefreshToken       string  `json:"refresh_token"`
	AccessExpiredTime  float64 `json:"access_expired_time"`
	RefreshExpiresTime float64 `json:"refresh_expires_time"`
	Success            bool    `json:"success"`
}
