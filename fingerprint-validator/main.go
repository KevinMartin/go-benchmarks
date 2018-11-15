package fingerprintValidator

import (
	"strconv"
	"strings"
)

// ValidateStringsSplitStringsCount counts colons using `strings.Count` then
// loops through the pairs after splitting with `strings.Split`.
func ValidateStringsSplitStringsCount(pairs int, fp string) bool {
	if strings.Count(fp, ":") != pairs-1 {
		return false
	}

	hexes := strings.Split(fp, ":")
	for i := range hexes {
		if len(hexes[i]) != 2 {
			return false
		}
		if _, err := strconv.ParseUint(hexes[i], 16, 64); err != nil {
			return false
		}
	}
	return true
}

// ValidateStringsSplitStrings loops through the pairs after splitting with `strings.Split`.
func ValidateStringsSplitStrings(pairs int, fp string) bool {
	// 3 = len(hex) + len(":")
	if len(fp) != pairs*3-1 {
		return false
	}

	hexes := strings.Split(fp, ":")
	for i := range hexes {
		if len(hexes[i]) != 2 {
			return false
		}
		if _, err := strconv.ParseUint(hexes[i], 16, 64); err != nil {
			return false
		}
	}
	return true
}

// ValidateLoopPairsCount creates a loop based on expected number of pairs.
func ValidateLoopPairsCount(pairs int, fp string) bool {
	// 3 = len(hex) + len(":")
	if len(fp) != pairs*3-1 {
		return false
	}

	for i := 0; i < pairs; i++ {
		pos := i * 3
		hex := fp[pos : pos+2]
		if len(fp) >= pos+3 && fp[pos+2:pos+3] != ":" {
			return false
		}
		if _, err := strconv.ParseUint(hex, 16, 64); err != nil {
			return false
		}
	}

	return true
}
