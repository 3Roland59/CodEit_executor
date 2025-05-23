package utils

import (
    "errors"
    "reflect"
    "strconv"
    "strings"
)

// CastValue parses a string value into the specified type
func CastValue(value string, valueType string) (interface{}, error) {
    switch strings.ToLower(valueType) {
    case "integer":
        return strconv.Atoi(value)
    case "string":
        return value, nil
    case "array":
        return strings.Split(value, ","), nil
    case "boolean":
        return strconv.ParseBool(value)
    default:
        return nil, errors.New("unsupported type")
    }
}

// ParseActualOutput tries to parse the actual output string into the type of expected value
func ParseActualOutput(actual string, expected interface{}) (interface{}, error) {
    switch expected.(type) {
    case int:
        return strconv.Atoi(strings.TrimSpace(actual))
    case string:
        return strings.TrimSpace(actual), nil
    case []string:
        // Assume comma-separated output for arrays
        parts := strings.Split(strings.TrimSpace(actual), ",")
        // Trim spaces from each element
        for i := range parts {
            parts[i] = strings.TrimSpace(parts[i])
        }
        return parts, nil
    case bool:
        return strconv.ParseBool(strings.TrimSpace(actual))
    default:
        return nil, errors.New("unsupported expected output type")
    }
}

// IsOutputCorrect compares expected vs actual output after parsing actual output to expected's type
func IsOutputCorrect(expected interface{}, actual string) bool {
    parsedActual, err := ParseActualOutput(actual, expected)
    if err != nil {
        return false
    }
    return reflect.DeepEqual(expected, parsedActual)
}

