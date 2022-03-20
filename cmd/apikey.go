package cmd

import (
	"fmt"
	"os"

	"github.com/Nishith-Savla/wakacli/api"
	"github.com/spf13/cobra"
)

// apikeyCmd represents the apikey command
var apikeyCmd = &cobra.Command{
	Use:   "apikey",
	Short: "Add WakaTime API Key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			_, err := fmt.Fprintln(os.Stderr, "Please provide an API Key")
			if err != nil {
				panic(err)
			}
			return
		}

		if err := api.SetAPIKey(args[0]); err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("API Key added successfully")
	},
}

func init() {
	rootCmd.AddCommand(apikeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apikeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apikeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
