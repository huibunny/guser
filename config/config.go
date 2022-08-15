package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App    `yaml:"app"`
		Log    `yaml:"logger"`
		Consul `yaml:"consul"`
		PG     `yaml:"postgres"`
		RMQ    `yaml:"rabbitmq"`
		Wx     `yaml:"wx"`
	}

	// App -.
	App struct {
		Name        string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version     string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		TokenExpire int64  `env-required:"true" yaml:"token_expire" env:"APP_TOKEN_EXPIRE"`
		Secret      string `env-required:"true" yaml:"secret" env:"APP_SECRET"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// Consul -.
	Consul struct {
		CheckApi string   `env-required:"true" yaml:"check_api"    env:"CONSUL_CHECKAPI"`
		Interval string   `env-required:"true" yaml:"interval"    env:"CONSUL_INTERVAL"`
		Timeout  string   `env-required:"true" yaml:"timeout"    env:"CONSUL_TIMEOUT"`
		Tags     []string `yaml:"tags"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true" yaml:"url"      env:"PG_URL"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env-required:"true" yaml:"rpc_server_exchange" env:"RMQ_RPC_SERVER"`
		ClientExchange string `env-required:"true" yaml:"rpc_client_exchange" env:"RMQ_RPC_CLIENT"`
		URL            string `env-required:"true" yaml:"url"                 env:"RMQ_URL"`
	}

	// WX -.
	Wx struct {
		AppID     string `env-required:"true" yaml:"appid" env:"APP_APPID"`
		AppSecret string `env-required:"true" yaml:"appsecret" env:"APP_APPSECRET"`
	}
)

// NewConfig returns app config.
func NewConfig(configFile string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(configFile, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
