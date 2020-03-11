package main

import (
	"fmt"
	"github.com/metamatex/metamatemono/metactl/cmd/metactl/v0"
	"github.com/metamatex/metamatemono/metactl/pkg/v0/types"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev-0.0.0"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "metactl",
	Short: "",
	Long:  "",
}

func init() {
	v := types.Version{
		Version: version,
		Commit: commit,
		Date: date,
	}

	v0.AddV0(rootCmd, false, v)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
