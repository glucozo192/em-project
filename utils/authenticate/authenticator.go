package authenticate

import "time"

type Authenticator interface {
	CreateToken(userID string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
