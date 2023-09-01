package cmd

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "Create and upload data to etcd",
	Run: func(cmd *cobra.Command, args []string) {
		uploadToEtcd()

		etcdKey := keyFlag
		etcdValue := valueFlag

		var url string = "http://localhost:8181/" + "servers/"
		line := "{" + "\"EtcdKey\"" + ":" + "\"" + etcdKey + "\","
		line += "\"EtcdValue\"" + ":" + "\"" + etcdValue + "\"}"

		fmt.Println(line)
		fmt.Println(url)
		jsonStr := []byte(line)
		responseBody := bytes.NewBuffer(jsonStr)
		resp, err := http.Post(url, "application/json", responseBody)
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			defer resp.Body.Close()
		}

		if resp.StatusCode == 200 {
			fmt.Printf("Key: %s has been metered as %s successfully\n", etcdKey, etcdValue)
		}
	},
}

var keyFlag string
var valueFlag string

func init() {
	itldims.AddCommand(create)

	create.Flags().StringVarP(&keyFlag, "key", "", "", "Actual KEY")
	create.Flags().StringVarP(&valueFlag, "value", "", "", "Actual VALUE")

	// create.MarkFlagRequired("key")
	// create.MarkFlagRequired("value")
}

// ... (rest of your code)
