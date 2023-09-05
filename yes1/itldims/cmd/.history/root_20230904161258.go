package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	itldims = &cobra.Command{
		Use:   "itldims",
		Short: "Interact with the etcd API",
		Long:  "A command-line tool to interact with the etcd API and check connection",
		Run: func(cmd *cobra.Command, args []string) { // in slice of string args, user input stored, cmd holds info on how things to be executed
			response, err := http.Get("http://localhost:8181/servers/") //we always get back 2: err and response
			if err != nil {
				fmt.Printf("\n\n ------: Failed to connect to the etcd API at http://localhost:8181/servers/:-------\n\n\n\n\n")
			}
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK { // could be replaced with status 200
				fmt.Println("Connected to API. Interaction with etcd can be done.")
			}
		},
	}
)

func init() {
	itldims.AddCommand(create) //acts like a function call to all the other commands in the directory
	itldims.AddCommand(get)
	itldims.AddCommand(delete)
	itldims.AddCommand(update)
}

func Execute() {
	if err := itldims.Execute(); err != nil {
		log.Fatal(err)
	}
}
