
# Tax Calculator API

This is a simple Go-based Tax Calculator API. It calculates the tax owed based on different tax brackets and returns the result in JSON format. It allows you to input annual income and tax year and provides tax calculations based on predefined brackets.

Features
Calculate Tax Owed: Calculates tax owed based on income and year.
Support for Multiple Tax Brackets: Tax brackets with different rates, including a last bracket with no upper limit.
Formatted Results: The tax owed is calculated and displayed with two decimal points for precision.

## Installation

To run this project locally, follow these steps:

Prerequisites
- Go installed on your machine (version 1.16 or higher).
- Go Modules enabled.


1. Steps
Extract the .zip archive into a folder.

Navigate into the project directory:

```bash
  cd TaxCalcPoints
```
 2. Install Go dependencies (if needed):

```bash
  go mod tidy
```
 4. Run the API:

```bash
  go run main.go
```
  4. Open url:

```bash
 http://localhost:8000/tax?income=50000&year=2019
```
                   
## QA

Unit test

```bash
  go test -v ./models  
  go test -v ./controllers   
  go test -v ./models  
  go test -v ./routes  
```

