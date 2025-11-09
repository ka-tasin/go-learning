func sum_n_numbers(n int) int {
	

	/* ---> not optimized - O(n)

	 var sum int = 0
	 for i := 0; i <= n; i++ {
	 	sum = sum + i
	 }
	return sum

	*/

	// optimized - O(1)
	return n * (n+1) / 2
}