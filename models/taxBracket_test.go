package models

import (
	"os"
	"testing"
)

// floatPtr is a helper function to create a pointer from a float64
func floatPtr(value float64) *float64 {
	return &value
}

// TestCalcTax verifies if CalcTax correctly calculates the tax based on different brackets
func TestCalcTax(t *testing.T) {
	// Define test tax brackets
	brackets := []Bracket{
		{MinIncome: 0, MaxIncome: floatPtr(50000), TaxRate: 0.15},
		{MinIncome: 50000, MaxIncome: floatPtr(100000), TaxRate: 0.205},
		{MinIncome: 100000, MaxIncome: nil, TaxRate: 0.26}, // Last bracket has no upper limit
	}

	// Test cases
	tests := []struct {
		income   float64
		expected float64
	}{
		{income: 40000, expected: 6000},
		{income: 60000, expected: 9550.00},
		{income: 120000, expected: 22950.00},
		{income: 50000, expected: 7500.00},
		{income: 100000, expected: 17750.00},
	}

	// Loop through test cases
	for _, test := range tests {
		result := CalcTax(test.income, brackets)
		if result != test.expected {
			t.Errorf("For income %.2f, expected %.2f, but got %.2f", test.income, test.expected, result)
		}
	}
}

// TestLoadTaxBrackets checks if LoadTaxBrackets correctly loads and parses the JSON data
func TestLoadTaxBrackets(t *testing.T) {
	// Create a temporary JSON file for testing
	jsonData := `[
    {
        "min": 0,
        "max": 53359,
        "rate": 0.15
    },
    {
        "min": 53359,
        "max": 106717,
        "rate": 0.205
    },
    {
        "min": 106717,
        "max": 165430,
        "rate": 0.26
    },
    {
        "min": 165430,
        "max": 235675,
        "rate": 0.29
    },
    {
        "min": 235675,
        "rate": 0.33
    }
]`
	tmpFile := "taxBracketsByYears/taxBrackets--2019.json"

	// Create directory if it does not exist
	_ = os.Mkdir("taxBracketsByYears", os.ModePerm)

	// Write the JSON data to the file
	err := os.WriteFile(tmpFile, []byte(jsonData), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile) // Cleanup after test

	// Run LoadTaxBrackets function
	brackets, err := LoadTaxBrackets("2019")
	if err != nil {
		t.Errorf("LoadTaxBrackets failed: %v", err)
	}

	// Validate the results
	expectedLen := 5
	if len(brackets) != expectedLen {
		t.Errorf("Expected %d brackets, but got %d", expectedLen, len(brackets))
	}

	expectedFirstRate := 0.15
	if brackets[0].TaxRate != expectedFirstRate {
		t.Errorf("Expected first bracket rate to be %.2f, but got %.2f", expectedFirstRate, brackets[0].TaxRate)
	}
}
