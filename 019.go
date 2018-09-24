package main

import (
	"fmt"
)

func main() {
	startDay := 1     // 1-31
	startMonth := 1   // 1-12
	startYear := 1901 // n
	weekDay := 2      // 1(monday)..7(sunday)

	targetDay := 31    // 1-31
	targetMonth := 12  // 1-12
	targetYear := 2000 // n

	sundays := 0 // counter

	for year := startYear; year <= targetYear; year++ {
		month := 1
		endMonth := 12

		// skip useless months of first year
		if year == startYear {
			month = startMonth
		}

		// skip unless months of last year
		if year == targetYear {
			endMonth = targetMonth
		}

		for ; month <= endMonth; month++ {
			day := 1

			// skip useless days of first month of first year
			if year == startYear && month == startMonth {
				day = startDay
			}

			nDays := 31

			if month == 4 || month == 6 || month == 9 || month == 11 {
				nDays = 30
			}

			if month == 2 {
				nDays = 28

				// Leap year (year divisible by 4 but not centuries unless it's divisible by 400)
				if year%4 == 0 && (!(year%100 == 0) || year%400 == 0) {
					nDays = 29
				}
			}

			// skip useless days of last month of last year
			if year == targetYear && month == targetMonth {
				nDays = targetDay
			}

			for ; day <= nDays; day++ {
				// actual problem condition
				if weekDay == 7 && day == 1 {
					sundays++
				}

				// loop weekday 1 (monday) .. 7 (sunday)
				if weekDay < 7 {
					weekDay++
				} else {
					weekDay = 1
				}
			}
		}
	}

	fmt.Println("== Euler project - Problem 19 ==")
	fmt.Println("Sundays falling on the first of the month:", sundays)
}
