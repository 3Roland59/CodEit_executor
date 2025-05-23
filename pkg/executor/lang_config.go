package executor

import (
	"encoding/json"
	"os"
	"sync"
	"fmt"
)

type LangConfig struct {
	Image     string `json:"image"`
	Extension string `json:"extension"`
	Command   string `json:"command"`
}

var (
	langConfigs map[string]LangConfig
	once        sync.Once
)

func LoadLangConfigs() error {
	var err error
	once.Do(func() {
		data, e := os.ReadFile("configs/languages.json")
		if e != nil {
			err = e
			return
		}
		err = json.Unmarshal(data, &langConfigs)
		if err != nil {
			return
		}
	})
	return err
}

// GetLangConfig returns config for the given language string
func GetLangConfig(lang string) (*LangConfig, error) {
	if langConfigs == nil {
		if err := LoadLangConfigs(); err != nil {
			return nil, fmt.Errorf("failed to load language configs: %w", err)
		}
	}
	config, ok := langConfigs[lang]
	if !ok {
		return nil, fmt.Errorf("language config not found for: %s", lang)
	}
	return &config, nil
}

