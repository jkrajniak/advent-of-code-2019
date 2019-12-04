package main

import (
	"strconv"
	"strings"
)

func CheckPassword(passCode []int) bool {
	// Length of 6 digits
	if len(passCode) != 6 {
		return false
	}
	diff := getDiff(passCode)

	// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
	if negativeInSlice(diff) {
		return false
	}

	// Two adjacent digits are the same
	if !inSlice(0, diff) {
		return false
	}

	return true
}

func CheckPasswordExtended(passCode []int) bool {
	// Length of 6 digits
	if len(passCode) != 6 {
		return false
	}
	diff := getDiff(passCode)

	// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
	if negativeInSlice(diff) {
		return false
	}

	// Two adjacent digits are the same but not three
	if !hasAtMostTwoDigitsInBlock(passCode) {
		return false
	}

	return true
}

func GetPassCodes(startRange, endRange int) []int64 {
	var passCodes []int64
	//var wg sync.WaitGroup
	//wg.Add(stopRange - startRange + 1)

	for i := startRange; i <= endRange; i++ {
		//go func (rawPassCode int64) {
		//	defer wg.Done()
		passCode := convertToArrayOfInts(int64(i))
		if CheckPassword(passCode) {
			passCodes = append(passCodes, int64(i))
		}
		//}(int64(i))
	}
	//wg.Wait()
	return passCodes
}

func GetPassCodesExtended(startRange, endRange int) []int64 {
	var passCodes []int64
	//var wg sync.WaitGroup
	//wg.Add(stopRange - startRange + 1)

	for i := startRange; i <= endRange; i++ {
		//go func (rawPassCode int64) {
		//	defer wg.Done()
		passCode := convertToArrayOfInts(int64(i))
		if CheckPasswordExtended(passCode) {
			passCodes = append(passCodes, int64(i))
		}
		//}(int64(i))
	}
	//wg.Wait()
	return passCodes
}

func convertToArrayOfInts(i int64) []int {
	s := strings.Split(strconv.FormatInt(i, 10), "")
	vsm := make([]int, len(s))
	for i, v := range s {
		vsm[i], _ = strconv.Atoi(v)
	}
	return vsm
}

func getDiff(passCode []int) []int {
	diff := make([]int, len(passCode)-1)
	for i := 1; i < len(passCode); i++ {
		diff[i-1] = passCode[i] - passCode[i-1]
	}
	return diff
}

func inSlice(item int, items []int) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func negativeInSlice(items []int) bool {
	for _, i := range items {
		if i < 0 {
			return true
		}
	}
	return false
}

func hasAtMostTwoDigitsInBlock(passCode []int) bool {
	blocks := []int{}
	blockSize := 1
	for i := 0; i < len(passCode)-1; i++ {
		if passCode[i] == passCode[i+1] {
			blockSize++
		} else {
			blocks = append(blocks, blockSize)
			blockSize = 1
		}
	}
	blocks = append(blocks, blockSize)
	for _, b := range blocks {
		if b == 2 {
			return true
		}
	}

	return false

}