package main

import (
	"go.uber.org/zap"
	"net/url"
	"time"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url, _ := url.Parse("http")
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
