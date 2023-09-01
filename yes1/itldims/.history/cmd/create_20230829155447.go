package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// Your etcd host address and port
const (
	etcdHost = "localhost"
	etcdPort = "2379"
)

func postSpecificKeyAnoop(w http.ResponseWriter, r *http.Request) {
	// Existing function code
	// ...
}

func putCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "put",
		Short: "Enter data into the etcd database",
		Run: func(cmd *cobra.Command, args []string) {
			// Your code to call postSpecificKeyAnoop function here
			httpReq, err := http.NewRequest("GET", "http://example.com", nil) // Create a dummy request, replace with actual data
			if err != nil {
				log.Fatal(err)
			}
			postSpecificKeyAnoop(nil, httpReq)
		},
	}
}
