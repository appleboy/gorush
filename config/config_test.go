package config

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// Test file is missing
func TestMissingFile(t *testing.T) {
	filename := "test"
	_, err := LoadConfYaml(filename)

	assert.NotNil(t, err)
}

// Test wrong json format
func TestWrongYAMLormat(t *testing.T) {
	content := []byte(`Wrong format`)

	filename := "tempfile"

	if err := ioutil.WriteFile(filename, content, 0644); err != nil {
		log.Fatalf("WriteFile %s: %v", filename, err)
	}

	// clean up
	defer os.Remove(filename)

	// parse JSON format error
	_, err := LoadConfYaml(filename)

	assert.NotNil(t, err)
}

// Test config file.
func TestReadConfig(t *testing.T) {
	config, err := LoadConfYaml("../config/config.yml")

	assert.Nil(t, err)
	assert.Equal(t, "8088", config.Core.Port)
	assert.True(t, config.Android.Enabled)
}

// Test config file.
func TestDefaultConfig(t *testing.T) {
	config := BuildDefaultPushConf()

	assert.Equal(t, "8088", config.Core.Port)
	assert.Equal(t, "/sys/stats", config.API.SysStatURI)
	assert.False(t, config.Android.Enabled)
}
