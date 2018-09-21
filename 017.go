package main

import (
	"fmt"
	"math"
	"strings"
	//"strconv"
)

func main() {
	const max = 1000
	const skipConcat = false
	const debug = false

	lettersCount := 0

	for i := 1; i <= max; i++ {
		s := getWordForInt(i)

		if !skipConcat {
			s = strings.Replace(s, " ", "", -1)
			s = strings.Replace(s, "-", "", -1)
		}

		if debug {
			fmt.Println(s)
		}

		lettersCount += len(s)
	}

	fmt.Println("== Euler project - Problem 17 ==")
	fmt.Println(lettersCount)
}

func getWordForInt(n int) string {
	below20 := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if n < 20 {
		return below20[n-1]
	}

	if n < 100 {
		tensIndex := int(math.Floor(float64(n) / 10))
		remaining := n - tensIndex*10
		remainingStr := ""

		if remaining > 0 {
			remainingStr = "-" + getWordForInt(remaining)
		}

		tensStr := tens[tensIndex-2] + remainingStr
		return tensStr
	}

	if n < 1000 {
		hundredsIndex := int(math.Floor(float64(n) / 100))
		remaining := n - hundredsIndex*100
		hundredsStr := ""

		hundredsStr += below20[hundredsIndex-1] + " hundred"

		if remaining > 0 {
			hundredsStr += " and " + getWordForInt(remaining)
		}

		return hundredsStr
	}

	if n < 10000 {
		thousands := int(math.Floor(float64(n) / 1000))
		remaining := n - thousands*1000

		thousandsStr := getWordForInt(thousands) + " thousand"

		if remaining > 0 {
			if remaining < 100 {
				thousandsStr += " and "
			} else {
				thousandsStr += ", "
			}
			thousandsStr += getWordForInt(remaining)
		}

		return thousandsStr
	}

	return ""
}
