package app

import (
	"fmt"

	"github.com/appleboy/gorush/config"
)

// MergeConfig merges CLI options into the configuration.
// CLI options take precedence over config file values.
func MergeConfig(cfg *config.ConfYaml, opts *Options) error {
	// iOS options
	if opts.Conf.Ios.KeyPath != "" {
		cfg.Ios.KeyPath = opts.Conf.Ios.KeyPath
	}
	if opts.Conf.Ios.KeyID != "" {
		cfg.Ios.KeyID = opts.Conf.Ios.KeyID
	}
	if opts.Conf.Ios.TeamID != "" {
		cfg.Ios.TeamID = opts.Conf.Ios.TeamID
	}
	if opts.Conf.Ios.Password != "" {
		cfg.Ios.Password = opts.Conf.Ios.Password
	}
	if opts.Conf.Ios.Production {
		cfg.Ios.Production = opts.Conf.Ios.Production
	}

	// Android options
	if opts.Conf.Android.KeyPath != "" {
		cfg.Android.KeyPath = opts.Conf.Android.KeyPath
	}

	// Huawei options
	if opts.Conf.Huawei.AppSecret != "" {
		cfg.Huawei.AppSecret = opts.Conf.Huawei.AppSecret
	}
	if opts.Conf.Huawei.AppID != "" {
		cfg.Huawei.AppID = opts.Conf.Huawei.AppID
	}

	// Storage options
	if opts.Conf.Stat.Engine != "" {
		cfg.Stat.Engine = opts.Conf.Stat.Engine
	}
	if opts.Conf.Stat.Redis.Addr != "" {
		cfg.Stat.Redis.Addr = opts.Conf.Stat.Redis.Addr
	}

	// Server options with validation
	if opts.Conf.Core.Port != "" {
		if err := config.ValidatePort(opts.Conf.Core.Port); err != nil {
			return fmt.Errorf("invalid port from command line: %w", err)
		}
		cfg.Core.Port = opts.Conf.Core.Port
	}
	if opts.Conf.Core.Address != "" {
		if err := config.ValidateAddress(opts.Conf.Core.Address); err != nil {
			return fmt.Errorf("invalid address from command line: %w", err)
		}
		cfg.Core.Address = opts.Conf.Core.Address
	}
	if opts.Conf.Core.HTTPProxy != "" {
		cfg.Core.HTTPProxy = opts.Conf.Core.HTTPProxy
	}

	// PID options
	if opts.Conf.Core.PID.Path != "" {
		if err := config.ValidatePIDPath(opts.Conf.Core.PID.Path); err != nil {
			return fmt.Errorf("invalid PID path from command line: %w", err)
		}
		cfg.Core.PID.Path = opts.Conf.Core.PID.Path
		cfg.Core.PID.Enabled = true
		cfg.Core.PID.Override = true
	}

	return nil
}

// ValidateAndMerge loads config, merges CLI options, and validates.
func ValidateAndMerge(opts *Options) (*config.ConfYaml, error) {
	cfg, err := config.LoadConf(opts.ConfigFile)
	if err != nil {
		return nil, fmt.Errorf("load yaml config file error: %w", err)
	}

	if err := MergeConfig(cfg, opts); err != nil {
		return nil, err
	}

	if err := config.ValidateConfig(cfg); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}
