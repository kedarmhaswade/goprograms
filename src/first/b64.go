package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {

	uuidstr := "B831964C-CC24-469E-8E7E-0D6320D8B48A"
	hexBytes, err := hex.DecodeString(strings.ReplaceAll(uuidstr, "-", ""))
	if err != nil {
		return
	}
	b64Str := base64.RawURLEncoding.EncodeToString(hexBytes)
	fmt.Printf("%s, len(b64str)= %v, len(uuid)=%v\n", b64Str, len(b64Str), len(uuidstr))
	//_, _ := base64.RawURLEncoding.DecodeString(b64Str)


}
