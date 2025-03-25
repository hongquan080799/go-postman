package handlers

import (
    "encoding/json"
    "io"
    "net/http"
    "strings"

    "github.com/hongquan080799/go-postman/models"
)

// SendRequest handles API requests from the frontend
func SendRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Read request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Parse JSON request
    var reqData models.RequestData
    if err := json.Unmarshal(body, &reqData); err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }

    // Create new request
    req, err := http.NewRequest(reqData.Method, reqData.URL, strings.NewReader(reqData.Body))
    if err != nil {
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
        return
    }

    // Set headers
    for key, value := range reqData.Headers {
        req.Header.Set(key, value)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, "Request failed", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Read response body
    respBody, _ := io.ReadAll(resp.Body)

    // Prepare response data
    responseData := map[string]interface{}{
        "status":  resp.Status,
        "headers": resp.Header,
        "body":    string(respBody),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(responseData)
}

