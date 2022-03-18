package cmd

import (
	"github.com/Nishith-Savla/wakacli/api"
	"github.com/Nishith-Savla/wakacli/cmd/utils"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "View total duration for today",
	Long:  `View the total duration for today`,
	Run:   utils.PrintDuration(api.GetToday),
}

func init() {
	rootCmd.AddCommand(todayCmd)
	todayCmd.Flags().BoolP("sec", "s", false, "Include seconds in output")
}
