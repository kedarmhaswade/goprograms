package second

import "strings"

func RemoveNonNumeric1(phoneNo string) string {
	replacer := strings.NewReplacer(" ", "", "-", "", "(", "", ")", "")
	return replacer.Replace(phoneNo)
}
func RemoveNonNumeric2(phoneNo string) string {
	strLen := len(phoneNo)
	buf := make([]byte, 0, strLen) // minimize allocation
	l := 0
	for i := 0; i < strLen; i++ {
		b := phoneNo[i]
		if b >= '0' && b <='9' {
			buf = append(buf, b)
			l += 1
		}
	}
	return string(buf[:l])
}
