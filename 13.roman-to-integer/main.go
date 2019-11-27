package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	if "" == s {
		return 0
	}
	romanMap := map[string]int {
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	sum := 0
	for i:=0; i < len(s); i++ {
		if i+1 > len(s) -1 {
			k := s[i:i+1]
			if _, ok := romanMap[k]; ok {
				sum += romanMap[k]
			}
		} else {
			k := s[i:i+2]
			if _, ok := romanMap[k]; ok {
				sum += romanMap[k]
				i++
			} else {
				k := s[i:i+1]
				if _, ok := romanMap[k]; ok {
					sum+=romanMap[k]
				}
			}
		}
	}
	return sum
}
