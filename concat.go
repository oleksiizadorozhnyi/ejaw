package optimization

import (
	"strings"
)

func Concat(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s
	}
	return result
}

func ConcatOptimized(strs []string) string {
	var builder strings.Builder

	totalLen := 0
	for _, s := range strs {
		totalLen += len(s)
	}
	builder.Grow(totalLen)

	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}
