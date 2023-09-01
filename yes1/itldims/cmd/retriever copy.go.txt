package cmd

import (
	"fmt"
	"io" // i/o operations
	"net/http"
	"strings"
)

func fetchDataFromAPI() (string, error) { //returns string and error
	fmt.Println("in fetchDataFromAPI")
	//////////////////////////////
	//fmt.Println(ar)

	////////////////////////

	response, err := http.Get("http://localhost:8181/servers/") //Get request sent to the API URL for fetching data
	if err != nil {
		fmt.Printf("%s", err)
	}

	data, err := io.ReadAll(response.Body) //Reads response body using the io.ReadAll
	if err != nil {
		fmt.Printf("%s", err)
	}

	return string(data), nil //returns the fetched data as a string
}

// //////////////  new function by anoop ////////////
func fetchDataFromAPIWithKey(key string) (string, error) { //returns string and error
	fmt.Println("in fetchDataFromAPI")
	//////////////////////////////
	var url string
	url = "http://localhost:8181/" + key
	////////////////////////

	//https://www.soberkoder.com/consume-rest-api-go/
	//https://thedevelopercafe.com/articles/make-post-request-in-go-d9756284d70b
	// etcdKey  and etcdValue
	// etcdKey="/servers/Physical/10.246.40.139/Hostname"
	// etcdValue="vahanapp04"
	//http.Post()                    //Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
	//http.Get()                     //Get(url string) (resp *http.Response, err error)
	response, err := http.Get(url) //Get request sent to the API URL for fetching data
	if err != nil {
		fmt.Printf("%s", err)
	}

	data, err := io.ReadAll(response.Body) //Reads response body using the io.ReadAll
	if err != nil {
		fmt.Printf("%s", err)
	}

	return string(data), nil //returns the fetched data as a string
}

// /////////////////////////////////////
func parseKeyValuePairs(data string) map[string]string { //string as input and returns a map of strings.
	result := make(map[string]string) //KeyValue pairse to be stored here

	keyValuePairs := strings.Split(data, "Key:")

	for _, kv := range keyValuePairs { //Each keyvalue is gone through the keyValuePairs

		lines := strings.Split(kv, "Value:") //data split into keyvaluepairs is split into kv, both keyvalue are stored here
		if len(lines) == 2 {                 //if split created 2 lines then key value were split successfuly, key = lines[0], value = lines[1]

			key := strings.TrimSpace(lines[0])
			value := strings.TrimSpace(lines[1])
			result[key] = value // result map is set as key = key, value = value
		}
	}

	return result
}
