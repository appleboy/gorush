package gorush

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

var version string

// SetVersion for setup version string.
func SetVersion(ver string) {
	version = ver
}

// GetVersion for get current version.
func GetVersion() string {
	return version
}

// PrintGoRushVersion provide print server engine
func PrintGoRushVersion() {
	fmt.Printf(`GoRush %s, Compiler: %s %s, Copyright (C) 2019 Bo-Yi Wu, Inc.`,
		version,
		runtime.Compiler,
		runtime.Version())
	fmt.Println()
}

// VersionMiddleware : add version on header.
func VersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		c.Header("X-GORUSH-VERSION", version)
		c.Next()
	}
}
