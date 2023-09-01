package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var create = &cobra.Command{
	Use:  "create",
	Long: `For creating keys or resources - Allows adding values`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}
