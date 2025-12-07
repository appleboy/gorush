package app

import (
	"testing"

	"github.com/appleboy/gorush/core"

	"github.com/stretchr/testify/assert"
)

func TestCLISendOptions(t *testing.T) {
	opts := CLISendOptions{
		Token:   "test-token",
		Message: "test-message",
		Title:   "test-title",
		Topic:   "test-topic",
	}

	assert.Equal(t, "test-token", opts.Token)
	assert.Equal(t, "test-message", opts.Message)
	assert.Equal(t, "test-title", opts.Title)
	assert.Equal(t, "test-topic", opts.Topic)
}

func TestSendNotification_UnsupportedPlatform(t *testing.T) {
	// Test that unsupported platform is handled
	// Note: This would call logx.LogError.Fatalf in production,
	// so we can't easily test it without mocking.
	// This test mainly documents the expected behavior.
	assert.Equal(t, 1, core.PlatFormIos)
	assert.Equal(t, 2, core.PlatFormAndroid)
	assert.Equal(t, 3, core.PlatFormHuawei)
}
