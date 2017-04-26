package q

// START1 OMIT
// BenchmarkExpensive shows how to execute some expensive setup function
// and then reset the performance counters afterward; the goal is to not
// count the expensive setup in the execution time of the benchmark.
func BenchmarkExpensive(b *testing.B) {
	boringAndExpensiveSetup()
	b.ResetTimer() // HL
	for n := 0; n < b.N; n++ {
		// function under test
	}
}
// END1 OMIT

// START2 OMIT
// BenchmarkComplicated shows how to remove from the timings an
// expensive setup operation that must be performed for each iteration
// of the test.
func BenchmarkComplicated(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer() // HL
		complicatedSetup()
		b.StartTimer() // HL
		// function to test goes here
	}
}
// END2 OMIT
