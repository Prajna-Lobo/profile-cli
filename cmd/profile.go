/*
Copyright © 2024 Prajna Lobo shineprajna@gmail.com
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// profileCmd represents the get command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Welcome to my profile",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("cat", "profile.md").Output()
		if err != nil {
			fmt.Printf("Cannot see profile %v", err)
		}
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	profileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	profileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
