package models

type AuthenticationResponse struct {
	User
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	Message   string `json:"message"`
}
