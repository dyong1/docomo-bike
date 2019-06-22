package auth

type Service interface {
	JWTAuthorize(userID string, plainPassword string) (JWTAuthResult, error)
}

type JWTAuthResult struct{}
