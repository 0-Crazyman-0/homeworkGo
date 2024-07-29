package main

import "fmt"

func main() {
	//var i int
	//for i = 0; i <= 9; i++ {
	//	println(i)
	//}

	//println(i)

	//num := 546
	//for num > 0 {
	//	a := num % 10
	//	println(a)
	//	num = num / 10
	//}

	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		continue
	// 	}
	// 	println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	for j := 1; j < 10; j++ {
	// 		fmt.Printf("%d * %d = %d \n", i, j, i*j)
	// 	}
	// 	fmt.Printf("__________________")
	// }

	// n := 5
	// var s float64 = 0
	// for i := 1; i <= n; i++ {
	// 	s = s + 1/float64(i)
	// }
	// fmt.Printf("%.4f", s)
	// var p float64 = 1
	// var n float64
	// for i := 1; i < 10; i++ {
	// 	fmt.Scan(&n)
	// 	p = p * n
	// }
	// fmt.Println(p)

	var n int = 5
	var c float32
	for i := 0; i < n; i++ {
		fmt.Scan(&c)
		c = c - float32(int(c))
		fmt.Printf("%.2f", c)
	}

}
