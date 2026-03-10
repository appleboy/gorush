package config

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Test file is missing
func TestMissingFile(t *testing.T) {
	filename := "nonexistent_file.yml"
	conf, err := LoadConf(filename)

	assert.Nil(t, conf)
	require.Error(t, err)
	// Check that the error message includes the filename and describes the issue
	assert.Contains(t, err.Error(), "failed to read config file")
	assert.Contains(t, err.Error(), filename)
}

// Test invalid YAML content
func TestInvalidYAMLFile(t *testing.T) {
	// Create a temporary file with invalid YAML
	tmpFile := "test_invalid.yml"
	content := []byte("invalid: yaml: content: [unclosed")

	// Write invalid content to a temporary file
	err := os.WriteFile(tmpFile, content, 0o600)
	require.NoError(t, err)
	defer os.Remove(tmpFile) // Clean up

	conf, err := LoadConf(tmpFile)

	assert.Nil(t, conf)
	require.Error(t, err)
	// Check that the error message includes the filename and describes parsing failure
	assert.Contains(t, err.Error(), "failed to parse config file")
	assert.Contains(t, err.Error(), tmpFile)
}

// Test improved error messages for default config loading
func TestDefaultConfigLoadFailure(t *testing.T) {
	// Backup original defaultConf
	originalDefaultConf := defaultConf
	defer func() {
		defaultConf = originalDefaultConf
	}()

	// Set invalid default config
	defaultConf = []byte(`invalid: yaml: [unclosed`)

	conf, err := LoadConf()

	assert.Nil(t, conf)
	require.Error(t, err)
	// Check that the error message describes the default config loading failure
	assert.Contains(t, err.Error(), "failed to load default config and no config file found")
}

func TestEmptyConfig(t *testing.T) {
	conf, err := LoadConf("testdata/empty.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}

	assert.Equal(t, uint(100), conf.Ios.MaxConcurrentPushes)
}

type ConfigTestSuite struct {
	suite.Suite
	ConfGorushDefault *ConfYaml
	ConfGorush        *ConfYaml
}

func (suite *ConfigTestSuite) SetupTest() {
	var err error
	suite.ConfGorushDefault, err = LoadConf()
	if err != nil {
		panic("failed to load default config.yml")
	}
	suite.ConfGorush, err = LoadConf("testdata/config.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}
}

func (suite *ConfigTestSuite) TestValidateConfDefault() {
	// Core
	suite.Empty(suite.ConfGorushDefault.Core.Address)
	suite.Equal("8088", suite.ConfGorushDefault.Core.Port)
	suite.Equal(int64(30), suite.ConfGorushDefault.Core.ShutdownTimeout)
	suite.True(suite.ConfGorushDefault.Core.Enabled)
	suite.Equal(int64(runtime.NumCPU()), suite.ConfGorushDefault.Core.WorkerNum)
	suite.Equal(int64(8192), suite.ConfGorushDefault.Core.QueueNum)
	suite.Equal("release", suite.ConfGorushDefault.Core.Mode)
	suite.False(suite.ConfGorushDefault.Core.Sync)
	suite.Empty(suite.ConfGorushDefault.Core.FeedbackURL)
	suite.Empty(suite.ConfGorushDefault.Core.FeedbackHeader)
	suite.Equal(int64(10), suite.ConfGorushDefault.Core.FeedbackTimeout)
	suite.False(suite.ConfGorushDefault.Core.SSL)
	suite.Equal("cert.pem", suite.ConfGorushDefault.Core.CertPath)
	suite.Equal("key.pem", suite.ConfGorushDefault.Core.KeyPath)
	suite.Empty(suite.ConfGorushDefault.Core.KeyBase64)
	suite.Empty(suite.ConfGorushDefault.Core.CertBase64)
	suite.Equal(int64(100), suite.ConfGorushDefault.Core.MaxNotification)
	suite.Empty(suite.ConfGorushDefault.Core.HTTPProxy)
	// Pid
	suite.False(suite.ConfGorushDefault.Core.PID.Enabled)
	suite.Equal("gorush.pid", suite.ConfGorushDefault.Core.PID.Path)
	suite.True(suite.ConfGorushDefault.Core.PID.Override)
	suite.False(suite.ConfGorushDefault.Core.AutoTLS.Enabled)
	suite.Equal(".cache", suite.ConfGorushDefault.Core.AutoTLS.Folder)
	suite.Empty(suite.ConfGorushDefault.Core.AutoTLS.Host)

	// Api
	suite.Equal("/api/push", suite.ConfGorushDefault.API.PushURI)
	suite.Equal("/api/stat/go", suite.ConfGorushDefault.API.StatGoURI)
	suite.Equal("/api/stat/app", suite.ConfGorushDefault.API.StatAppURI)
	suite.Equal("/api/config", suite.ConfGorushDefault.API.ConfigURI)
	suite.Equal("/sys/stats", suite.ConfGorushDefault.API.SysStatURI)
	suite.Equal("/metrics", suite.ConfGorushDefault.API.MetricURI)
	suite.Equal("/healthz", suite.ConfGorushDefault.API.HealthURI)

	// Android
	suite.True(suite.ConfGorushDefault.Android.Enabled)
	suite.Empty(suite.ConfGorushDefault.Android.KeyPath)
	suite.Empty(suite.ConfGorushDefault.Android.Credential)
	suite.Equal(0, suite.ConfGorushDefault.Android.MaxRetry)

	// iOS
	suite.False(suite.ConfGorushDefault.Ios.Enabled)
	suite.Empty(suite.ConfGorushDefault.Ios.KeyPath)
	suite.Empty(suite.ConfGorushDefault.Ios.KeyBase64)
	suite.Equal("pem", suite.ConfGorushDefault.Ios.KeyType)
	suite.Empty(suite.ConfGorushDefault.Ios.Password)
	suite.False(suite.ConfGorushDefault.Ios.Production)
	suite.Equal(uint(100), suite.ConfGorushDefault.Ios.MaxConcurrentPushes)
	suite.Equal(0, suite.ConfGorushDefault.Ios.MaxRetry)
	suite.Empty(suite.ConfGorushDefault.Ios.KeyID)
	suite.Empty(suite.ConfGorushDefault.Ios.TeamID)

	// queue
	suite.Equal("local", suite.ConfGorushDefault.Queue.Engine)
	suite.Equal("127.0.0.1:4150", suite.ConfGorushDefault.Queue.NSQ.Addr)
	suite.Equal("gorush", suite.ConfGorushDefault.Queue.NSQ.Topic)
	suite.Equal("gorush", suite.ConfGorushDefault.Queue.NSQ.Channel)

	suite.Equal("127.0.0.1:4222", suite.ConfGorushDefault.Queue.NATS.Addr)
	suite.Equal("gorush", suite.ConfGorushDefault.Queue.NATS.Subj)
	suite.Equal("gorush", suite.ConfGorushDefault.Queue.NATS.Queue)

	suite.Equal("127.0.0.1:6379", suite.ConfGorushDefault.Queue.Redis.Addr)
	suite.Equal("gorush", suite.ConfGorushDefault.Queue.Redis.StreamName)
	suite.Equal("gorush", suite.ConfGorushDefault.Queue.Redis.Group)
	suite.Equal("gorush", suite.ConfGorushDefault.Queue.Redis.Consumer)
	suite.Empty(suite.ConfGorushDefault.Queue.Redis.Username)
	suite.Empty(suite.ConfGorushDefault.Queue.Redis.Password)
	suite.False(suite.ConfGorushDefault.Queue.Redis.WithTLS)
	suite.Equal(0, suite.ConfGorushDefault.Queue.Redis.DB)

	// log
	suite.Equal("string", suite.ConfGorushDefault.Log.Format)
	suite.Equal("stdout", suite.ConfGorushDefault.Log.AccessLog)
	suite.Equal("debug", suite.ConfGorushDefault.Log.AccessLevel)
	suite.Equal("stderr", suite.ConfGorushDefault.Log.ErrorLog)
	suite.Equal("error", suite.ConfGorushDefault.Log.ErrorLevel)
	suite.True(suite.ConfGorushDefault.Log.HideToken)
	suite.False(suite.ConfGorushDefault.Log.HideMessages)

	suite.Equal("memory", suite.ConfGorushDefault.Stat.Engine)
	suite.False(suite.ConfGorushDefault.Stat.Redis.Cluster)
	suite.Equal("localhost:6379", suite.ConfGorushDefault.Stat.Redis.Addr)
	suite.Empty(suite.ConfGorushDefault.Stat.Redis.Username)
	suite.Empty(suite.ConfGorushDefault.Stat.Redis.Password)
	suite.Equal(0, suite.ConfGorushDefault.Stat.Redis.DB)

	suite.Equal("bolt.db", suite.ConfGorushDefault.Stat.BoltDB.Path)
	suite.Equal("gorush", suite.ConfGorushDefault.Stat.BoltDB.Bucket)

	suite.Equal("bunt.db", suite.ConfGorushDefault.Stat.BuntDB.Path)
	suite.Equal("level.db", suite.ConfGorushDefault.Stat.LevelDB.Path)
	suite.Equal("badger.db", suite.ConfGorushDefault.Stat.BadgerDB.Path)

	// gRPC
	suite.False(suite.ConfGorushDefault.GRPC.Enabled)
	suite.Equal("9000", suite.ConfGorushDefault.GRPC.Port)
}

func (suite *ConfigTestSuite) TestValidateConf() {
	// Core
	suite.Equal("8088", suite.ConfGorush.Core.Port)
	suite.Equal(int64(30), suite.ConfGorush.Core.ShutdownTimeout)
	suite.True(suite.ConfGorush.Core.Enabled)
	suite.Equal(int64(runtime.NumCPU()), suite.ConfGorush.Core.WorkerNum)
	suite.Equal(int64(8192), suite.ConfGorush.Core.QueueNum)
	suite.Equal("release", suite.ConfGorush.Core.Mode)
	suite.False(suite.ConfGorush.Core.Sync)
	suite.Empty(suite.ConfGorush.Core.FeedbackURL)
	suite.Equal(int64(10), suite.ConfGorush.Core.FeedbackTimeout)
	suite.Len(suite.ConfGorush.Core.FeedbackHeader, 1)
	suite.Equal(
		"x-gorush-token:4e989115e09680f44a645519fed6a976",
		suite.ConfGorush.Core.FeedbackHeader[0],
	)
	suite.False(suite.ConfGorush.Core.SSL)
	suite.Equal("cert.pem", suite.ConfGorush.Core.CertPath)
	suite.Equal("key.pem", suite.ConfGorush.Core.KeyPath)
	suite.Empty(suite.ConfGorush.Core.CertBase64)
	suite.Empty(suite.ConfGorush.Core.KeyBase64)
	suite.Equal(int64(100), suite.ConfGorush.Core.MaxNotification)
	suite.Empty(suite.ConfGorush.Core.HTTPProxy)
	// Pid
	suite.False(suite.ConfGorush.Core.PID.Enabled)
	suite.Equal("gorush.pid", suite.ConfGorush.Core.PID.Path)
	suite.True(suite.ConfGorush.Core.PID.Override)
	suite.False(suite.ConfGorush.Core.AutoTLS.Enabled)
	suite.Equal(".cache", suite.ConfGorush.Core.AutoTLS.Folder)
	suite.Empty(suite.ConfGorush.Core.AutoTLS.Host)

	// Api
	suite.Equal("/api/push", suite.ConfGorush.API.PushURI)
	suite.Equal("/api/stat/go", suite.ConfGorush.API.StatGoURI)
	suite.Equal("/api/stat/app", suite.ConfGorush.API.StatAppURI)
	suite.Equal("/api/config", suite.ConfGorush.API.ConfigURI)
	suite.Equal("/sys/stats", suite.ConfGorush.API.SysStatURI)
	suite.Equal("/metrics", suite.ConfGorush.API.MetricURI)
	suite.Equal("/healthz", suite.ConfGorush.API.HealthURI)

	// Android
	suite.True(suite.ConfGorush.Android.Enabled)
	suite.Equal("key.json", suite.ConfGorush.Android.KeyPath)
	suite.Equal("CREDENTIAL_JSON_DATA", suite.ConfGorush.Android.Credential)
	suite.Equal(0, suite.ConfGorush.Android.MaxRetry)

	// iOS
	suite.False(suite.ConfGorush.Ios.Enabled)
	suite.Equal("key.pem", suite.ConfGorush.Ios.KeyPath)
	suite.Empty(suite.ConfGorush.Ios.KeyBase64)
	suite.Equal("pem", suite.ConfGorush.Ios.KeyType)
	suite.Empty(suite.ConfGorush.Ios.Password)
	suite.False(suite.ConfGorush.Ios.Production)
	suite.Equal(uint(100), suite.ConfGorush.Ios.MaxConcurrentPushes)
	suite.Equal(0, suite.ConfGorush.Ios.MaxRetry)
	suite.Empty(suite.ConfGorush.Ios.KeyID)
	suite.Empty(suite.ConfGorush.Ios.TeamID)

	// log
	suite.Equal("string", suite.ConfGorush.Log.Format)
	suite.Equal("stdout", suite.ConfGorush.Log.AccessLog)
	suite.Equal("debug", suite.ConfGorush.Log.AccessLevel)
	suite.Equal("stderr", suite.ConfGorush.Log.ErrorLog)
	suite.Equal("error", suite.ConfGorush.Log.ErrorLevel)
	suite.True(suite.ConfGorush.Log.HideToken)

	suite.Equal("memory", suite.ConfGorush.Stat.Engine)
	suite.False(suite.ConfGorush.Stat.Redis.Cluster)
	suite.Equal("localhost:6379", suite.ConfGorush.Stat.Redis.Addr)
	suite.Empty(suite.ConfGorush.Stat.Redis.Username)
	suite.Empty(suite.ConfGorush.Stat.Redis.Password)
	suite.Equal(0, suite.ConfGorush.Stat.Redis.DB)

	suite.Equal("bolt.db", suite.ConfGorush.Stat.BoltDB.Path)
	suite.Equal("gorush", suite.ConfGorush.Stat.BoltDB.Bucket)

	suite.Equal("bunt.db", suite.ConfGorush.Stat.BuntDB.Path)
	suite.Equal("level.db", suite.ConfGorush.Stat.LevelDB.Path)
	suite.Equal("badger.db", suite.ConfGorush.Stat.BadgerDB.Path)

	// gRPC
	suite.False(suite.ConfGorush.GRPC.Enabled)
	suite.Equal("9000", suite.ConfGorush.GRPC.Port)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func TestLoadConfigFromEnv(t *testing.T) {
	t.Setenv("GORUSH_CORE_PORT", "9001")
	t.Setenv("GORUSH_GRPC_ENABLED", "true")
	t.Setenv("GORUSH_CORE_MAX_NOTIFICATION", "200")
	t.Setenv("GORUSH_IOS_KEY_ID", "ABC123DEFG")
	t.Setenv("GORUSH_IOS_TEAM_ID", "DEF123GHIJ")
	t.Setenv("GORUSH_API_HEALTH_URI", "/healthz")
	t.Setenv("GORUSH_CORE_FEEDBACK_HOOK_URL", "http://example.com")
	t.Setenv("GORUSH_CORE_FEEDBACK_HEADER", "x-api-key:1234567890 x-auth-key:0987654321")

	ConfGorush, err := LoadConf("testdata/config.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}
	assert.Equal(t, "9001", ConfGorush.Core.Port)
	assert.Equal(t, int64(200), ConfGorush.Core.MaxNotification)
	assert.True(t, ConfGorush.GRPC.Enabled)
	assert.Equal(t, "ABC123DEFG", ConfGorush.Ios.KeyID)
	assert.Equal(t, "DEF123GHIJ", ConfGorush.Ios.TeamID)
	assert.Equal(t, "/healthz", ConfGorush.API.HealthURI)
	assert.Equal(t, "http://example.com", ConfGorush.Core.FeedbackURL)
	assert.Equal(t, "x-api-key:1234567890", ConfGorush.Core.FeedbackHeader[0])
	assert.Equal(t, "x-auth-key:0987654321", ConfGorush.Core.FeedbackHeader[1])
}

func TestRedisDBConfiguration(t *testing.T) {
	// Test loading Redis DB configuration from file
	conf, err := LoadConf("testdata/redis_db_config.yml")
	if err != nil {
		t.Fatalf("failed to load redis_db_config.yml: %v", err)
	}

	// Test queue.redis.db is properly loaded
	assert.Equal(t, "redis", conf.Queue.Engine)
	assert.Equal(t, 5, conf.Queue.Redis.DB)

	// Test stat.redis.db is properly loaded
	assert.Equal(t, "redis", conf.Stat.Engine)
	assert.Equal(t, 3, conf.Stat.Redis.DB)
}

func TestRedisDBConfigurationFromEnv(t *testing.T) {
	// Test loading Redis DB configuration from environment variables
	t.Setenv("GORUSH_QUEUE_REDIS_DB", "7")
	t.Setenv("GORUSH_STAT_REDIS_DB", "9")

	conf, err := LoadConf()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Test queue.redis.db is properly loaded from env
	assert.Equal(t, 7, conf.Queue.Redis.DB)

	// Test stat.redis.db is properly loaded from env
	assert.Equal(t, 9, conf.Stat.Redis.DB)
}

func TestLoadWrongDefaultYAMLConfig(t *testing.T) {
	// Backup original defaultConf
	originalDefaultConf := defaultConf
	defer func() {
		defaultConf = originalDefaultConf
	}()

	defaultConf = []byte(`a`)
	_, err := LoadConf()
	assert.Error(t, err)
}

func TestValidatePort(t *testing.T) {
	tests := []struct {
		name    string
		port    string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty port should be valid",
			port:    "",
			wantErr: false,
		},
		{
			name:    "valid port 80",
			port:    "80",
			wantErr: false,
		},
		{
			name:    "valid port 8080",
			port:    "8080",
			wantErr: false,
		},
		{
			name:    "valid port 65535",
			port:    "65535",
			wantErr: false,
		},
		{
			name:    "valid port 1",
			port:    "1",
			wantErr: false,
		},
		{
			name:    "invalid port 0",
			port:    "0",
			wantErr: true,
			errMsg:  "port out of range",
		},
		{
			name:    "invalid port 65536",
			port:    "65536",
			wantErr: true,
			errMsg:  "port out of range",
		},
		{
			name:    "invalid port format",
			port:    "abc",
			wantErr: true,
			errMsg:  "invalid port format",
		},
		{
			name:    "invalid port with injection",
			port:    "80;rm -rf /",
			wantErr: true,
			errMsg:  "invalid port format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePort(tt.port)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ValidatePort() expected error but got none")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ValidatePort() error = %v, want error containing %v", err, tt.errMsg)
				}
			} else if err != nil {
				t.Errorf("ValidatePort() error = %v, want nil", err)
			}
		})
	}
}

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name    string
		addr    string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty address should be valid",
			addr:    "",
			wantErr: false,
		},
		{
			name:    "valid IPv4 localhost",
			addr:    "127.0.0.1",
			wantErr: false,
		},
		{
			name:    "valid IPv4 all interfaces",
			addr:    "0.0.0.0",
			wantErr: false,
		},
		{
			name:    "valid IPv6 localhost",
			addr:    "::1",
			wantErr: false,
		},
		{
			name:    "valid hostname",
			addr:    "localhost",
			wantErr: false,
		},
		{
			name:    "invalid address too long",
			addr:    strings.Repeat("a", 254),
			wantErr: true,
			errMsg:  "invalid address format",
		},
		{
			name:    "invalid address with double dots",
			addr:    "test..example.com",
			wantErr: true,
			errMsg:  "invalid address format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAddress(tt.addr)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ValidateAddress() expected error but got none")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf(
						"ValidateAddress() error = %v, want error containing %v",
						err,
						tt.errMsg,
					)
				}
			} else if err != nil {
				t.Errorf("ValidateAddress() error = %v, want nil", err)
			}
		})
	}
}

func TestValidatePIDPath(t *testing.T) {
	tests := []struct {
		name    string
		pidPath string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty path should be valid",
			pidPath: "",
			wantErr: false,
		},
		{
			name:    "valid relative path",
			pidPath: "gorush.pid",
			wantErr: false,
		},
		{
			name:    "valid absolute path in tmp",
			pidPath: "/tmp/gorush.pid",
			wantErr: false,
		},
		{
			name:    "valid absolute path with .. components (cleaned)",
			pidPath: "/tmp/foo/../gorush.pid",
			wantErr: false,
		},
		{
			name:    "valid absolute path with multiple .. components",
			pidPath: "/home/user/subdir/../gorush.pid",
			wantErr: false,
		},
		{
			name:    "relative path traversal attack",
			pidPath: "../../../etc/passwd",
			wantErr: true,
			errMsg:  "path traversal detected",
		},
		{
			name:    "simple relative traversal",
			pidPath: "../gorush.pid",
			wantErr: true,
			errMsg:  "path traversal detected",
		},
		{
			name:    "complex relative traversal",
			pidPath: "subdir/../../gorush.pid",
			wantErr: true,
			errMsg:  "path traversal detected",
		},
		{
			name:    "attempt to write to /etc",
			pidPath: "/etc/gorush.pid",
			wantErr: true,
			errMsg:  "sensitive directory",
		},
		{
			name:    "attempt to write to /usr",
			pidPath: "/usr/bin/gorush.pid",
			wantErr: true,
			errMsg:  "sensitive directory",
		},
		{
			name:    "attempt to write to /var",
			pidPath: "/var/log/gorush.pid",
			wantErr: true,
			errMsg:  "sensitive directory",
		},
		{
			name:    "attempt to write to /sys",
			pidPath: "/sys/gorush.pid",
			wantErr: true,
			errMsg:  "sensitive directory",
		},
		{
			name:    "attempt to write to /proc",
			pidPath: "/proc/gorush.pid",
			wantErr: true,
			errMsg:  "sensitive directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePIDPath(tt.pidPath)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ValidatePIDPath() expected error but got none")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf(
						"ValidatePIDPath() error = %v, want error containing %v",
						err,
						tt.errMsg,
					)
				}
			} else if err != nil {
				t.Errorf("ValidatePIDPath() error = %v, want nil", err)
			}
		})
	}
}

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *ConfYaml
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid config",
			cfg: &ConfYaml{
				Core: SectionCore{
					Port:    "8088",
					Address: "0.0.0.0",
				},
				Stat: SectionStat{
					Engine: "memory",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid port in config",
			cfg: &ConfYaml{
				Core: SectionCore{
					Port: "99999",
				},
			},
			wantErr: true,
			errMsg:  "invalid core port",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateConfig(tt.cfg)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ValidateConfig() expected error but got none")
					return
				}

				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf(
						"ValidateConfig() error = %v, want error containing %v",
						err,
						tt.errMsg,
					)
				}
			} else if err != nil {
				t.Errorf("ValidateConfig() error = %v, want nil", err)
			}
		})
	}
}

// Benchmark tests for security validation functions
func BenchmarkValidatePort(b *testing.B) {
	for b.Loop() {
		_ = ValidatePort("8080")
	}
}

func BenchmarkValidateAddress(b *testing.B) {
	for b.Loop() {
		_ = ValidateAddress("127.0.0.1")
	}
}

func BenchmarkValidatePIDPath(b *testing.B) {
	for b.Loop() {
		_ = ValidatePIDPath("/tmp/gorush.pid")
	}
}

// Integration test for security validation
func TestSecurityValidationIntegration(t *testing.T) {
	// Test that all validation functions work together
	t.Run("complete security validation", func(t *testing.T) {
		// Test valid inputs
		if err := ValidatePort("8088"); err != nil {
			t.Errorf("Valid port should not error: %v", err)
		}

		if err := ValidateAddress("127.0.0.1"); err != nil {
			t.Errorf("Valid address should not error: %v", err)
		}

		if err := ValidatePIDPath("/tmp/test.pid"); err != nil {
			t.Errorf("Valid PID path should not error: %v", err)
		}

		// Test malicious inputs
		if err := ValidatePort("8080; rm -rf /"); err == nil {
			t.Error("Malicious port should error")
		}

		if err := ValidatePIDPath("../../../etc/passwd"); err == nil {
			t.Error("Path traversal should error")
		}

		if err := ValidatePIDPath("/etc/malicious.pid"); err == nil {
			t.Error("Sensitive directory write should error")
		}
	})
}

func TestAPIDefaultsFromEnv(t *testing.T) {
	// Test that API endpoint defaults can be overridden by environment variables
	t.Setenv("GORUSH_API_PUSH_URI", "/custom/push")
	t.Setenv("GORUSH_API_STAT_GO_URI", "/custom/stat/go")
	t.Setenv("GORUSH_API_METRIC_URI", "/custom/metrics")

	conf, err := LoadConf()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	assert.Equal(t, "/custom/push", conf.API.PushURI)
	assert.Equal(t, "/custom/stat/go", conf.API.StatGoURI)
	assert.Equal(t, "/custom/metrics", conf.API.MetricURI)
}

func TestLogDefaultsFromEnv(t *testing.T) {
	// Test that log level defaults can be overridden by environment variables
	t.Setenv("GORUSH_LOG_ACCESS_LEVEL", "info")
	t.Setenv("GORUSH_LOG_ERROR_LEVEL", "warn")
	t.Setenv("GORUSH_LOG_FORMAT", "json")

	conf, err := LoadConf()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	assert.Equal(t, "info", conf.Log.AccessLevel)
	assert.Equal(t, "warn", conf.Log.ErrorLevel)
	assert.Equal(t, "json", conf.Log.Format)
}

func TestLogLevelDefaultsWhenEmpty(t *testing.T) {
	// Test that when no log level is specified, defaults are used
	// This was the original bug - empty log levels caused initialization to fail
	conf, err := LoadConf()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Verify defaults are set
	assert.NotEmpty(t, conf.Log.AccessLevel, "access level should have default value")
	assert.NotEmpty(t, conf.Log.ErrorLevel, "error level should have default value")
	assert.Equal(t, "debug", conf.Log.AccessLevel)
	assert.Equal(t, "error", conf.Log.ErrorLevel)
}

func TestAllAPIEndpointsHaveDefaults(t *testing.T) {
	// Test that all API endpoints have proper defaults
	conf, err := LoadConf()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Verify all API endpoints have non-empty defaults
	assert.NotEmpty(t, conf.API.PushURI, "push_uri should have default")
	assert.NotEmpty(t, conf.API.StatGoURI, "stat_go_uri should have default")
	assert.NotEmpty(t, conf.API.StatAppURI, "stat_app_uri should have default")
	assert.NotEmpty(t, conf.API.ConfigURI, "config_uri should have default")
	assert.NotEmpty(t, conf.API.SysStatURI, "sys_stat_uri should have default")
	assert.NotEmpty(t, conf.API.MetricURI, "metric_uri should have default")
	assert.NotEmpty(t, conf.API.HealthURI, "health_uri should have default")

	// Verify correct default values
	assert.Equal(t, "/api/push", conf.API.PushURI)
	assert.Equal(t, "/api/stat/go", conf.API.StatGoURI)
	assert.Equal(t, "/api/stat/app", conf.API.StatAppURI)
	assert.Equal(t, "/api/config", conf.API.ConfigURI)
	assert.Equal(t, "/sys/stats", conf.API.SysStatURI)
	assert.Equal(t, "/metrics", conf.API.MetricURI)
	assert.Equal(t, "/healthz", conf.API.HealthURI)
}

func TestSanitizedCopy(t *testing.T) {
	cfg := &ConfYaml{}
	cfg.Core.CertBase64 = "cert-data"
	cfg.Core.KeyBase64 = "key-data"
	cfg.Core.HTTPProxy = "http://proxy:8080"
	cfg.Android.KeyPath = "/path/to/key.json"
	cfg.Android.Credential = "fcm-credential"
	cfg.Huawei.AppSecret = "hms-secret"
	cfg.Ios.KeyPath = "/path/to/ios.p8"
	cfg.Ios.KeyBase64 = "ios-key-data"
	cfg.Ios.Password = "ios-password"
	cfg.Ios.KeyID = "ABCDE12345"
	cfg.Ios.TeamID = "TEAM123456"
	cfg.Queue.Redis.Username = "queue-user"
	cfg.Queue.Redis.Password = "queue-pass"
	cfg.Stat.Redis.Username = "stat-user"
	cfg.Stat.Redis.Password = "stat-pass"

	// Set a non-sensitive field to verify it's preserved
	cfg.Core.Port = "8088"

	sanitized := cfg.SanitizedCopy()

	// All sensitive fields should be redacted
	assert.Equal(t, "[REDACTED]", sanitized.Core.CertBase64)
	assert.Equal(t, "[REDACTED]", sanitized.Core.KeyBase64)
	assert.Equal(t, "[REDACTED]", sanitized.Core.HTTPProxy)
	assert.Equal(t, "[REDACTED]", sanitized.Android.KeyPath)
	assert.Equal(t, "[REDACTED]", sanitized.Android.Credential)
	assert.Equal(t, "[REDACTED]", sanitized.Huawei.AppSecret)
	assert.Equal(t, "[REDACTED]", sanitized.Ios.KeyPath)
	assert.Equal(t, "[REDACTED]", sanitized.Ios.KeyBase64)
	assert.Equal(t, "[REDACTED]", sanitized.Ios.Password)
	assert.Equal(t, "[REDACTED]", sanitized.Ios.KeyID)
	assert.Equal(t, "[REDACTED]", sanitized.Ios.TeamID)
	assert.Equal(t, "[REDACTED]", sanitized.Queue.Redis.Username)
	assert.Equal(t, "[REDACTED]", sanitized.Queue.Redis.Password)
	assert.Equal(t, "[REDACTED]", sanitized.Stat.Redis.Username)
	assert.Equal(t, "[REDACTED]", sanitized.Stat.Redis.Password)

	// Non-sensitive fields should be preserved
	assert.Equal(t, "8088", sanitized.Core.Port)

	// Original config must NOT be modified
	assert.Equal(t, "cert-data", cfg.Core.CertBase64)
	assert.Equal(t, "fcm-credential", cfg.Android.Credential)
	assert.Equal(t, "ios-password", cfg.Ios.Password)
	assert.Equal(t, "stat-pass", cfg.Stat.Redis.Password)
}

func TestSanitizedCopyEmptyFields(t *testing.T) {
	cfg := &ConfYaml{}
	// All sensitive fields are empty by default

	sanitized := cfg.SanitizedCopy()

	// Empty fields should remain empty, not become "[REDACTED]"
	assert.Empty(t, sanitized.Core.CertBase64)
	assert.Empty(t, sanitized.Core.KeyBase64)
	assert.Empty(t, sanitized.Core.HTTPProxy)
	assert.Empty(t, sanitized.Android.KeyPath)
	assert.Empty(t, sanitized.Android.Credential)
	assert.Empty(t, sanitized.Huawei.AppSecret)
	assert.Empty(t, sanitized.Ios.KeyPath)
	assert.Empty(t, sanitized.Ios.KeyBase64)
	assert.Empty(t, sanitized.Ios.Password)
	assert.Empty(t, sanitized.Ios.KeyID)
	assert.Empty(t, sanitized.Ios.TeamID)
	assert.Empty(t, sanitized.Queue.Redis.Username)
	assert.Empty(t, sanitized.Queue.Redis.Password)
	assert.Empty(t, sanitized.Stat.Redis.Username)
	assert.Empty(t, sanitized.Stat.Redis.Password)
}
