package cmd

import (
	"fmt"
	"os"

	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd.Execute()
}

var flags struct {
	bump    string
	version string
}

var rootCmd = &cobra.Command{
	Use: "semver",
	Run: func(cmd *cobra.Command, args []string) {
		v, err := semver.NewVersion(flags.version)
		dieOnError(err, "failed to validate version "+flags.version)
		s := *v
		bumps := map[string]func() semver.Version{
			"patch": s.IncPatch,
			"minor": s.IncMinor,
			"major": s.IncMajor,
		}

		inc, ok := bumps[flags.bump]
		if !ok {
			dieOnError(fmt.Errorf("bump type %s does not exists", flags.bump), "")
		}
		s = inc()
		fmt.Printf("::set-output name=version::%s\n", s.String())
	},
}

func init() {
	rootCmd.Flags().StringVar(&flags.bump, "bump", "patch", "")
	rootCmd.Flags().StringVar(&flags.version, "version", "", "")
}

func dieOnError(err error, msg string) {
	if err == nil {
		return
	}

	if msg == "" {
		fmt.Fprintf(os.Stderr, "[ERROR]: %v\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "[ERROR]: %s: %v\n", msg, err)
	}

	os.Exit(1)
}
