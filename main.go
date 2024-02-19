package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "strings"
    "os"
)

// SheetData represents the structure of the data to be served as JSON
type SheetData struct {
    Headers []string   `json:"headers"`
    Values  [][]string `json:"values"`
}

// fetchSheetData fetches and parses the Google Sheet CSV data
func fetchSheetData(sheetURL string) (SheetData, error) {
    resp, err := http.Get(sheetURL)
    if err != nil {
        return SheetData{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return SheetData{}, fmt.Errorf("received non-200 response code")
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return SheetData{}, err
    }

    r := csv.NewReader(strings.NewReader(string(body)))
    records, err := r.ReadAll()
    if err != nil {
        return SheetData{}, err
    }

    // Assuming the first row contains headers
    headers := records[0]
    values := records[1:]

    return SheetData{Headers: headers, Values: values}, nil
}

// handleRequest is the HTTP handler function
func handleRequest(w http.ResponseWriter, r *http.Request) {
    sheetURL := os.Getenv("GSHEET_PUBLISHED_URL")
    if sheetURL == "" {
        http.Error(w, "GSHEET_PUBLISHED_URL is not set", http.StatusInternalServerError)
        return
    }
    data, err := fetchSheetData(sheetURL)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/", handleRequest)
    fmt.Println("Server started on port 5678")
    log.Fatal(http.ListenAndServe(":5678", nil))
}

