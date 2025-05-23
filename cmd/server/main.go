package main

import (
    "log"
    "net/http"

    "github.com/3roland59/CodEdit_executor/internal/handler"
)

func main() {
    http.HandleFunc("/execute", handler.ExecuteCodeHandler)
    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
