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
	//  fmt.Println(input)
	input = strings.Replace(input, " ", "", -1)
	runex := []rune(input)
	lenArr := utf8.RuneCountInString(input)
	//  fmt.Println(lenArr)
	if lenArr == 0 {
		err = fmt.Errorf("%w", errorEmptyInput)
		return "", err
	}
	var t3Sign, fSign, tSign, fDig, tDig, FinalStr string

	for i := 0; i <= lenArr-1; i++ {

		switch {
		case (runex[i] == '+' || runex[i] == '-') && fDig == "":
			fSign = string(runex[i])
		case (runex[i] == '+' || runex[i] == '-') && fDig != "" && tDig == "":
			tSign = string(runex[i])
		case (runex[i] == '+' || runex[i] == '-'):
			//	fmt.Println("ааааа")
			t3Sign = string(runex[i])
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		case (runex[i] >= '0' && runex[i] <= '9') && tSign == "":
			fDig += string(runex[i])
		case (runex[i] >= '0' && runex[i] <= '9') && t3Sign == "" && tSign != "":
			tDig += string(runex[i])
		case (runex[i] >= '0' && runex[i] <= '9'):
			t3Sign = string(runex[i])
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		default:
			// fmt.Println(string(runex[i]))
			err = fmt.Errorf("attention in line detected NLO")
			return "", err
		}
	}
	// fmt.Println(fSign, fDig, tSign, tDig)

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
	// fmt.Println(finalint)
	FinalStr = strconv.Itoa(finalint)
	// fmt.Println(FinalStr)

	return FinalStr, nil
}
