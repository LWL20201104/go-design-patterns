package singleton

import "testing"

func BenchmarkGetInstanceV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInstanceV3()
	}
}

func BenchmarkGetInstanceV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInstanceV4()
	}
}

func BenchmarkGetInstanceV5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInstanceV5()
	}
}

// go test bench=.
// BenchmarkGetInstanceV3-8   	132847944	         8.669 ns/op
// BenchmarkGetInstanceV4-8   	616986142	         1.945 ns/op
// BenchmarkGetInstanceV5-8   	622991500	         1.926 ns/op
