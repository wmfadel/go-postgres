package models

type AuthenticationResponse struct {
	User
	Token     string `json:"token"`
	ExpiresAt int    `json:"expires_at"`
	Message   string `json:"message"`
}
