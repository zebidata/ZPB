package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// VersionCmd prints out the current sdk version
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the app version",
		Run:   printVersion,
	}
)

// GetVersion return version of CLI/node and commit hash
func GetVersion() string {
	v := Version
	if GitCommit != "" {
		v = v + "-" + GitCommit
	}
	return v
}

// CMD
func printVersion(_ *cobra.Command, _ []string) {
	v := GetVersion()
	fmt.Println(v)
}