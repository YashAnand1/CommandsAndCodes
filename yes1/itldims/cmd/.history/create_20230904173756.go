package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "Create and upload data to etcd",
	// Args:  cobra.ExactArgs(2), //
	Run: func(cmd *cobra.Command, args []string) {
		var url string = "http://192.168.122.128:8181/" + "servers/Anoop"
		etcdKey := key
		etcdValue := value

		line := "{" + "\"EtcdKey\"" + ":" + "\"" + etcdKey + "\","     //added user arg 1
		line = line + "\"EtcdValue\"" + ":" + "\"" + etcdValue + "\"}" //added user arg 2

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
			fmt.Printf("Key: %s has been metered as %s succesfully\n", strings.ToUpper(etcdKey), strings.ToUpper(etcdValue))
		}
		/////
		// }

	},
}

var key string
var value string

func init() {
	itldims.AddCommand(create)

	create.PersistentFlags().StringVarP(&key, "key", "k", "", "A flag for the KEY")
	create.PersistentFlags().StringVarP(&value, "value", "v", "", "A flag for the VALUE")

}
