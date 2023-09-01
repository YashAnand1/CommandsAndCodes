package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var (
	itldims = &cobra.Command{
		Use:   "itldims",
		Short: "Interact with the etcd API",
		Long:  "A command-line tool to interact with the etcd API and tell if the connection has been made",
		Run: func(cmd *cobra.Command, args []string) {
			response, err := http.Get("http://localhost:8181/servers/")
			if err != nil {
				log.Fatalf("Failed to connect to the etcd API.")
			}
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {
				fmt.Println("Successfully connected with API. Interaction with etcd can be done.")
			}
		},
	}

	get = &cobra.Command{
		Use:   "get",
		Short: "Get keys with specific inputs from etcd API",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			data, err := fetchDataFromEtcdAPI()
			if err != nil {
				log.Fatalf("Failed to fetch data from the etcd API: %v", err)
			}

			if len(args) == 1 {
				args = append(args, "servers")
			}

			for key, value := range data {
				if strings.Contains(key, "{") || strings.Contains(key, "}") ||
					strings.Contains(value, "{") || strings.Contains(value, "}") {
					continue
				}

				if !strings.Contains(key, "data") &&
					(strings.Contains(key, args[0]) || strings.Contains(value, args[0])) &&
					(strings.Contains(key, args[1]) || strings.Contains(value, args[1])) {
					fmt.Printf("key=%s\nvalue=%s\n\n", key, value)
				}
			}
		},
	}
)

func fetchDataFromEtcdAPI() (map[string]string, error) {
	response, err := http.Get("http://localhost:8181/servers/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the etcd API: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data from the etcd API. Status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return parseKeyValuePairs(string(body)), nil
}

func parseKeyValuePairs(data string) map[string]string {
	result := make(map[string]string)
	lines := strings.Split(data, "\n")

	for i := 0; i < len(lines); i += 2 {
		key := strings.TrimSpace(lines[i])
		if i+1 >= len(lines) {
			break
		}

		value := strings.TrimSpace(lines[i+1])
		if value == "" {
			// If the value is empty, it means we have multi-line value.
			// Concatenate all consecutive non-empty lines as the value.
			var builder strings.Builder
			for j := i + 2; j < len(lines); j++ {
				line := strings.TrimSpace(lines[j])
				if line == "" {
					break
				}
				builder.WriteString(line + "\n")
			}
			value = strings.TrimSpace(builder.String())
		}

		result[key] = value
	}

	return result
}

func init() {
	itldims.AddCommand(get)
}

func main() {
	if err := itldims.Execute(); err != nil {
		log.Fatal(err)
	}
}
