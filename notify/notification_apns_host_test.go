package notify

import (
	"context"
	"sync"
	"testing"

	"github.com/appleboy/gorush/config"

	"github.com/sideshow/apns2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// getApnsClient must pick the APNs host from the per-request override first,
// then fall back to the iOS config default.
func TestGetApnsClientHost(t *testing.T) {
	cfg, err := config.LoadConf()
	require.NoError(t, err)
	cfg.Ios.Enabled = true
	cfg.Ios.KeyPath = testKeyPath
	require.NoError(t, InitAPNSClient(context.Background(), cfg))

	tests := []struct {
		name          string
		cfgProduction bool
		req           *PushNotification
		wantHost      string
	}{
		{"request production override", false, &PushNotification{Production: true}, apns2.HostProduction},
		{"request development override", true, &PushNotification{Development: true}, apns2.HostDevelopment},
		{"config production default", true, &PushNotification{}, apns2.HostProduction},
		{"config development default", false, &PushNotification{}, apns2.HostDevelopment},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg.Ios.Production = tt.cfgProduction
			got := getApnsClient(cfg, tt.req)
			assert.Equal(t, tt.wantHost, got.Host)
		})
	}
}

// Resolving a client for one request must not change the host of a client
// already handed to another request. Before the fix, getApnsClient returned the
// shared global and mutated its Host in place, so a later development request
// silently repointed an earlier production client at the sandbox.
func TestGetApnsClientNoSharedMutation(t *testing.T) {
	cfg, err := config.LoadConf()
	require.NoError(t, err)
	cfg.Ios.Enabled = true
	cfg.Ios.KeyPath = testKeyPath
	require.NoError(t, InitAPNSClient(context.Background(), cfg))

	prod := getApnsClient(cfg, &PushNotification{Production: true})
	require.Equal(t, apns2.HostProduction, prod.Host)

	// Unrelated later request for the other environment.
	dev := getApnsClient(cfg, &PushNotification{Development: true})

	assert.Equal(t, apns2.HostProduction, prod.Host,
		"production client host was mutated by a later development request")
	assert.Equal(t, apns2.HostDevelopment, dev.Host)
	assert.NotSame(t, ApnsClient, prod, "must not hand back the shared global client")
}

// Concurrent resolution for different environments must not race on the shared
// global. Run with -race to exercise the regression.
func TestGetApnsClientConcurrent(t *testing.T) {
	cfg, err := config.LoadConf()
	require.NoError(t, err)
	cfg.Ios.Enabled = true
	cfg.Ios.KeyPath = testKeyPath
	require.NoError(t, InitAPNSClient(context.Background(), cfg))

	var wg sync.WaitGroup
	for range 50 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			c := getApnsClient(cfg, &PushNotification{Production: true})
			assert.Equal(t, apns2.HostProduction, c.Host)
		}()
		go func() {
			defer wg.Done()
			c := getApnsClient(cfg, &PushNotification{Development: true})
			assert.Equal(t, apns2.HostDevelopment, c.Host)
		}()
	}
	wg.Wait()
}
