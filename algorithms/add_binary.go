package algorithms

import "fmt"

// import "fmt"

func AddBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	carry := 0
	result := ""
	for lenA > 0 || lenB > 0 || carry > 0 {
		digitA := 0
		digitB := 0
		if lenA > 0 {
			digitA = int(a[lenA-1])
		}
		if lenB > 0 {
			digitB = int(b[lenB-1])
		}
		sum := digitA + digitB + carry
		carry = sum / 2
		digit := sum % 2
		result = fmt.Sprintf("%d%s", digit, result)
		lenA--
		lenB--
	}
	return result
}
