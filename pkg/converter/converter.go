package converter

import (
	"math"
	"strings"
)

const base = 62

func buildEncodingMaps() (map[int]rune, map[rune]int) {
	charRanges := []struct {
		begin rune
		end   rune
	}{
		{'a', 'z'},
		{'A', 'Z'},
		{'0', '9'},
	}
	encodeMap := make(map[int]rune, base)
	decodeMap := make(map[rune]int, base)
	ind := 0

	for _, charRange := range charRanges {
		for i := charRange.begin; i <= charRange.end; i++ {
			encodeMap[ind] = i
			decodeMap[i] = ind
			ind += 1
		}
	}

	return encodeMap, decodeMap
}

var encodeMap, decodeMap = buildEncodingMaps()

func ConvertToURL(id int) string {
	digits := []int{}

	for id > 0 {
		digits = append(digits, id%base)
		id = id / base
	}

	var result strings.Builder

	for i := len(digits) - 1; i >= 0; i-- {
		result.WriteRune(encodeMap[digits[i]])
	}

	return result.String()
}

func ConvertToID(url string) int {
	power := len(url) - 1
	id := 0

	for _, char := range url {
		digit := decodeMap[char]
		id += digit * int(math.Pow(float64(base), float64(power)))
		power -= 1
	}

	return id
}
