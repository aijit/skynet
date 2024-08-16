package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	skynetConfig = new(SkynetConfig)
)

type SkynetConfig struct {
	Thread     int
	Harbor     int
	Profile    int
	Daemon     string
	ModulePath string
	Bootstrap  string
	Logger     string
	Logservice string
}

func GetConfig() *SkynetConfig {
	return skynetConfig
}

func (cfg *SkynetConfig) Load() (err error) {
	env := NewSkynetEnv()
	defer env.Close()
	if err = env.Load(os.Args[1]); err != nil {
		return
	}
	cfg.Daemon = env.String("daemon")
	cfg.Harbor = env.Int("harbor")
	str, _ := json.MarshalIndent(cfg, "", " ")
	fmt.Println(string(str))
	return
}
