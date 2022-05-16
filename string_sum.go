package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf
// Проверить строку на (только) пробелы
// Проверить строку на (только) пробелы
func StringSum(input string) (output string, err error) {

	input = strings.Replace(input, " ", "", -1)
	runex := []rune(input)
	lenArr := utf8.RuneCountInString(input)
	if lenArr == 0 {
		return "", errorEmptyInput
	}
	var fSign, tSign, fDig, tDig, finalStr string

	for i := 0; i <= lenArr-1; i++ {
		switch {
		case (runex[i] == '+' || runex[i] == '-') && fDig == "":
			fSign = string(runex[i])
		case (runex[i] == '+' || runex[i] == '-') && fDig != "":
			tSign = string(runex[i])
		case (runex[i] > '0' && runex[i] < '9') && fDig == "":
			fDig = string(runex[i])
		case (runex[i] > '0' && runex[i] < '9') && fDig != "":
			tDig = string(runex[i])
		default:
			err = fmt.Errorf("attention in line detected NLO")
			return "", err
		}
	}

	if fDig == "" || tDig == "" {
		return "", errorNotTwoOperands

	}

	finalint1, err1 := strconv.Atoi(fSign + fDig)

	finalint2, err2 := strconv.Atoi(tSign + tDig)

	if err1 != nil || err2 != nil {

		err = fmt.Errorf("attention in line detected Super NLO")
		return "", err
	}

	finalint := finalint1 + finalint2

	finalStr = strconv.Itoa(finalint)

	return finalStr, nil
}
