package models

// TestCase defines the structure for a single test case
type TestCase struct {
	InputType   string `json:"inputType"`   // e.g., "Integer", "String", "Array"
	InputValue  string `json:"inputValue"`  // raw input value, to be cast
	OutputType  string `json:"outputType"`  // expected type of output
	OutputValue string `json:"outputValue"` // expected output value
}

// ExecutionRequest represents the request payload for code execution
type ExecutionRequest struct {
	Code      string     `json:"code"`      // The code to run
	Language  string     `json:"language"`  // Programming language (e.g., python, java)
	TestCases []TestCase `json:"testCases"` // List of test cases to run against the code
}

