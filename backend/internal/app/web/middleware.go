package web

import (
	"net/http"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
)

func (a *App) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(contextutils.WithLogger(r.Context(), a.logger))

		next.ServeHTTP(w, r)
	})
}
