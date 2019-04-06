package golangexamples

import (
	"fmt"
	"strings"
)

func ConcatSlice(slicetoConcat []byte) string {
	var str []string
	for i := 0; i < len(slicetoConcat); i++ {
		str = append(str, string(slicetoConcat[i]))

	}
	return strings.Join(str, "-")
}

func Encrypt(slicetoEncrypt []byte, ceaserCount int) {
	var str []string
	for i := 0; i < len(slicetoEncrypt); i++ {
		s := int(slicetoEncrypt[i]) + ceaserCount
		if s > 'z' {
			s = (s - 26)
		}
		str = append(str, string(s))
	}
	fmt.Println(str, "\n")
}
