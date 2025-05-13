package main
import "fmt"

func main() {
	const NMAX int = 100
	var A [NMAX]int
	var n int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	used := make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				hundreds := A[i]
				tens := A[j]
				ones := A[k]
				number := hundreds*100 + tens*10 + ones
				if hundreds != 0 && number%2 == 0 {
					used[number] = true
				}
			}
		}
	}

	for num := 100; num <= 998; num += 2 {
		if used[num] {
			fmt.Print(num, " ")
		}
	}
	fmt.Println()
}
