package adder31plus

import (
	"fmt"
	"strconv"
	"strings"
)

type Outcome int

const (
	Same Outcome = iota
	Larger
	Less
)

type Number struct {
	Sign  string
	Value []int
}

func Add(val1 []int, val2 []int) string {
	temp := 0
	res := make([]string, len(val1)+1)
	for i := len(val1) - 1; i >= 0; i-- {
		digit := val1[i] + val2[i] + temp
		temp = digit / 10
		digit = digit % 10
		res[i+1] = strconv.Itoa(digit)
	}
	if temp == 1 {
		res[0] = "1"
	} else {
		res = res[1:]
	}
	return strings.Join(res, "")
}

func Subtract(val1 []int, val2 []int) string {
	res := make([]int, len(val1))
	for i := range val1 {
		if val1[i] < val2[i] {
			res[i-1] = res[i-1] - 1
			res[i] = 10 + val1[i] - val2[i]
		} else {
			res[i] = val1[i] - val2[i]
		}
	}
	// Ex: (XXX is the last output)
	// res						-> [3]{1,2,3}
	// fmt.Sprint(res)			-> "[1 2 3]"
	// strings.Fields(XXX)		-> [5]string{"[", "1", "2", "3", "]"}
	// strings.Join(XXX, "")	-> "[123]"
	// strings.Trim(XXX, "[]")	-> "123"
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), ""), "[]")
}

func PaddingZero(val []int, digit int) []int {
	if len(val) >= digit {
		return val
	}
	zeros := make([]int, digit-len(val))
	return append(zeros, val...)
}

// Compare if val1 > val2
func CompareValue(val1 []int, val2 []int) Outcome {
	for i := range val1 {
		if val1[i] > val2[i] {
			return Larger
		} else if val1[i] < val2[i] {
			return Less
		}
	}
	return Same
}

// Remove the first few zero from string except only one zero
func RemoveZero(str string) string {
	idx := 0
	for ; idx < len(str) && string(str[idx]) == "0"; idx++ {
	}

	return str[idx:]
}

// Get sign and value
func ExtractSignAndNum(str string) (string, string) {
  firstChar := string(str[0])
	switch firstChar {
	case "-", "+":
		return firstChar, str[1:]
	default:
		return "+", str[:]
	}
}

// Parse input string to number structure
func StrToNumber(str string) Number {
	var number = Number{}
	sign, numStr := ExtractSignAndNum(str)
	number.Sign = sign
	number.Value = make([]int, 0, len(numStr))

	for _, val := range strings.Split(RemoveZero(numStr), "") {
		valInt, _ := strconv.Atoi(val)
		number.Value = append(number.Value, valInt)
	}
	return number
}

// Input two strings of number, return the result of adding
func Adder(num1 string, num2 string) string {
	number1 := StrToNumber(num1)
	number2 := StrToNumber(num2)

	// pad zero depends on the larger num
	// Ex: [2]int{1,2}, [3]int{1,2,3} -> [3]{0,1,2}, [3]int{1,2,3}
	number2.Value = PaddingZero(number2.Value, len(number1.Value))
	number1.Value = PaddingZero(number1.Value, len(number2.Value))

	// If signs are the same, do add() and add sign on the result
	if number1.Sign == number2.Sign {
		return strings.TrimPrefix(number1.Sign+RemoveZero(Add(number1.Value, number2.Value)), "+")
	}

	// If signs are different, use large value to subtract small value
	// If value of negative input is larger, add "-" to the result
	// If values are the same, return 0
	switch CompareValue(number1.Value, number2.Value) {
	case Larger:
		res := RemoveZero(Subtract(number1.Value, number2.Value))
		return strings.TrimPrefix(number1.Sign+res, "+")
	case Less:
		res := RemoveZero(Subtract(number2.Value, number1.Value))
		return strings.TrimPrefix(number2.Sign+res, "+")
	default:
		return "0"
	}
}

