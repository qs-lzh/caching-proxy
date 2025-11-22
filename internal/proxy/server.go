package proxy

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/qs-lzh/caching-proxy/internal/cache"
)

type ProxyConfig struct {
	Port   int
	Origin string
}

func StartProxyServer(config ProxyConfig) error {
	redisAddr := "localhost:6379"
	cache := cache.NewRedisCache(redisAddr)

	r := chi.NewRouter()
	h := NewProxyHandler(handlerConfig{origin: config.Origin}, cache)
	r.Get("/*", h.handleProxy)

	proxyAddr := fmt.Sprintf("localhost:%d", config.Port)
	srv := http.Server{
		Addr:    proxyAddr,
		Handler: r,
	}
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
