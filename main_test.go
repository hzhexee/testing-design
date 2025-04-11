package taxcalculator

import (
	"math"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		income   int
		expected float64
	}{
		{50000, 5000},
		{150000, 22500},
		{600000, 120000},
		{1500000, 375000},
		{0, 0},
		{-50000, -5000},
		{-1, -0.1},
		// Граничные значения
		{100000, 10000},      // Граница первого диапазона
		{100001, 15000.15},   // Первое значение второго диапазона
		{500000, 75000},      // Граница второго диапазона
		{500001, 100000.2},   // Первое значение третьего диапазона
		{1000000, 200000},    // Граница третьего диапазона
		{1000001, 250000.25}, // Первое значение четвертого диапазона
	}

	for _, test := range tests {
		result := calculateTax(test.income)
		// Используем допустимую погрешность при сравнении
		epsilon := 0.00001
		if math.Abs(result-test.expected) > epsilon {
			t.Errorf("For income %d, expected %.2f but got %.2f", test.income, test.expected, result)
		}
	}
}
