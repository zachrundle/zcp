package setup

import (
	"fmt"
	"os"
	"path/filepath"
)

// SetupDirectories creates the ~/.zcp/templates/ and ~/.zcp/install-configs directories
// TODO: look into better utilizing MkdirAll function so the code is less repetitive
func SetupDirectories() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v",err)
	}
	zcpDir := filepath.Join(homeDir, ".zcp")
	if err := os.MkdirAll(zcpDir, 0755); err != nil {
		return fmt.Errorf("failed to create zcp directory: %v", err)
	}
	templatesDir := filepath.Join(zcpDir, "templates")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		return fmt.Errorf("failed to create templates directory: %v", err)
	}
	installconfigDir := filepath.Join(zcpDir, "install-configs")
	if err := os.MkdirAll(installconfigDir, 0755); err != nil {
		return fmt.Errorf("failed to create install-configs direcotry: %v", err)
	}
	return nil
}

// BaseTemplateExists checks if the base.yaml file exists in the ~/.zcp/templates/ directory
func BaseTemplateExists() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("failed to get home directory: %v", err)
	}
	baseTemplateFilePath := filepath.Join(homeDir, ".zcp", "templates", "base.yaml")
    info, err := os.Stat(baseTemplateFilePath)
    if err != nil {
        return false
    }
    return info.IsDir()
}