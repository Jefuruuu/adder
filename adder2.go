package adder31plus

import (
	"fmt"
	"strconv"
	"strings"
)

type Number struct {
	Sign string
	Value []int
}

func Add(val1 []int, val2 []int) string {
	temp := 0
	res := ""
	for i := len(val1)-1; i >= 0; i -- {
		digit := val1[i] + val2[i] + temp
		if digit > 9 {
			temp = 1
			res = strconv.Itoa(digit - 10) + res
		} else {
			temp = 0
			res = strconv.Itoa(digit) + res
		}
	}
	if temp == 1 {
		res = strconv.Itoa(temp) + res
	}
	return res
}

func Subtract(val1 []int, val2 []int) string {
	res := make([]int, len(val1))
	for i := range val1 {
		if val1[i] < val2[i] {
			res[i - 1] = res[i - 1] - 1
			res[i] = 10 + val1[i] - val2[i]
		} else {
			res[i] = val1[i] - val2[i]
		}
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), ""), "[]")
}

func PaddingZero(val []int, digit int) []int {
	zeros := make([]int, digit-len(val))
	val = append(zeros, val...)
	return val
}

func CompareValue(val1 []int, val2 []int) string {
	for i := range val1 {
		if val1[i] > val2[i] {
			return "True"
		} else if val1[i] < val2[i] {
			return "False"
		}
	}
	return "Same"
}

func RemoveZero(str string) string {
	strTrim := ""
	for {
		strTrim = strings.TrimPrefix(str, "0")
		if str == strTrim || len(str) == 1 {
			return str
		}
		str = strTrim
	}
}

func ExtractSignAndNum(str string) (string, string) {
	switch string(str[0]) {
	case "-":
		return "-", str[1:]
	case "+":
		return "+", str[1:]
	default:
		return "+", str[:]
	}
}

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

func Adder(num1 string, num2 string) string {
	number1 := StrToNumber(num1)
	number2 := StrToNumber(num2)

	if len(number1.Value) > len(number2.Value) {
		number2.Value = PaddingZero(number2.Value, len(number1.Value))
	} else if len(number1.Value) < len(number2.Value) {
		number1.Value = PaddingZero(number1.Value, len(number2.Value))
	}

	if number1.Sign == number2.Sign {
		return strings.TrimPrefix(number1.Sign + RemoveZero(Add(number1.Value, number2.Value)), "+")
	}
	if CompareValue(number1.Value, number2.Value) == "True" {
		res := RemoveZero(Subtract(number1.Value, number2.Value))
		return strings.TrimPrefix(number1.Sign + res, "+")
	} else if CompareValue(number1.Value, number2.Value) == "False" {
		res := RemoveZero(Subtract(number2.Value, number1.Value))
		return strings.TrimPrefix(number2.Sign + res, "+")
	} else {
		return "0"
	}
}