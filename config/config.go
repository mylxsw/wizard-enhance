package config

import (
	"encoding/json"

	"github.com/mylxsw/glacier/infra"
)

type Config struct {
	Listen          string `json:"listen"`
	Debug           bool   `json:"debug"`
	LogPath         string `json:"log_path"`
	APISecret       string `json:"-"`
	Version         string `json:"version"`
	GitCommit       string `json:"git_commit"`
	DBConnStr       string `json:"-"`
	GotenbergServer string `json:"gotenberg_server"`
	StoragePath     string `json:"storage_path"`
}

func (conf *Config) Serialize() string {
	rs, _ := json.Marshal(conf)
	return string(rs)
}

// Get return config object from container
func Get(cc infra.Resolver) *Config {
	return cc.MustGet(&Config{}).(*Config)
}
