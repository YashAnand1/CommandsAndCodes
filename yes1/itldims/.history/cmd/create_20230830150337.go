package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const etcdHost = "your_etcd_host"
const csvFile = "path_to_your_csv_file"

func uploadToEtcd() {
	// ... (your existing uploadToEtcd function code)
}

func postSpecificKeyAnoop(w http.ResponseWriter, r *http.Request) {
	// ... (your existing postSpecificKeyAnoop function code)
}

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create and upload data to etcd",
		Run: func(cmd *cobra.Command, args []string) {
			uploadToEtcd()
		},
	}

	rootCmd.AddCommand(createCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
