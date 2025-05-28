package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

// zap extractLevel
func ZapExtractLevel(p []byte) string {
	var temp map[string]interface{}
	if err := json.Unmarshal(p, &temp); err != nil {
		return "UNKNOWN"
	}
	fmt.Print("temp: ", temp)
	if lvl, ok := temp["level"].(string); ok {
		return strings.ToLower(lvl)
	}
	return "UNKNOWN"
}

// zap extractMessage
func ZapExtractMessage(p []byte) string {
	var temp map[string]interface{}
	if err := json.Unmarshal(p, &temp); err != nil {
		return "UNKNOWN"
	}
	if msg, ok := temp["msg"].(string); ok {
		return msg
	}
	return "UNKNOWN"
}

// zap extractTimestamp
func ZapExtractTimestamp(p []byte) string {
	var temp map[string]interface{}
	if err := json.Unmarshal(p, &temp); err != nil {
		return "UNKNOWN"
	}
	if ts, ok := temp["time"].(string); ok {
		return ts
	}
	return "UNKNOWN"
}