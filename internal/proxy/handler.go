package proxy

import (
	"fmt"
	"io"
	"net/http"

	"github.com/qs-lzh/caching-proxy/internal/cache"
)

type proxyHandler struct {
	config handlerConfig
	cache  cache.Cache
}

type handlerConfig struct {
	origin string
}

func NewProxyHandler(config handlerConfig, cache cache.Cache) *proxyHandler {
	return &proxyHandler{
		config: config,
		cache:  cache,
	}
}

func (h *proxyHandler) handleProxy(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	url := fmt.Sprintf("%s%s", h.config.origin, path)

	val, err := h.cache.Get(url)
	if err == nil {
		fmt.Printf("get resp from cache\n")
		w.Write([]byte(val))
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get resp from url: %v", err)
	}
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read resp.Body\n")
		return
	}
	fmt.Printf("get resp from %s website\n", url)
	if err = h.cache.Set(url, string(respBodyBytes)); err != nil {
		fmt.Printf("failed to set key %s in redis: %v\n", url, err)
		return
	}
	fmt.Printf("set key %s in redis successfully\n", url)

	w.Write(respBodyBytes)
}
