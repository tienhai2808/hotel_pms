package port

import "time"


type JWTProvider interface {
	GenerateToken(userID int64, tokenVersion int, ttl time.Duration) (string, error)

	ParseToken(tokenStr string) (int64, int, time.Duration, error)
}