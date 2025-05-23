package handler

import (
	"github.com/3roland59/CodEdit_executor/models"
	"github.com/3roland59/CodEdit_executor/pkg/executor"
	"encoding/json"
	"net/http"
)

// ExecuteCodeHandler handles the code execution endpoint
func ExecuteCodeHandler(w http.ResponseWriter, r *http.Request) {
	var req models.ExecutionRequest

	// Decode the JSON request body into ExecutionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Run the execution engine with code, language, and test cases
	res := executor.Execute(req.Code, req.Language, req.TestCases)

	// Respond with the structured ExecutionResponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

