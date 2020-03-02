package main

func main() {
	println(intToRoman(3))
	println(intToRoman(4))
	println(intToRoman(9))
	println(intToRoman(58))
	println(intToRoman(1994))
}

func intToRoman(num int) string {
	roman := ""
	if num <= 0 {
		return roman
	}

	numArr := [15]int{1, 2, 3, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romanArr := [15]string{"I", "II", "III", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	for idx := 14; idx >= 0 && num > 0; idx-- {
		for num > numArr[idx] {
			num = num - numArr[idx]
			roman += romanArr[idx]
		}
		if numArr[idx] == num {
			roman += romanArr[idx]
			break
		}
	}

	return roman
}
