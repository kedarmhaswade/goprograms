package second

import "strings"

func RemoveNonNumeric1(phoneNo string) string {
	replacer := strings.NewReplacer(" ", "", "-", "", "(", "", ")", "")
	return replacer.Replace(phoneNo)
}
func RemoveNonNumeric2(phoneNo string) string {
	strLen := len(phoneNo)
	buf := make([]byte, strLen, strLen) // minimize allocation
	j := 0
	for i := 0; i < strLen; i++ {
		b := phoneNo[i]
		if b != ' ' && b != '-' && b != '(' && b != ')' {
			buf[j] = b // not using append
			j += 1
		}
	}
	return string(buf[:j])
}
