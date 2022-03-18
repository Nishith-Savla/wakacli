package cmd

import (
	"github.com/Nishith-Savla/wakacli/api"
	"github.com/Nishith-Savla/wakacli/cmd/utils"
	"github.com/spf13/cobra"
)

var yesterdayCmd = &cobra.Command{
	Use:   "yesterday",
	Short: "View total duration for yesterday",
	Long:  `View the total duration for yesterday`,
	Run:   utils.PrintDuration(api.GetYesterday),
}

func init() {
	rootCmd.AddCommand(yesterdayCmd)
	yesterdayCmd.Flags().BoolP("sec", "s", false, "Include seconds in output")
}
