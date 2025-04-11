package taxcalculator

func calculateTax(income int) float64 {
	incomeFloat := float64(income)
	var taxRate float64
	if income <= 100000 {
		taxRate = 0.1
	} else if income > 100000 && income <= 500000 {
		taxRate = 0.15
	} else if income > 500000 && income <= 1000000 {
		taxRate = 0.2
	} else {
		taxRate = 0.25
	}

	taxAmount := incomeFloat * taxRate
	return taxAmount
}