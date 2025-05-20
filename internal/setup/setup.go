package setup

import (
	"os"
	"path/filepath"
)

// setup Base config (apiVersion, default template, etc)
// setup zcp config file (like default provider, etc)
// initial setup should land the base config in ~/.zcp/base_template
// zcp config file should also go to ~/.zcp/config if user opts to create it

type TUIConfig struct {
	Height         int
	Width          int
	templateConfig string
	templatePath   string
	storagePath    string
	done           bool
	KeyPrompt      string
}

func New() TUIConfig {
	homeDir, _ := os.UserHomeDir()
	templatePath := filepath.Join(homeDir, ".zcp", "templates")
	storagePath := filepath.Join(homeDir, ".zcp", "install-configs")

	keyPrompt := "Press Enter to continue..."

	return TUIConfig{
		Height:       24,
		Width:        80,
		templatePath: templatePath,
		storagePath:  storagePath,
		KeyPrompt:    keyPrompt,
	}
}
