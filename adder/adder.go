package adder

import (
	"fmt"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func Add(str1 string, str2 string) (string) {
	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")
	sym1 := ""
	sym2 := ""
	result := ""

	// remove symbol
	if arr1[0] == "-" {
		sym1 = "-"
		arr1 = append(arr1[:0], arr1[1:]...)
	} else if arr1[0] == "+" {
		sym1 = "+"
		arr1 = append(arr1[:0], arr1[1:]...)
	} else {
		sym1 = "+"
	}
	
	if arr2[0] == "-" {
		sym2 = "-"
		arr2 = append(arr2[:0], arr2[1:]...)
	} else if arr2[0] == "+" {
		sym2 = "+"
		arr2 = append(arr2[:0], arr2[1:]...)
	} else {
		sym2 = "+"
	}

	// remove zero in the front
	for {
		if arr1[0] == "0" && len(arr1) != 1 {
			arr1 = append(arr1[:0], arr1[1:]...)
		} else {
			break;
		}
	}

	for {
		if arr2[0] == "0" && len(arr2) != 1 {
			arr2 = append(arr2[:0], arr2[1:]...)
		} else {
			break;
		}
	}

	// add calculate (++ || --)
	if sym1 == sym2 {
		// padding zero
		lenTotal := len(arr1) + len(arr2)
		for {
			if len(arr1) < lenTotal {
				arr1 = append([]string{"0"}, arr1...)
			} else {
				break
			}
		}
		for {
			if len(arr2) < lenTotal {
				arr2 = append([]string{"0"}, arr2...)
			} else {
				break
			}
		}

		// calculate by element
		digi := 0
		for i := len(arr1)-1; i >= 0; i-- {
			temp := Atoi(arr1[i]) + Atoi(arr2[i]) + digi
			if temp >= 10 {
				digi = 1
				result = Itoa(temp - 10) + result
			} else {
				digi =0
				result = Itoa(temp) + result
			}
		}
		if digi == 1 {
			result = Itoa(digi) + result
		}

		// remove zero
		for {
			if string(result[0]) == "0" && len(result) > 1 {
				result = strings.TrimPrefix(result, "0")
			} else {
				break
			}
		}

		// add symbol
		if sym1 == "-" {
			return "-" + result
		} else {
			return result
		}
	}

	// subtract calculate (+L-S || +S-L)
	// 1 > 2
	if len(arr1) > len(arr2) {
		arrLarge := arr1
		symLarge := sym1
		arrSmall := arr2
		resultsStr := ""
		lenTotal := len(arrLarge) + len(arrSmall)
		results := make([]string, lenTotal)
		
		// padding zero
		for {
			if len(arrLarge) < lenTotal {
				arrLarge = append([]string{"0"}, arrLarge...)
			} else {
				break
			}
		}
		for {
			if len(arrSmall) < lenTotal {
				arrSmall = append([]string{"0"}, arrSmall...)
			} else {
				break
			}
		}

		// 1 + 2 -
		for i := range arrLarge {
			if Atoi(arrLarge[i]) >= Atoi(arrSmall[i]) {
				results[i] = Itoa(Atoi(arrLarge[i]) - Atoi(arrSmall[i]))
				} else {
				results[i - 1] = Itoa(Atoi(results[i - 1]) - 1)
				results[i] = Itoa(Atoi(arrLarge[i]) - Atoi(arrSmall[i]) + 10)
			}
		}
		for {
			if results[0] == "0" && len(results) != 1 {
				results = append(results[:0], results[1:]...)
			} else {
				break;
			}
		}
		for i := range results {
			resultsStr = resultsStr + results[i]
		}

		if symLarge == "+" {
			return resultsStr
		} else {
			return "-" + resultsStr
		}
		// 2 > 1
	} else if len(arr2) > len(arr1) {
		arrLarge := arr2
		symLarge := sym2
		arrSmall := arr1
		resultsStr := ""
		lenTotal := len(arrLarge) + len(arrSmall)
		results := make([]string, lenTotal)
		
		// padding zero
		for {
			if len(arrLarge) < lenTotal {
				arrLarge = append([]string{"0"}, arrLarge...)
			} else {
				break
			}
		}
		for {
			if len(arrSmall) < lenTotal {
				arrSmall = append([]string{"0"}, arrSmall...)
			} else {
				break
			}
		}

		// 1 + 2 -
		for i := range arrLarge {
			if Atoi(arrLarge[i]) >= Atoi(arrSmall[i]) {
				
				results[i] = Itoa(Atoi(arrLarge[i]) - Atoi(arrSmall[i]))
			} else {
				
				results[i - 1] = Itoa(Atoi(results[i - 1]) - 1)
				results[i] = Itoa(Atoi(arrLarge[i]) - Atoi(arrSmall[i]) + 10)
			}
		}
		for {
			if results[0] == "0" && len(results) != 1 {
				results = append(results[:0], results[1:]...)
			} else {
				break;
			}
		}
		for i := range results {
			resultsStr = resultsStr + results[i]
		}
		if symLarge == "+" {
			return resultsStr
		} else {
			return "-" + resultsStr
		}
		// same length
	} else {
		for i := range arr1 {
			// 1 > 2
			if arr1[i] > arr2[i] {
				arrLarge := arr1
				symLarge := sym1
				arrSmall := arr2
				resultsStr := ""
				lenTotal := len(arrLarge) + len(arrSmall)
				results := make([]string, lenTotal)
				
				// padding zero
				for {
					if len(arrLarge) < lenTotal {
						arrLarge = append([]string{"0"}, arrLarge...)
					} else {
						break
					}
				}
				for {
					if len(arrSmall) < lenTotal {
						arrSmall = append([]string{"0"}, arrSmall...)
					} else {
						break
					}
				}

				// 1 + 2 -
				for i := range arrLarge {
					if Atoi(arrLarge[i]) >= Atoi(arrSmall[i]) {
						
						results[i] = Itoa(Atoi(arrLarge[i]) - Atoi(arrSmall[i]))
					} else {
						
						results[i - 1] = Itoa(Atoi(results[i - 1]) - 1)
						results[i] = Itoa(Atoi(arrLarge[i]) - Atoi(arrSmall[i]) + 10)
					}
				}
				for {
					if results[0] == "0" && len(results) != 1 {
						results = append(results[:0], results[1:]...)
					} else {
						break;
					}
				}
				for i := range results {
					resultsStr = resultsStr + results[i]
				}
				if symLarge == "+" {
					return resultsStr
				} else {
					return "-" + resultsStr
				}
			} else if arr2[i] > arr1[i] {
				arrLarge := arr2
				symLarge := sym2
				arrSmall := arr1
				resultsStr := ""
				lenTotal := len(arrLarge) + len(arrSmall)
				results := make([]string, lenTotal)
				
				// padding zero
				for {
					if len(arrLarge) < lenTotal {
						arrLarge = append([]string{"0"}, arrLarge...)
					} else {
						break
					}
				}
				for {
					if len(arrSmall) < lenTotal {
						arrSmall = append([]string{"0"}, arrSmall...)
					} else {
						break
					}
				}

				// 1 + 2 -
				for i := range arrLarge {
					if Atoi(arrLarge[i]) >= Atoi(arrSmall[i]) {
						
						results[i] = Itoa(Atoi(arrLarge[i]) - Atoi(arrSmall[i]))
					} else {
						
						results[i - 1] = Itoa(Atoi(results[i - 1]) - 1)
						results[i] = Itoa(Atoi(arrLarge[i]) - Atoi(arrSmall[i]) + 10)
					}
				}
				for {
					if results[0] == "0" && len(results) != 1 {
						results = append(results[:0], results[1:]...)
					} else {
						break;
					}
				}
				for i := range results {
					resultsStr = resultsStr + results[i]
				}
				if symLarge == "+" {
					return resultsStr
				} else {
					return "-" + resultsStr
				}
			}
		}
		return "0"
	}
}

func main(){
	fmt.Println(Add("123", "456") == "579")
	fmt.Println(Add("-123", "-456") == "-579")
	fmt.Println(Add("999", "1") == "1000")
	fmt.Println(Add("-999", "-1") == "-1000")
	fmt.Println(Add("1", "999") == "1000")
	fmt.Println(Add("-1", "-999") == "-1000")
	fmt.Println(Add("33", "-44") == "-11")
	fmt.Println(Add("-44", "33") == "-11")
	fmt.Println(Add("0", "-9") == "-9")
	fmt.Println(Add("-9", "0") == "-9")
	fmt.Println(Add("9", "-10") == "-1")
	fmt.Println(Add("-10", "9") == "-1")
	fmt.Println(Add("-98", "9") == "-89")
	fmt.Println(Add("9", "-98") == "-89")
	fmt.Println(Add("123", "-456") == "-333")
	fmt.Println(Add("-456", "123") == "-333")
	fmt.Println(Add("123", "-4567") == "-4444")
	fmt.Println(Add("-4567", "123") == "-4444")
	fmt.Println(Add("999", "-456") == "543")
	fmt.Println(Add("-456", "999") == "543")
	fmt.Println(Add("999999", "-1000000") == "-1")
	fmt.Println(Add("-1000000", "999999") == "-1")
	fmt.Println(Add("01", "02") == "3")
	fmt.Println(Add("-001", "02") == "1")
	fmt.Println(Add("+1", "-1") == "0")
}