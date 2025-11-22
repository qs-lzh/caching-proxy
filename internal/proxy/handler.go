package proxy

import (
	"fmt"
	"io"
	"net/http"

	"github.com/qs-lzh/caching-proxy/internal/cache"
)

type proxyHandler struct {
	config handlerConfig
	cache  *cache.RedisCache
}

type handlerConfig struct {
	origin string
}

func NewProxyHandler(config handlerConfig, cache *cache.RedisCache) *proxyHandler {
	return &proxyHandler{
		config: config,
		cache:  cache,
	}
}

func (h *proxyHandler) handleProxy(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	url := fmt.Sprintf("%s%s", h.config.origin, path)

	cachedResp, err := h.cache.Get(url)
	if err == nil {
		fmt.Printf("get resp from cache\n")
		for k, v := range cachedResp.Header {
			w.Header()[k] = v
		}
		w.Header()["X-Cache"] = []string{"HIT"}
		w.WriteHeader(cachedResp.StatusCode)
		w.Write(cachedResp.Body)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get resp from url: %v\n", err)
		return
	}
	fmt.Printf("get resp from %s website\n", url)
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read from resp.Body: %v\n", err)
		return
	}
	cachedResp = cache.CachedResponse{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       respBodyBytes,
	}
	if err = h.cache.Set(url, cachedResp); err != nil {
		fmt.Printf("failed to set key %s in redis: %v\n", url, err)
		return
	}
	fmt.Printf("set key %s in redis successfully\n", url)

	w.Write(respBodyBytes)
}
