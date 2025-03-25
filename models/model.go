package models


type RequestData struct {
    URL     string `json:"url"`
    Method  string `json:"method"`
    Headers map[string]string `json:"headers"`
    Body    string `json:"body"`
}

