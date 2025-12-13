package core

import (
	"encoding/json"
	"testing"
)

func TestPlatformStringAndValidity(t *testing.T) {
	cases := []struct {
		p        Platform
		expected string
	}{
		{PlatformIOS, "ios"},
		{PlatformAndroid, "android"},
		{PlatformHuawei, "huawei"},
	}

	for _, c := range cases {
		if !c.p.IsValid() {
			t.Fatalf("expected %v to be valid", c.p)
		}
		if got := c.p.String(); got != c.expected {
			t.Fatalf("String() mismatch: got %q want %q", got, c.expected)
		}
	}

	if Platform(0).IsValid() {
		t.Fatalf("expected zero value to be invalid")
	}
}

func TestParsePlatform(t *testing.T) {
	p, err := ParsePlatform(" IOS ")
	if err != nil || p != PlatformIOS {
		t.Fatalf("ParsePlatform ios failed: p=%v err=%v", p, err)
	}
	p, err = ParsePlatform("android")
	if err != nil || p != PlatformAndroid {
		t.Fatalf("ParsePlatform android failed: p=%v err=%v", p, err)
	}
	p, err = ParsePlatform("HuAwEi")
	if err != nil || p != PlatformHuawei {
		t.Fatalf("ParsePlatform huawei failed: p=%v err=%v", p, err)
	}
	if _, err = ParsePlatform("unknown"); err == nil {
		t.Fatalf("expected error for unknown platform")
	}
}

func TestPlatformTextMarshaling(t *testing.T) {
	b, err := PlatformIOS.MarshalText()
	if err != nil || string(b) != "ios" {
		t.Fatalf("MarshalText: got %q err=%v", string(b), err)
	}
	var p Platform
	if err := p.UnmarshalText([]byte("android")); err != nil || p != PlatformAndroid {
		t.Fatalf("UnmarshalText: p=%v err=%v", p, err)
	}
}

func TestPlatformJSONMarshaling(t *testing.T) {
	b, err := json.Marshal(PlatformHuawei)
	if err != nil || string(b) != "\"huawei\"" {
		t.Fatalf("MarshalJSON: got %s err=%v", string(b), err)
	}
	var p Platform
	if err := json.Unmarshal([]byte("\"ios\""), &p); err != nil || p != PlatformIOS {
		t.Fatalf("UnmarshalJSON string: p=%v err=%v", p, err)
	}
	// legacy numeric values
	if err := json.Unmarshal([]byte("2"), &p); err != nil || p != PlatformAndroid {
		t.Fatalf("UnmarshalJSON legacy number: p=%v err=%v", p, err)
	}
	if err := json.Unmarshal([]byte("99"), &p); err == nil {
		t.Fatalf("expected error for invalid numeric JSON")
	}
}

func TestLogBlock(t *testing.T) {
	if !LogSucceededPush.IsValid() || !LogFailedPush.IsValid() {
		t.Fatalf("expected log blocks to be valid")
	}
	if LogBlock("x").IsValid() {
		t.Fatalf("expected arbitrary value to be invalid")
	}
}

