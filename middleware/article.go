package middleware

import (
	"context"
	"net/http"

	"github.com/pressly/chi"
)

// ArticleCtx 文章节点
func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, "articleID")
		ctx := context.WithValue(r.Context(), "article", articleID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
