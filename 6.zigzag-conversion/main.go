package main

func main() {
	s := "PAYPALISHIRING"
	println(convert(s, 4))
	println(convert(s, 3))
}

func convert(s string, numRows int) string {
	if s == "" || numRows == 1 || len(s) <= numRows {
		return s
	}
	length := len(s)
	strByte := make([]byte, length)
	step := 2 * (numRows - 1)

	k := 0
	num := (length-1)/step + 1
	remain := (length - 1) % step
	for i := 0; i < numRows; i++ {
		if i == 0 {
			for j := 0; j < num; j++ {
				strByte[k] = s[j*step]
				k++
			}
		} else if i == numRows-1 {
			for j := 0; j < num; j++ {
				if (j != num-1) || (j == (num-1) && remain >= (numRows-1)) {
					strByte[k] = s[j*step+numRows-1]
					k++
				}
			}
		} else {
			for j := 0; j < num; j++ {
				if j*step+i < length {
					strByte[k] = s[j*step+i]
					k++
				}
				if (j+1)*step-i < length {
					strByte[k] = s[(j+1)*step-i]
					k++
				}
			}
		}
	}
	return string(strByte)
}
