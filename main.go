package main

import (
    "net/http"
    "html/template"
    "github.com/hongquan080799/go-postman/handlers"
)


func main() {
    tmpl := template.Must(template.ParseFiles("templates/index.html"));
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, nil);
    });

    http.HandleFunc("/send-request", handlers.SendRequest)
    http.ListenAndServe(":8080", nil);
}
