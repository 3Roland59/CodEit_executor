package executor

import (
	"github.com/3roland59/CodEdit_executor/models"
	"github.com/3roland59/CodEdit_executor/runner"
	"github.com/3roland59/CodEdit_executor/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Execute(code string, language string, testCases []models.TestCase) models.ExecutionResponse {
	langConfig, err := GetLangConfig(language)
	if err != nil {
		return models.ExecutionResponse{
			Success: false,
			Message: err.Error(),
			Score:   0,
			Error:   err.Error(),
		}
	}

	codeFile := filepath.Join("/tmp", "Main."+langConfig.Extension)
	if err := os.WriteFile(codeFile, []byte(code), 0644); err != nil {
		return models.ExecutionResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to write code file: %v", err),
			Score:   0,
			Error:   err.Error(),
		}
	}

	passed := 0
	var results []models.TestCaseResult
	var combinedOutput []string

	for _, tc := range testCases {
		input, inputErr := utils.CastValue(tc.InputValue, tc.InputType)
		expected, expectedErr := utils.CastValue(tc.OutputValue, tc.OutputType)

		if inputErr != nil || expectedErr != nil {
			return models.ExecutionResponse{
				Success: false,
				Message: "Failed to cast input/output values",
				Score:   0,
				Error:   "Invalid input or output type",
			}
		}

		start := time.Now()
		output, err := runner.RunDocker(langConfig.Image, langConfig.Command, code, input)
		duration := time.Since(start)

		correct := err == nil && utils.IsOutputCorrect(expected, output)

		if correct {
			passed++
		}

		results = append(results, models.TestCaseResult{
			Input:         fmt.Sprintf("%v", input),
			Expected:      fmt.Sprintf("%v", expected),
			Actual:        strings.TrimSpace(output),
			Passed:        correct,
			ErrorMessage:  errorMessage(err),
			ExecutionTime: duration.String(),
		})

		combinedOutput = append(combinedOutput, strings.TrimSpace(output))
	}

	score := float64(passed) / float64(len(testCases)) * 100
	success := passed == len(testCases)

	return models.ExecutionResponse{
		Success:         success,
		Message:         "Execution completed",
		Score:           score,
		TestCaseResults: results,
		Output:          strings.Join(combinedOutput, "\n"),
	}
}

func errorMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

