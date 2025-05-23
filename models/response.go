package models

// TestCaseResult stores the outcome of a single test case
type TestCaseResult struct {
	Input         string `json:"input"`
	Expected      string `json:"expected"`
	Actual        string `json:"actual"`
	Passed        bool   `json:"passed"`
	ErrorMessage  string `json:"errorMessage,omitempty"` // populated only if an error occurred
	ExecutionTime string `json:"executionTime,omitempty"` // optional, for benchmarking
}

// ExecutionResponse represents the response after running the code
type ExecutionResponse struct {
	Success         bool              `json:"success"`
	Message         string            `json:"message"`
	Score         float64            `json:"score"`
	TestCaseResults []TestCaseResult  `json:"testCaseResults,omitempty"` // one per test case
	Output          string            `json:"output,omitempty"`          // raw stdout output
	Error           string            `json:"error,omitempty"`           // if any execution error
}

