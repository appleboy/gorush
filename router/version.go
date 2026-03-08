package router

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	version string
	commit  string
)

// SetVersion for setup version string.
func SetVersion(ver string) {
	version = ver
}

// SetCommit for setup commit string.
func SetCommit(ver string) {
	commit = ver
}

// GetVersion for get current version.
func GetVersion() string {
	return version
}

// PrintGoRushVersion provide print server engine
func PrintGoRushVersion() {
	if len(commit) > 7 {
		commit = commit[:7]
	}

	fmt.Fprintf(
		os.Stdout,
		"GoRush %s, Commit: %s, Compiler: %s %s, Copyright (C) %d Bo-Yi Wu, Inc.\n",
		version,
		commit,
		runtime.Compiler,
		runtime.Version(),
		time.Now().Year(),
	)
}

// VersionMiddleware : add version on header.
func VersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		c.Header("X-GORUSH-VERSION", version)
		c.Next()
	}
}
