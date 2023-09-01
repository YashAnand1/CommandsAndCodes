package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"go.etcd.io/etcd/clientv3"
)

// Your etcd host address and port
const (
	etcdHost = "localhost"
	etcdPort = "2379"
)

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
		Endpoints: []string{etcdHost + ":" + etcdPort},
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

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(putCommand())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
