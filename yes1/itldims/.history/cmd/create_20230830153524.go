package cmd

import (
	"github.com/spf13/cobra"
)

var (
	etcdHost = "localhost:2379"
)

var createCmd = &cobra.Command{
	Use:   "create <key>",
	Short: "Create and upload data to etcd",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		// Call the function to handle key creation with the provided key value
		postSpecificKeyAnoop(key)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func postSpecificKeyAnoop(key string) {
	// Your function implementation here
}
