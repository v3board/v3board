package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version   string
	GoVersion string
	GitCommit string
	Built     string
	OsArch    string

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show v3board version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%-20s\t%s\n%-20s\t%s\n%-20s\t%s\n%-20s\t%s\n%-20s\t%s\n", "Version:", Version, "Go version:", GoVersion, "Git commit:", GitCommit, "OS/Arch:", OsArch, "Built:", Built)
		},
	}
)
