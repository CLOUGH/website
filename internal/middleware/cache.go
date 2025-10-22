package middleware

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// CacheEntry stores cached response data
type CacheEntry struct {
	Data        []byte
	ContentType string
	StatusCode  int
	ExpiresAt   time.Time
}

// Cache manages response caching
type Cache struct {
	store map[string]*CacheEntry
	mu    sync.RWMutex
	ttl   time.Duration
}

// NewCache creates a new cache with specified TTL
func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		store: make(map[string]*CacheEntry),
		ttl:   ttl,
	}

	// Cleanup expired entries every minute
	go cache.cleanup()

	return cache
}

// generateKey creates a cache key from request
func generateKey(c *gin.Context) string {
	// Include path, query params, and relevant headers
	key := c.Request.URL.Path + "?" + c.Request.URL.RawQuery
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}

// get retrieves cached entry if valid
func (cache *Cache) get(key string) (*CacheEntry, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	entry, exists := cache.store[key]
	if !exists {
		return nil, false
	}

	if time.Now().After(entry.ExpiresAt) {
		return nil, false
	}

	return entry, true
}

// set stores a cache entry
func (cache *Cache) set(key string, entry *CacheEntry) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry.ExpiresAt = time.Now().Add(cache.ttl)
	cache.store[key] = entry
}

// cleanup removes expired entries
func (cache *Cache) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		cache.mu.Lock()
		now := time.Now()
		for key, entry := range cache.store {
			if now.After(entry.ExpiresAt) {
				delete(cache.store, key)
			}
		}
		cache.mu.Unlock()
	}
}

// Middleware returns a Gin middleware handler for caching GET requests
func (cache *Cache) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only cache GET requests
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		key := generateKey(c)

		// Check cache
		if entry, found := cache.get(key); found {
			// Set cache headers
			c.Header("X-Cache", "HIT")
			c.Header("Cache-Control", "public, max-age=300")
			c.Data(entry.StatusCode, entry.ContentType, entry.Data)
			c.Abort()
			return
		}

		// Capture response
		writer := &responseWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		c.Header("X-Cache", "MISS")
		c.Next()

		// Cache successful responses (200-299)
		if writer.Status() >= 200 && writer.Status() < 300 {
			entry := &CacheEntry{
				Data:        writer.body.Bytes(),
				ContentType: writer.Header().Get("Content-Type"),
				StatusCode:  writer.Status(),
			}
			cache.set(key, entry)

			// Add cache control header for client-side caching
			c.Header("Cache-Control", "public, max-age=300")
		}
	}
}

// responseWriter captures the response for caching
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
