// package main

// import (
// 	"context"
// 	"encoding/csv"
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/tealeg/xlsx"
// 	clientv3 "go.etcd.io/etcd/client/v3"
// )

// var (
// 	// File paths
// 	excelFile = "/home/user/sk/etcd-inventory/etcd.xlsx"
// 	csvFile   = "/home/user/sk/etcd-inventory/myetcd.csv"
// 	etcdHost  = "localhost:2379"
// )

// type ServerData map[string]string

// func convertExcelToCSV(excelFile, csvFile string) {
// 	// Open the Excel file
// 	xlFile, err := xlsx.OpenFile(excelFile)
// 	if err != nil {
// 		log.Fatalf("Failed to open Excel file: %v", err)
// 	}

// 	// Create the CSV file
// 	file, err := os.Create(csvFile)
// 	if err != nil {
// 		log.Fatalf("Failed to create CSV file: %v", err)
// 	}
// 	defer file.Close()

// 	// Write data to the CSV file
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	// Iterate over sheets and rows in the Excel file
// 	for _, sheet := range xlFile.Sheets {
// 		for _, row := range sheet.Rows {
// 			var rowData []string
// 			for _, cell := range row.Cells {
// 				text := cell.String()
// 				rowData = append(rowData, text)
// 			}

// 			// Check if the row is empty
// 			isEmptyRow := true
// 			for _, field := range rowData {
// 				if field != "" {
// 					isEmptyRow = false
// 					break
// 				}
// 			}

// 			// Skip empty rows
// 			if !isEmptyRow {
// 				writer.Write(rowData)
// 			}
// 		}
// 	}
// }

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
// 	for _, record := range records[1:] {
// 		serverIP := record[0]
// 		serverType := record[1]
// 		serverData := make(ServerData)

// 		// Create server data dictionary
// 		for i := 2; i < len(headers); i++ {
// 			header := headers[i]
// 			value := record[i]
// 			//fmt.Println(value)
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

// 		// Check if creation date and time data is present
// 		var data map[string]interface{}
// 		err = json.Unmarshal(etcdValueData, &data)
// 		if err != nil {
// 			log.Printf("Failed to unmarshal JSON data: %v", err)
// 			continue
// 		}
// 		creationDateTime, ok := data["creation_date_time"].(string)
// 		if ok {
// 			fmt.Printf("Creation Date and Time: %s\n", creationDateTime)
// 		} else {
// 			fmt.Println("Creation Date and Time not found in JSON")
// 		}

// 		_, err = etcdClient.Put(context.Background(), etcdKeyData, string(etcdValueData))
// 		if err != nil {
// 			log.Printf("Failed to upload server data to etcd: %v", err)
// 		}
// 	}

// 	log.Println("Server details added to etcd successfully.")
// }

// func getServerData(w http.ResponseWriter, r *http.Request) {
// 	// Extract the server type and IP from the URL path
// 	log.Printf("response %v", r.URL.Path)

// 	// Connect to etcd
// 	ctx := context.TODO()
// 	etcdClient, err := clientv3.New(clientv3.Config{
// 		Endpoints: []string{etcdHost},
// 	})
// 	if err != nil {
// 		log.Printf("Failed to connect to etcd: %v", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer etcdClient.Close()

// 	// Construct the etcd key for the server data
// 	etcdKeyData := fmt.Sprintf(r.URL.Path)

// 	var revisions int
// 	response, err := etcdClient.Get(ctx, etcdKeyData, clientv3.WithSort(clientv3.SortByCreateRevision, clientv3.SortAscend))
// 	log.Printf("response %v", response)
// 	for _, kv := range response.Kvs {
// 		revisions = int(kv.ModRevision)
// 		log.Printf("revisions %v", revisions)
// 	}
// 	log.Printf("revision response: %v", revisions)

// 	var revisionslist []int64
// 	for i := revisions; i >= revisions-20; i-- {
// 		revisionslist = append(revisionslist, int64(i))
// 	}

// 	values, createRevisions, err := getRevisionValues(etcdClient, etcdKeyData, revisionslist)
// 	w.Header().Set("Content-Type", "text/plain")

// 	for i, value := range values {
// 		createRevision := createRevisions[i]
// 		createTime := time.Unix(createRevision/1000, 0) // Convert the create revision to a Unix timestamp
// 		fmt.Fprintf(w, "Value: %s, Creation Time: %s\n", value, createTime)
// 	}
// }

// func getRevisionValues(client *clientv3.Client, key string, revisions []int64) ([]string, []int64, error) {
// 	ctx := context.TODO()

// 	var values []string
// 	var createRevisions []int64

// 	for _, rev := range revisions {
// 		response, err := client.Get(ctx, key, clientv3.WithRev(rev))
// 		if err != nil {
// 			return nil, nil, err
// 		}

// 		if len(response.Kvs) > 0 {
// 			value := string(response.Kvs[0].Value)
// 			values = append(values, value)
// 			createRevisions = append(createRevisions, response.Kvs[0].CreateRevision)
// 		} else {
// 			fmt.Println("Value not found")
// 		}
// 	}

// 	return values, createRevisions, nil
// }

// func main() {
// 	// Convert Excel to CSV
// 	convertExcelToCSV(excelFile, csvFile)
// 	log.Println("Excel file converted to CSV successfully.")

// 	// Parse command-line flags
// 	flag.Parse()

// 	// Upload CSV data to etcd
// 	uploadToEtcd()

// 	// Start API server
// 	log.Println("Starting API server...")
// 	http.HandleFunc("/servers/", getServerData)
// 	err := http.ListenAndServe(":8181", nil)
// 	if err != nil {
// 		log.Fatalf("Failed to start API server: %v", err)
// 	}

// ;