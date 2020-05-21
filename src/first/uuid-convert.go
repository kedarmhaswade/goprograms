package main

// There are so many UUID's

import (
	"encoding/hex"
	"fmt"
	gofrs "github.com/gofrs/uuid"
	goog "github.com/google/uuid"
	"strings"
)

func main() {
	googUuid := goog.New()
	googUuidStrWithHyphens := googUuid.String()
	fmt.Printf("goog uuid str w hyphens: %v, len: %v\n", googUuidStrWithHyphens, len(googUuidStrWithHyphens))
	googUuidStrNoHyphens := strings.ReplaceAll(googUuidStrWithHyphens, "-", "")
	fmt.Printf("goog uuid str wo hyphens: %v, len: %v\n", googUuidStrNoHyphens, len(googUuidStrNoHyphens))
	h, _ := hex.DecodeString(googUuidStrNoHyphens)
	//fmt.Printf("decoded hex: %v\n", h)
	var b [16]byte
	copy(b[:], h)
	gofrsUUID := gofrs.UUID(b)
	gofrsUUIDStr := gofrsUUID.String()
	fmt.Printf("gofrs UUID str: %v\n", gofrsUUIDStr)
	fmt.Printf("uuids match: %v\n", strings.EqualFold(gofrsUUIDStr, googUuidStrWithHyphens))
}
