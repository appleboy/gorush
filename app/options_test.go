package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOptions(t *testing.T) {
	opts := NewOptions()
	assert.NotNil(t, opts)
	assert.False(t, opts.ShowVersion)
	assert.False(t, opts.Ping)
	assert.Empty(t, opts.ConfigFile)
	assert.Empty(t, opts.Token)
	assert.Empty(t, opts.Message)
}

func TestOptions_CLISendOptions(t *testing.T) {
	opts := &Options{
		Token:   "test-token",
		Message: "test-message",
		Title:   "test-title",
		Topic:   "test-topic",
	}

	sendOpts := opts.CLISendOptions()
	assert.Equal(t, "test-token", sendOpts.Token)
	assert.Equal(t, "test-message", sendOpts.Message)
	assert.Equal(t, "test-title", sendOpts.Title)
	assert.Equal(t, "test-topic", sendOpts.Topic)
}

func TestOptions_IsCLIMode(t *testing.T) {
	tests := []struct {
		name     string
		opts     *Options
		expected bool
	}{
		{
			name:     "no platform enabled",
			opts:     &Options{},
			expected: false,
		},
		{
			name: "android enabled",
			opts: func() *Options {
				o := NewOptions()
				o.Conf.Android.Enabled = true
				return o
			}(),
			expected: true,
		},
		{
			name: "ios enabled",
			opts: func() *Options {
				o := NewOptions()
				o.Conf.Ios.Enabled = true
				return o
			}(),
			expected: true,
		},
		{
			name: "huawei enabled",
			opts: func() *Options {
				o := NewOptions()
				o.Conf.Huawei.Enabled = true
				return o
			}(),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.opts.IsCLIMode())
		})
	}
}
