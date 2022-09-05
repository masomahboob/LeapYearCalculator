package main

import (
	"fmt"
)

func main() {

	leapYearCalc(0)

}

func leapYearCalc(choice int) []int {

	var result []int

	// determin if a given year is a leap year
	if choice == 0 {

		var year int

		fmt.Print("Please enter a year: ")

		fmt.Scan(&year)

		if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
			result = append(result, year)
			fmt.Printf("%d is a leap year.\n", year)
		} else {
			fmt.Printf("%d is not a leap year.\n", year)
		}

		// determin multiple years to be leap year or not
	} else if choice == 1 {

		startingYear := 2000

		var upToYear int

		fmt.Print("Please enter a year: ")

		fmt.Scan(&upToYear)

		for y := startingYear; y <= upToYear; y++ {
			if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
				result = append(result, y)
			}
		}

		fmt.Println(result)

	} else {
		fmt.Println("Please choose a viable option! (0 or 1)")
	}

	return result
}
