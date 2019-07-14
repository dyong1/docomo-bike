package env

import "strings"

func IsProd(env string) bool {
	return strings.ToLower(env) == "production"
}
func IsDev(env string) bool {
	return strings.ToLower(env) == "development"
}
