package letcode_test

import (
	"fmt"
	"testing"
)

func cov(str string) int {
	fmt.Println(str)
	if str == "" {
		return 0
	}

	var i, s int

	for i < len(str) && i < 18 {
		if str[i] > '9' || str[i] < '0' {
			return s
		}
		s = s*10 + int(str[i]-'0')
		fmt.Println(s)
		i++
	}
	return s
}

const INT_MAX = 2147483647
const INT_MIN = -2147483648

func check(i int) int {
	fmt.Println(i)
	if i > INT_MAX {
		return INT_MAX
	}

	if i < INT_MIN {
		return INT_MIN
	}

	return i
}

func myAtoi(str string) int {
	if str == "" {
		return 0
	}

	var i, s int
	for i < len(str) {
		if str[i] != ' ' {
			s = i
			break
		}
		i++
	}

	str = str[s:]

	if len(str) > 0 && (str[0] > '9' || str[0] < '0') {
		if str[0] != '+' && str[0] != '-' {
			return 0
		}
		sign := 1

		if str[0] == '-' {
			sign = -1
		}

		return check(sign * cov(str[1:]))

	}
	return check(cov(str))
}
func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("9223372036854775809"))
	fmt.Println(int32(0x7FFFFFFF))

}
