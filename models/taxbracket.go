package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// Bracket represents a tax bracket with min and max income and corresponding tax rate
type Bracket struct {
	MinIncome float64  `json:"min"`
	MaxIncome *float64 `json:"max,omitempty"`
	TaxRate   float64  `json:"rate"`
}

// CalcTax calculates the total tax based on income and the tax brackets
func CalcTax(income float64, brackets []Bracket) float64 {
	tax := 0.0
	for _, bracket := range brackets {
		if income > bracket.MinIncome {
			// Set max limit to income if it's within the current bracket
			maxLimit := bracket.MaxIncome
			if maxLimit == nil || income < *maxLimit {
				maxLimit = &income
			}
			// Calculate tax for this bracket
			tax += (*maxLimit - bracket.MinIncome) * bracket.TaxRate
		} else {
			break
		}
	}
	return tax
}

// LoadTaxBrackets loads the tax brackets from a JSON file based on the provided year
func LoadTaxBrackets(year string) ([]Bracket, error) {
	// Monta o caminho do arquivo com base no ano
	filePath := fmt.Sprintf("taxBracketsByYears/taxBrackets--%s.json", year)

	// Abre o arquivo JSON
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decodifica o JSON diretamente do arquivo
	var taxData []Bracket
	err = json.NewDecoder(file).Decode(&taxData)
	if err != nil {
		return nil, err
	}

	return taxData, nil
}
