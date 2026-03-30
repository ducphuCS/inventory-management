package services

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// ProductAnalysis represents the structured output of the Python analysis script.
type ProductAnalysis struct {
	ProductName     string `json:"product_name"`
	PopularityScore int    `json:"popularity_score"`
	RestockPriority string `json:"restock_priority"`
	LastMarketTrend string `json:"last_market_trend"`
	PythonEngine    string `json:"python_engine"`
	Error           string `json:"error"`
}

// AnalyzeWithPython provides an interface to execute the Python analysis script.
// It returns a ProductAnalysis object by capturing and parsing the script's stdout.
func AnalyzeWithPython(productName string) (*ProductAnalysis, error) {
	// 1. Execute using 'uv run' to ensure the correct environment and dependencies.
	// We call 'python3' through 'uv' to handle virtual environments automatically.
	cmd := exec.Command("uv", "run", "python3", "scripts/analyze_product.py", productName)

	// 2. Execute the command and capture the standard output.
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute python script: %v", err)
	}

	// 3. Unmarshal the JSON output into our Go struct.
	var analysis ProductAnalysis
	if err := json.Unmarshal(output, &analysis); err != nil {
		return nil, fmt.Errorf("failed to parse python script output: %v", err)
	}

	// 4. Return the parsed analysis.
	return &analysis, nil
}
