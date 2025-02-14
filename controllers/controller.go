package controllers

import (
	"TaxCalcPoints/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// TaxHandler handles the request to calculate tax based on income and year
func TaxHandler(w http.ResponseWriter, r *http.Request) {
	// Get parameters From URL
	GetIncome := r.URL.Query()
	incomeReturn := GetIncome.Get("income")
	incomeyear := GetIncome.Get("year")

	// Converts the income input from string to float
	income, err := strconv.ParseFloat(incomeReturn, 64)
	if err != nil || income < 0 {
		http.Error(w, `"error": "Income isn't a valid number"`, http.StatusBadRequest)
		return
	}

	// Loads the tax brackets based on the year
	taxData, err := models.LoadTaxBrackets(incomeyear)
	if err != nil {
		http.Error(w, `"error": "Failed to load tax data"`, http.StatusBadRequest)
		return
	}

	tax := models.CalcTax(income, taxData)

	// JSON response
	res := map[string]interface{}{
		"income":   income,
		"tax_owed": fmt.Sprintf("%.2f", tax),
		"year":     incomeyear,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
