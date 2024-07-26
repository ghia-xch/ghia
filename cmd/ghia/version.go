package main

import (
	"fmt"
	"github.com/ghia-xch/ghia/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version",
	Long:  `Prints version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pkg.SemVer)
	},
}
