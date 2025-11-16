package core

import (
	"encoding/json"
	"errors"
	"strings"
)

// Backward-compatible integer constants kept as-is.
const (
	// PlatFormIos constant is 1 for iOS
	PlatFormIos = iota + 1
	// PlatFormAndroid constant is 2 for Android
	PlatFormAndroid
	// PlatFormHuawei constant is 3 for Huawei
	PlatFormHuawei
)

// Log block string constants (backward-compatible)
const (
	// SucceededPush is log block
	SucceededPush = "succeeded-push"
	// FailedPush is log block
	FailedPush = "failed-push"
)

// Platform is a typed enum for push target platforms.
// This complements the existing integer constants.
type Platform uint8

const (
	// Typed equivalents of platform values.
	PlatformIOS     Platform = Platform(PlatFormIos)
	PlatformAndroid Platform = Platform(PlatFormAndroid)
	PlatformHuawei  Platform = Platform(PlatFormHuawei)
)

// String returns a stable lowercase name for the platform.
func (p Platform) String() string {
	switch p {
	case PlatformIOS:
		return "ios"
	case PlatformAndroid:
		return "android"
	case PlatformHuawei:
		return "huawei"
	default:
		return "unknown"
	}
}

// IsValid reports whether the platform value is supported.
func (p Platform) IsValid() bool {
	return p == PlatformIOS || p == PlatformAndroid || p == PlatformHuawei
}

// ParsePlatform parses a string (case-insensitive) into a Platform.
func ParsePlatform(s string) (Platform, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "ios":
		return PlatformIOS, nil
	case "android":
		return PlatformAndroid, nil
	case "huawei":
		return PlatformHuawei, nil
	default:
		return 0, errors.New("unknown platform: " + s)
	}
}

// MarshalText encodes Platform as its string form.
func (p Platform) MarshalText() ([]byte, error) {
	if !p.IsValid() {
		return nil, errors.New("invalid platform")
	}
	return []byte(p.String()), nil
}

// UnmarshalText decodes Platform from its string form.
func (p *Platform) UnmarshalText(text []byte) error {
	v, err := ParsePlatform(string(text))
	if err != nil {
		return err
	}
	*p = v
	return nil
}

// MarshalJSON encodes Platform as a JSON string.
func (p Platform) MarshalJSON() ([]byte, error) {
	if !p.IsValid() {
		return nil, errors.New("invalid platform")
	}
	return json.Marshal(p.String())
}

// UnmarshalJSON decodes Platform from a JSON string or legacy number.
func (p *Platform) UnmarshalJSON(data []byte) error {
	// Try string first.
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		return p.UnmarshalText([]byte(s))
	}
	// Fallback to legacy numeric values 1/2/3.
	var n int
	if err := json.Unmarshal(data, &n); err == nil {
		switch n {
		case PlatFormIos:
			*p = PlatformIOS
			return nil
		case PlatFormAndroid:
			*p = PlatformAndroid
			return nil
		case PlatFormHuawei:
			*p = PlatformHuawei
			return nil
		}
	}
	return errors.New("invalid platform JSON")
}

// LogBlock is a typed alias for log block kinds.
type LogBlock string

const (
	LogSucceededPush LogBlock = LogBlock(SucceededPush)
	LogFailedPush    LogBlock = LogBlock(FailedPush)
)

// IsValid checks a LogBlock value.
func (l LogBlock) IsValid() bool {
	return l == LogSucceededPush || l == LogFailedPush
}
