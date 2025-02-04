package optimization

import (
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	strs := make([]string, 30)
	for i := 0; i < 30; i++ {
		strs[i] = "тестове значення"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Concat(strs)
	}
}

func BenchmarkConcatOptimized(b *testing.B) {
	strs := make([]string, 30)
	for i := 0; i < 30; i++ {
		strs[i] = "тестове значення"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ConcatOptimized(strs)
	}
}
