package pkg

import "testing"

func BenchmarkNewGeneratedString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewGeneratedString()
		if err != nil {
			panic(err)
		}
	}
}
