package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Main(input []byte) string {
	chunked := splitBinaryInto5BitChunks(toBinary(input))

	result := []rune{}

	for _, c := range chunked {
		i, err := strconv.ParseInt(c, 2, 10)
		if err != nil {
			panic(err)
		}

		result = append(result, convert(int(i)))
	}

	return string(result)
}

func toBinary(input []byte) []rune {
	return []rune(string(input))
}

func splitBinaryInto5BitChunks(binary []rune) []string {
	binaryStr := ""
	for _, b := range binary {
		binaryStr += fmt.Sprintf("%08b", b)
	}

	chunked := splitToN([]rune(binaryStr), 5, true, "0")
	last := chunked[len(chunked)-1]
	if len(last) < 5 {
		chunked[len(chunked)-1] = string(last) + strings.Repeat("0", 5-len(last))
	}

	return chunked
}

func splitToN(runes []rune, length int, isPadding bool, paddingStr string) []string {
	result := []string{}

	for i := 0; i < len(runes); i += length {
		if i+length < len(runes) {
			result = append(result, string(runes[i:i+length]))
			continue
		}

		result = append(result, string(runes[i:]))
	}

	return result
}

const S = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

func convert(b int) rune {
	return rune(S[b])
}
