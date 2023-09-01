package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var get = &cobra.Command{ // Cobra command is 'get'
	Use:   "get",
	Short: "Search Attributes & Values from etcd API",
	Long: `Data retrieval can be done using the following combinations:
 - 'itldims get <KeyComponent> <KeyComponent/Value>'
 - 'itldims get <KeyComponent/Value>'
`,
	Args: cobra.RangeArgs(1, 2), //number of allowed arguments=1&2
	Run: func(cmd *cobra.Command, args []string) {
		var key1 string
		var numberof_key int
		////////////  anoop ///////////
		if len(args) == 1 {
			key1 = args[0]
			numberof_key = len(strings.Split(key1, "/"))
		}

		//fmt.Printf("key1:%s", key1)

		if numberof_key == 4 {
			data1, err1 := fetchDataFromAPIWithKey(key1)
			if err1 != nil {
				fmt.Printf("Failed to fetch data from the etcd API: %v", err1) // Print an error message if error occured
			}

			fmt.Println(data1)
		}

		if numberof_key < 4 {
			//////////////////////////////
			data, err := fetchDataFromAPI()
			if err != nil {
				fmt.Printf("Failed to fetch data from the etcd API: %v", err) // Print an error message if error occured
			}

			for key, value := range parseKeyValuePairs(data) { // data is gone through and key value variables entered

				// For skipping/continuing data key completely
				if strings.Contains(key, "{") || strings.Contains(key, "}") || strings.Contains(key, "data") ||
					strings.Contains(value, "{") || strings.Contains(value, "}") {
					continue
				}

				ATs := make(map[string]string) // a MAP for attributes created for storing server attributes
				splitKey := strings.Split(key, "/")
				serverAtr := splitKey[4]
				serverIP := splitKey[3]
				ATs[serverAtr] = serverAtr // Store the server attribute in the map

				// IF SERVERIP IS MENTIONED
				if strings.Contains(args[0], ".") && (strings.Contains(key, args[0])) {
					if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
						continue
					}
					fmt.Printf("%s:\n%s\n\n", serverAtr, value)
				} else { // IF SERVERIP IS NOT MENTIONED
					if strings.Contains(key, args[0]) || strings.Contains(value, args[0]) {
						if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
							continue
						}
						fmt.Printf("Server IP:\n%s\n%s:%s\n\n", serverIP, serverAtr, value) // Print server IP, attribute, and value
					}
				}
			}

		} // end of else statement

		////  anoop new post function  ///////
		if numberof_key == 5 {

			var url string = "http://localhost:8181/" + "servers/anoop"
			////////////////////////

			//https://thedevelopercafe.com/articles/make-post-request-in-go-d9756284d70b
			// etcdKey  and etcdValue
			// etcdKey="/servers/Physical/10.246.40.139/Hostname"
			// etcdValue="vahanapp04"
			//http.Post()                    //Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
			//http.Get()                     //Get(url string) (resp *http.Response, err error)

			var line string

			line = "{" + "\"EtcdKey\"" + ":" + "\"/servers/Physical/10.246.40.139/Hostname\","
			line = line + "\"EtcdValue\"" + ":" + "\"vahanapp06\"}"

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
				fmt.Printf("Key: %s id netered as %s succesfully", "servers/Physical/10.246.40.139/Hostname", "vahanapp05")
			}
			/////
		}

		/////////////////////////

		if numberof_key == 6 {
			var url string = "http://localhost:8181/" + "servers/delete"
			////////////////////////

			var line string = "{" + "\"EtcdKey\"" + ":" + "\"/servers/Physical/10.246.40.139/Hostname\"" + "}"
			fmt.Println(line)
			fmt.Println(url)
			var jsonStr = []byte(line)
			responseBody := bytes.NewBuffer(jsonStr)

			////////////////////////////////////////
			// create a new HTTP client
			client := &http.Client{}

			// create a new DELETE request
			req, err := http.NewRequest("DELETE", url, responseBody)
			if err != nil {
				panic(err)
			}

			// send the request
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				fmt.Printf("Key: %s deleted succesfully", line)
			}
			// read the response body
			// body, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	panic(err)
			// }

			// print the response body
			// fmt.Println(string(body))
			///////////////////////////////////////

			/////
		}

	},
}

// Initialize 'get' and add subcommands
func init() {
	get.AddCommand(put)
	get.AddCommand(attributes)
	get.AddCommand(Types)
	get.AddCommand(servers)
}
