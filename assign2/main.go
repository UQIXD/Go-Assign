package main

import "fmt"

func minAndMaxTransaction(expenseList []int) (float64, int, int) {
	if len(expenseList) == 0 {
		return 0, 0, 0 // Return default values if the list is empty
	}

	// Initialize variables
	sum := 0
	min := expenseList[0]
	max := expenseList[0]

	// Iterate over the expense list
	for _, expense := range expenseList {
		sum += expense
		if expense < min {
			min = expense
		}
		if expense > max {
			max = expense
		}
	}

	// Calculate the mean
	mean := float64(sum) / float64(len(expenseList))

	return mean, min, max
}

func main() {
	expenseList1 := []int{50000, 30000, 80000, 45000, 72000, 58000, 65000}
	fmt.Println(minAndMaxTransaction(expenseList1)) // output: 57142.857143 30000 80000

	expenseList2 := []int{80000, 60000, 30000, 45000, 72000, 58000, 20000, 15000}
	fmt.Println(minAndMaxTransaction(expenseList2)) // output: average, min, max

	expenseList3 := []int{50000, 60000, 10000, 75000, 10000, 75000, 65000, 45000, 11000}
	fmt.Println(minAndMaxTransaction(expenseList3)) // output: average, min, max
}
