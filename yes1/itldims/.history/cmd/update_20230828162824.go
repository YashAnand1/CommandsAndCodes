package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var update = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  `For updating a key or resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}
