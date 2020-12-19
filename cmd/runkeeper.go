package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// runkeeperCmd represents the runkeeper command
var runkeeperCmd = &cobra.Command{
	Use:   "runkeeper",
	Short: "A simple Runkeeper client for uploading data",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		subConfig := viper.Sub("runkeeper")
		username := ""
		if subConfig.GetString("username") != "" {
			username = subConfig.GetString("username")
		}
		password := ""
		if subConfig.GetString("password") != "" {
			password = subConfig.GetString("password")
		}
		fmt.Println("Runkeeper config")
		fmt.Printf("username: %s \n", username)
		fmt.Printf("password: %s \n", password)
	},
}

func init() {
	rootCmd.AddCommand(runkeeperCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runkeeperCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runkeeperCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
