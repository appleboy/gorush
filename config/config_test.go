package config

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Test file is missing
func TestMissingFile(t *testing.T) {
	filename := "nonexistent_file.yml"
	conf, err := LoadConf(filename)

	assert.Nil(t, conf)
	assert.NotNil(t, err)
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
	err := os.WriteFile(tmpFile, content, 0o644)
	assert.NoError(t, err)
	defer os.Remove(tmpFile) // Clean up

	conf, err := LoadConf(tmpFile)

	assert.Nil(t, conf)
	assert.NotNil(t, err)
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
	assert.NotNil(t, err)
	// Check that the error message describes the default config loading failure
	assert.Contains(t, err.Error(), "failed to load default config and no config file found")
}

func TestEmptyConfig(t *testing.T) {
	conf, err := LoadConf("testdata/empty.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}

	assert.Equal(t, uint(DefaultMaxConcurrentPush), conf.Ios.MaxConcurrentPushes)
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
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.Address)
	assert.Equal(suite.T(), DefaultPort, suite.ConfGorushDefault.Core.Port)
	assert.Equal(suite.T(), int64(DefaultShutdownTimeout), suite.ConfGorushDefault.Core.ShutdownTimeout)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Core.Enabled)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorushDefault.Core.WorkerNum)
	assert.Equal(suite.T(), int64(DefaultQueueNum), suite.ConfGorushDefault.Core.QueueNum)
	assert.Equal(suite.T(), "release", suite.ConfGorushDefault.Core.Mode)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.Sync)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.FeedbackURL)
	assert.Equal(suite.T(), 0, len(suite.ConfGorushDefault.Core.FeedbackHeader))
	assert.Equal(suite.T(), int64(DefaultFeedbackTimeout), suite.ConfGorushDefault.Core.FeedbackTimeout)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.SSL)
	assert.Equal(suite.T(), "cert.pem", suite.ConfGorushDefault.Core.CertPath)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorushDefault.Core.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.KeyBase64)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.CertBase64)
	assert.Equal(suite.T(), int64(DefaultMaxNotification), suite.ConfGorushDefault.Core.MaxNotification)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.HTTPProxy)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.PID.Enabled)
	assert.Equal(suite.T(), "gorush.pid", suite.ConfGorushDefault.Core.PID.Path)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Core.PID.Override)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.AutoTLS.Enabled)
	assert.Equal(suite.T(), ".cache", suite.ConfGorushDefault.Core.AutoTLS.Folder)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.AutoTLS.Host)

	// Api
	assert.Equal(suite.T(), "/api/push", suite.ConfGorushDefault.API.PushURI)
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorushDefault.API.StatGoURI)
	assert.Equal(suite.T(), "/api/stat/app", suite.ConfGorushDefault.API.StatAppURI)
	assert.Equal(suite.T(), "/api/config", suite.ConfGorushDefault.API.ConfigURI)
	assert.Equal(suite.T(), "/sys/stats", suite.ConfGorushDefault.API.SysStatURI)
	assert.Equal(suite.T(), "/metrics", suite.ConfGorushDefault.API.MetricURI)
	assert.Equal(suite.T(), "/healthz", suite.ConfGorushDefault.API.HealthURI)

	// Android
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Android.Enabled)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Android.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Android.Credential)
	assert.Equal(suite.T(), 0, suite.ConfGorushDefault.Android.MaxRetry)

	// iOS
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Ios.Enabled)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.KeyBase64)
	assert.Equal(suite.T(), "pem", suite.ConfGorushDefault.Ios.KeyType)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.Password)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Ios.Production)
	assert.Equal(suite.T(), uint(DefaultMaxConcurrentPush), suite.ConfGorushDefault.Ios.MaxConcurrentPushes)
	assert.Equal(suite.T(), 0, suite.ConfGorushDefault.Ios.MaxRetry)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.KeyID)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.TeamID)

	// queue
	assert.Equal(suite.T(), "local", suite.ConfGorushDefault.Queue.Engine)
	assert.Equal(suite.T(), "127.0.0.1:4150", suite.ConfGorushDefault.Queue.NSQ.Addr)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.NSQ.Topic)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.NSQ.Channel)

	assert.Equal(suite.T(), "127.0.0.1:4222", suite.ConfGorushDefault.Queue.NATS.Addr)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.NATS.Subj)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.NATS.Queue)

	assert.Equal(suite.T(), "127.0.0.1:6379", suite.ConfGorushDefault.Queue.Redis.Addr)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.Redis.StreamName)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.Redis.Group)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.Redis.Consumer)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Queue.Redis.Username)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Queue.Redis.Password)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Queue.Redis.WithTLS)

	// log
	assert.Equal(suite.T(), "string", suite.ConfGorushDefault.Log.Format)
	assert.Equal(suite.T(), "stdout", suite.ConfGorushDefault.Log.AccessLog)
	assert.Equal(suite.T(), "debug", suite.ConfGorushDefault.Log.AccessLevel)
	assert.Equal(suite.T(), "stderr", suite.ConfGorushDefault.Log.ErrorLog)
	assert.Equal(suite.T(), "error", suite.ConfGorushDefault.Log.ErrorLevel)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Log.HideToken)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Log.HideMessages)

	assert.Equal(suite.T(), "memory", suite.ConfGorushDefault.Stat.Engine)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Stat.Redis.Cluster)
	assert.Equal(suite.T(), "localhost:6379", suite.ConfGorushDefault.Stat.Redis.Addr)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Stat.Redis.Username)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Stat.Redis.Password)
	assert.Equal(suite.T(), 0, suite.ConfGorushDefault.Stat.Redis.DB)

	assert.Equal(suite.T(), "bolt.db", suite.ConfGorushDefault.Stat.BoltDB.Path)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Stat.BoltDB.Bucket)

	assert.Equal(suite.T(), "bunt.db", suite.ConfGorushDefault.Stat.BuntDB.Path)
	assert.Equal(suite.T(), "level.db", suite.ConfGorushDefault.Stat.LevelDB.Path)
	assert.Equal(suite.T(), "badger.db", suite.ConfGorushDefault.Stat.BadgerDB.Path)

	// gRPC
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.GRPC.Enabled)
	assert.Equal(suite.T(), DefaultGRPCPort, suite.ConfGorushDefault.GRPC.Port)
}

func (suite *ConfigTestSuite) TestValidateConf() {
	// Core
	assert.Equal(suite.T(), DefaultPort, suite.ConfGorush.Core.Port)
	assert.Equal(suite.T(), int64(DefaultShutdownTimeout), suite.ConfGorush.Core.ShutdownTimeout)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.Enabled)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorush.Core.WorkerNum)
	assert.Equal(suite.T(), int64(DefaultQueueNum), suite.ConfGorush.Core.QueueNum)
	assert.Equal(suite.T(), "release", suite.ConfGorush.Core.Mode)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.Sync)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.FeedbackURL)
	assert.Equal(suite.T(), int64(DefaultFeedbackTimeout), suite.ConfGorush.Core.FeedbackTimeout)
	assert.Equal(suite.T(), 1, len(suite.ConfGorush.Core.FeedbackHeader))
	assert.Equal(suite.T(), "x-gorush-token:4e989115e09680f44a645519fed6a976", suite.ConfGorush.Core.FeedbackHeader[0])
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.SSL)
	assert.Equal(suite.T(), "cert.pem", suite.ConfGorush.Core.CertPath)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Core.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.CertBase64)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.KeyBase64)
	assert.Equal(suite.T(), int64(DefaultMaxNotification), suite.ConfGorush.Core.MaxNotification)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.HTTPProxy)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.PID.Enabled)
	assert.Equal(suite.T(), "gorush.pid", suite.ConfGorush.Core.PID.Path)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.PID.Override)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.AutoTLS.Enabled)
	assert.Equal(suite.T(), ".cache", suite.ConfGorush.Core.AutoTLS.Folder)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.AutoTLS.Host)

	// Api
	assert.Equal(suite.T(), "/api/push", suite.ConfGorush.API.PushURI)
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorush.API.StatGoURI)
	assert.Equal(suite.T(), "/api/stat/app", suite.ConfGorush.API.StatAppURI)
	assert.Equal(suite.T(), "/api/config", suite.ConfGorush.API.ConfigURI)
	assert.Equal(suite.T(), "/sys/stats", suite.ConfGorush.API.SysStatURI)
	assert.Equal(suite.T(), "/metrics", suite.ConfGorush.API.MetricURI)
	assert.Equal(suite.T(), "/healthz", suite.ConfGorush.API.HealthURI)

	// Android
	assert.Equal(suite.T(), true, suite.ConfGorush.Android.Enabled)
	assert.Equal(suite.T(), "key.json", suite.ConfGorush.Android.KeyPath)
	assert.Equal(suite.T(), "CREDENTIAL_JSON_DATA", suite.ConfGorush.Android.Credential)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Android.MaxRetry)

	// iOS
	assert.Equal(suite.T(), false, suite.ConfGorush.Ios.Enabled)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Ios.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.KeyBase64)
	assert.Equal(suite.T(), "pem", suite.ConfGorush.Ios.KeyType)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.Password)
	assert.Equal(suite.T(), false, suite.ConfGorush.Ios.Production)
	assert.Equal(suite.T(), uint(DefaultMaxConcurrentPush), suite.ConfGorush.Ios.MaxConcurrentPushes)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Ios.MaxRetry)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.KeyID)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.TeamID)

	// log
	assert.Equal(suite.T(), "string", suite.ConfGorush.Log.Format)
	assert.Equal(suite.T(), "stdout", suite.ConfGorush.Log.AccessLog)
	assert.Equal(suite.T(), "debug", suite.ConfGorush.Log.AccessLevel)
	assert.Equal(suite.T(), "stderr", suite.ConfGorush.Log.ErrorLog)
	assert.Equal(suite.T(), "error", suite.ConfGorush.Log.ErrorLevel)
	assert.Equal(suite.T(), true, suite.ConfGorush.Log.HideToken)

	assert.Equal(suite.T(), "memory", suite.ConfGorush.Stat.Engine)
	assert.Equal(suite.T(), false, suite.ConfGorush.Stat.Redis.Cluster)
	assert.Equal(suite.T(), "localhost:6379", suite.ConfGorush.Stat.Redis.Addr)
	assert.Equal(suite.T(), "", suite.ConfGorush.Stat.Redis.Username)
	assert.Equal(suite.T(), "", suite.ConfGorush.Stat.Redis.Password)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Stat.Redis.DB)

	assert.Equal(suite.T(), "bolt.db", suite.ConfGorush.Stat.BoltDB.Path)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Stat.BoltDB.Bucket)

	assert.Equal(suite.T(), "bunt.db", suite.ConfGorush.Stat.BuntDB.Path)
	assert.Equal(suite.T(), "level.db", suite.ConfGorush.Stat.LevelDB.Path)
	assert.Equal(suite.T(), "badger.db", suite.ConfGorush.Stat.BadgerDB.Path)

	// gRPC
	assert.Equal(suite.T(), false, suite.ConfGorush.GRPC.Enabled)
	assert.Equal(suite.T(), DefaultGRPCPort, suite.ConfGorush.GRPC.Port)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("GORUSH_CORE_PORT", "9001")
	os.Setenv("GORUSH_GRPC_ENABLED", "true")
	os.Setenv("GORUSH_CORE_MAX_NOTIFICATION", "200")
	os.Setenv("GORUSH_IOS_KEY_ID", "ABC123DEFG")
	os.Setenv("GORUSH_IOS_TEAM_ID", "DEF123GHIJ")
	os.Setenv("GORUSH_API_HEALTH_URI", "/healthz")
	os.Setenv("GORUSH_CORE_FEEDBACK_HOOK_URL", "http://example.com")
	os.Setenv("GORUSH_CORE_FEEDBACK_HEADER", "x-api-key:1234567890 x-auth-key:0987654321")
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

func TestLoadWrongDefaultYAMLConfig(t *testing.T) {
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
					t.Errorf("ValidateAddress() error = %v, want error containing %v", err, tt.errMsg)
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
					t.Errorf("ValidatePIDPath() error = %v, want error containing %v", err, tt.errMsg)
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
					Port:    DefaultPort,
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
					t.Errorf("ValidateConfig() error = %v, want error containing %v", err, tt.errMsg)
				}
			} else if err != nil {
				t.Errorf("ValidateConfig() error = %v, want nil", err)
			}
		})
	}
}

// Benchmark tests for security validation functions
func BenchmarkValidatePort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ValidatePort("8080")
	}
}

func BenchmarkValidateAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ValidateAddress("127.0.0.1")
	}
}

func BenchmarkValidatePIDPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ValidatePIDPath("/tmp/gorush.pid")
	}
}

// Integration test for security validation
func TestSecurityValidationIntegration(t *testing.T) {
	// Test that all validation functions work together
	t.Run("complete security validation", func(t *testing.T) {
		// Test valid inputs
		if err := ValidatePort(DefaultPort); err != nil {
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
