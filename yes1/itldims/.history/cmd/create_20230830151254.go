package cmd

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"go.etcd.io/etcd/clientv3"
)

var (
	// File paths
	excelFile = "/home/user/yes1/etcd-inventory/etcd.xlsx"
	csvFile   = "/home/user/yes1/etcd-inventory/myetcd.csv"
	etcdHost  = "localhost:2379"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create and upload data to etcd",
	Run: func(cmd *cobra.Command, args []string) {
		uploadToEtcd()
	},
}

func uploadToEtcd() {
	// Connect to etcd
	log.Println("Entered into function")
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	// Read the CSV file
	file, err := os.Open(csvFile)
	log.Println("reading file")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Iterate over the records and upload to etcd
	headers := records[0]
	for _, record := range records[1:] {
		serverIP := record[0]
		serverType := record[1]
		serverData := make(ServerData)

		// Create server data dictionary
		for i := 2; i < len(headers); i++ {
			header := headers[i]
			value := record[i]
			serverData[header] = value
		}

		// Set key-value pairs in etcd for each data field
		for header, value := range serverData {
			etcdKey := fmt.Sprintf("/servers/%s/%s/%s", serverType, serverIP, header)
			etcdValue := value
			fmt.Println(etcdKey)
			fmt.Println(etcdValue)
			_, err := etcdClient.Put(context.Background(), etcdKey, etcdValue)
			if err != nil {
				log.Printf("Failed to upload key-value to etcd: %v", err)
			}
		}

		// Set key-value pair for server data
		etcdKeyData := fmt.Sprintf("/servers/%s/%s/data", serverType, serverIP)
		etcdValueData, err := json.Marshal(serverData)
		if err != nil {
			log.Printf("Failed to marshal server data: %v", err)
			continue
		}
		_, err = etcdClient.Put(context.Background(), etcdKeyData, string(etcdValueData))
		if err != nil {
			log.Printf("Failed to upload server data to etcd: %v", err)
		}
	}

	log.Println("Server details added to etcd successfully.")
}

func postSpecificKeyAnoop(w http.ResponseWriter, r *http.Request) {
	// Extract the server type and IP from the URL path
	log.Printf("response %v", r.URL.Path)
	// Read response body
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response:", string(responseBody))

	type Keydata struct {
		EtcdKey   string
		EtcdValue string
	}
	var keydata1 Keydata

	err1 := json.Unmarshal(responseBody, &keydata1)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("Struct is:", keydata1)
	fmt.Printf("key:%s | value:%s\n", keydata1.EtcdKey, keydata1.EtcdValue)

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()
	_, err = etcdClient.Put(context.Background(), keydata1.EtcdKey, keydata1.EtcdValue)
	if err != nil {
		log.Printf("Failed to upload server data to etcd: %v", err)
	}

}
