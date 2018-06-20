package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version - Iris Version
const Version = "0.2.0"

// VersionCmd - The version of gaia
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s\n", Version)
	},
}
