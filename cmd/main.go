package cmd

import (
    "fmt"
    "os"
    // tea "github.com/charmbracelet/bubbletea"
	"github.com/zachrundle/zcp/internal/version"
	// "github.com/spf13/cobra"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("zcp version %s\n", version.Get())
		os.Exit(0)
	}
}
