package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create <key>",
	Short: "Create and upload data to etcd",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		uploadToEtcd()
		////  anoop new post function  ///////
		var key1 string
		var numberof_key int
		////////////  anoop ///////////
		if len(args) == 2 {
			key1 = args[0]
			numberof_key = len(strings.Split(key1, "/")) // number of key components
		}


			var url string = "http://localhost:8181/" + "servers/"

			var line string = args[0]
			etcdKey := args[0]
			etcdValue := args[1]
			line = "{" + "\"EtcdKey\"" + ":" + "\"" + etcdKey + "\","
			line = line + "\"EtcdValue\"" + ":" + "\"" + etcdValue + "\"}"

			// etcdKey := strings.Split(line, "/")[:4]
			// etcdValue := strings.Split(line, "/")[4] // Extract the last component of the key

			fmt.Println(line)
			fmt.Println(url)
			var jsonStr = []byte(line)
			responseBody := bytes.NewBuffer(jsonStr)
			resp, err := http.Post(url, "application/json", responseBody)
			if err != nil {
				fmt.Println(err)
			}
			if err == nil {
				defer resp.Body.Close()
			}

			if resp.StatusCode == 200 {
				// fmt.Printf("Key: %s is netered as %s succesfully", "servers/Physical/10.246.40.139/Hostname", "vahanapp00")
				fmt.Printf("Key: %s has been metered as %s succesfully\n", etcdKey, string(etcdValue))
			}
			/////
		}
	},
}
