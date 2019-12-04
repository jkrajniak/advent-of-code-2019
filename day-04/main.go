package main

import (
	"fmt"
)

func main() {
	startRange := 284639
	stopRange := 748759

	passCodes := GetPassCodes(startRange, stopRange)
	fmt.Println(len(passCodes))

	passCodesExtended := GetPassCodesExtended(startRange, stopRange)
	fmt.Println(len(passCodesExtended))

}


