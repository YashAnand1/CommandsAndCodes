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
	var url string = "http://localhost:8181/" + key
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

// func uploadToEtcd() {
// 	// Connect to etcd
// 	log.Println("Entered into function")
// 	etcdClient, err := clientv3.New(clientv3.Config{
// 		Endpoints: []string{etcdHost},
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to etcd: %v", err)
// 	}
// 	defer etcdClient.Close()

// 	// Read the CSV file
// 	file, err := os.Open(csvFile)
// 	log.Println("reading file")
// 	if err != nil {
// 		log.Fatalf("Failed to open CSV file: %v", err)
// 	}
// 	defer file.Close()

// 	// Parse the CSV file
// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		log.Fatalf("Failed to read CSV file: %v", err)
// 	}

// 	// Iterate over the records and upload to etcd
// 	headers := records[0]
// 	type ServerData map[string]string

// 	for _, record := range records[1:] {
// 		serverIP := record[0]
// 		serverType := record[1]
// 		serverData := make(ServerData)

// 		// Create server data dictionary
// 		for i := 2; i < len(headers); i++ {
// 			header := headers[i]
// 			value := record[i]
// 			serverData[header] = value
// 		}

// 		// Set key-value pairs in etcd for each data field
// 		for header, value := range serverData {
// 			etcdKey := fmt.Sprintf("/servers/%s/%s/%s", serverType, serverIP, header)
// 			etcdValue := value
// 			fmt.Println(etcdKey)
// 			fmt.Println(etcdValue)
// 			_, err := etcdClient.Put(context.Background(), etcdKey, etcdValue)
// 			if err != nil {
// 				log.Printf("Failed to upload key-value to etcd: %v", err)
// 			}
// 		}

// 		// Set key-value pair for server data
// 		etcdKeyData := fmt.Sprintf("/servers/%s/%s/data", serverType, serverIP)
// 		etcdValueData, err := json.Marshal(serverData)
// 		if err != nil {
// 			log.Printf("Failed to marshal server data: %v", err)
// 			continue
// 		}
// 		_, err = etcdClient.Put(context.Background(), etcdKeyData, string(etcdValueData))
// 		if err != nil {
// 			log.Printf("Failed to upload server data to etcd: %v", err)
// 		}
// 	}

// 	log.Println("Server details added to etcd successfully.")
// }
