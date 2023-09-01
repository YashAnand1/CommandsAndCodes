package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"go.etcd.io/etcd/clientv3"
)

var (
	etcdHost = "localhost:2379"
)

var create = &cobra.Command{
	Use:   "create <key>",
	Short: "Create and upload data to etcd",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		// Call the function to handle key creation with the provided key value
		postSpecificKeyAnoop(key)
	},
}

func postSpecificKeyAnoop(key string) {

	// Connect to etcd
	log.Println("Entered into function")
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer etcdClient.Close()

	// You don't need the following lines if you're not working with an HTTP request.
	// These lines are related to reading from an HTTP request, which is not applicable here.
	/*
		log.Printf("response %v", r.URL.Path)
		responseBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Response:", string(responseBody))
	*/

	type Keydata struct {
		EtcdKey   string
		EtcdValue string
	}
	keydata1 := Keydata{
		EtcdKey:   key,               // Use the provided key as the EtcdKey
		EtcdValue: "your-value-here", // Set your value here
	}

	_, err = etcdClient.Put(context.Background(), keydata1.EtcdKey, keydata1.EtcdValue)
	if err != nil {
		log.Printf("Failed to upload server data to etcd: %v", err)
	}
}
