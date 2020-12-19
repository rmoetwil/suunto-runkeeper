package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// suuntoCmd represents the suunto command
var suuntoCmd = &cobra.Command{
	Use:   "suunto",
	Short: "A simple Suunto client for downloading data",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		subConfig := viper.Sub("suunto")
		username := ""
		if subConfig.GetString("username") != "" {
			username = subConfig.GetString("username")
		}
		password := ""
		if subConfig.GetString("password") != "" {
			password = subConfig.GetString("password")
		}
		fmt.Println("Suunto config")
		fmt.Printf("username: %s \n", username)
		fmt.Printf("password: %s \n", password)
	},
}

func init() {
	rootCmd.AddCommand(suuntoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// suuntoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// suuntoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
