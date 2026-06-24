package middleware

import (
	"net/http"

	"github.com/butbeautifulv/veneno/engage/serve/internal/config"
	"github.com/butbeautifulv/veneno/pkg/api"
	"github.com/butbeautifulv/veneno/pkg/auth"
)

func Auth(stack *auth.Stack, strict bool, sec config.SecurityConfig, next http.Handler) http.Handler {
	return api.AuthMiddleware(stack, strict, sec.Prod, auth.PermEngageToolRun, next)
}
