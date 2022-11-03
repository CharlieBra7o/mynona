package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/exp/slices"
	"rsc.io/quote"
)

// Config
type Config struct {
	Credentials *Credentials `json:"credentials"`
	Domains     []string     `json:"domains"`
}

var config Config
var migadu *MigaduAPI

type Credentials struct {
	User   string `json:"user"`
	APIKey string `json:"apiKey"`
}

// Error handling
func main() {
	// Error handling
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}

// Logic
func run() error {
	fmt.Println(quote.Go())

	// Register API handlers
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/api/domains", domainsHandler)

	// Read config
	// var config Config
	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	json.Unmarshal(configBytes, &config)

	// Create new API instance
	migadu = NewMigaduAPI(config.Credentials.User, config.Credentials.APIKey)

	migadu.POST("/domains/prestige-worldwi.de/mailboxes/demo/identities")

	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

	return nil
}

// Migadu API
type MigaduAPI struct {
	user       string
	apiKey     string
	httpClient http.Client
}

func NewMigaduAPI(user string, apiKey string) *MigaduAPI {
	return &MigaduAPI{user, apiKey, http.Client{}}
}
func (api MigaduAPI) GET(path string) string {
	url := MigaduAPIEndpoint + path

	request, _ := http.NewRequest("GET", url, nil)
	request.SetBasicAuth(config.Credentials.User, config.Credentials.APIKey)

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	responseBodyBytes, _ := ioutil.ReadAll(response.Body)
	responseBodyString := string(responseBodyBytes)

	// var addressAliases = new(AddressAliases)
	// if err := json.Unmarshal(responseBodyBytes, &addressAliases); err != nil {
	// 	log.Fatal(err)
	// 	return ""
	// }

	return responseBodyString
}
func (api MigaduAPI) POST(path string) string {
	fmt.Println("CALL migadu.POST()")
	url := MigaduAPIEndpoint + path

	type Identity struct {
		Name       string `json:"name"`
		Local_part string `json:"local_part"`
		Password   string `json:"password"`
	}

	identity := Identity{Name: "foobar.com", Local_part: "foobar", Password: "superpass"}
	fmt.Println(identity)
	identityJson, err := json.Marshal(identity)
	fmt.Println(identityJson)
	if err != nil {
		log.Fatal(err)
	}

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(identityJson))
	request.SetBasicAuth(config.Credentials.User, config.Credentials.APIKey)
	request.Header.Set("Content-Type", "application/json")

	fmt.Println(request)

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	responseBodyBytes, _ := ioutil.ReadAll(response.Body)
	responseBodyString := string(responseBodyBytes)

	// var addressAliases = new(AddressAliases)
	// if err := json.Unmarshal(responseBodyBytes, &addressAliases); err != nil {
	// 	log.Fatal(err)
	// 	return ""
	// }

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return responseBodyString
}

// REST API Handlers
func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested path:", r.URL.Path)

	path := strings.Split(r.URL.Path, "/")[2:]
	allowedPaths := []string{"mailboxes", "identities", "aliases", "rewrites"}

	switch r.Method {
	case "GET":
		fmt.Println("GET")
		if 3 <= len(path) && len(path) <= 4 && slices.Contains(allowedPaths, path[2]) {
			fmt.Println("GET index aliases for", path[1])
			if slices.Contains(config.Domains, path[1]) {
				// Passthrough INDEX request
				response := migadu.GET(strings.Join(path, "/"))
				fmt.Println((response))
				fmt.Fprintln(w, response)
			} else {
				http.Error(w, "Invalid Domain", http.StatusNotFound)
			}
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	case "POST":
		fmt.Println("POST")
	default:
		http.Error(w, "Method not supported", http.StatusBadRequest)
	}
}

func domainsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(config.Domains)
}
