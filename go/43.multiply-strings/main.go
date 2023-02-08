package main

func main() {
	println(multiply("2", "3"))
	println(multiply("123", "456"))
	println(multiply("0", "0"))
}

func multiply(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)
	res := make([]int, n1+n2)
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			mul := int((num1[i] - '0') * (num2[j] - '0'))
			p1, p2 := i+j, i+j+1
			sum := mul + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}
	i := 0
	for ; i < len(res); i++ {
		if res[i] != 0 {
			break
		}
	}
	str := ""
	for ; i < len(res); i++ {
		str += string('0' + int32(res[i]))
	}
	if str == "" {
		return "0"
	}
	return str
}
