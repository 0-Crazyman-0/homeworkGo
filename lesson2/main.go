package main

import "fmt"

func main() {
	num := 15
	if num > 9 && num < 100 {
		println("двухзначина")
	}
	n := 6
	if n%2 == 1 {
		println("число нечотное")
	} else {
		println("число чотное")
	}

	m := 78
	if m >= 0 && m < 10 {
		println("число однозначное")
	} else if m >= 10 && m < 100 {
		println("число двухзначное")
	} else if m >= 100 && m < 1000 {
		println("сичло трехзначное")
	} else {
		println("число слишком болӣшое ")
	}

	a := 25
	b := 35
	if a > b {
		println("наибольшыйм евляется ", a)
	} else if a < b {
		println("наибольшыйм евляется ", b)
	} else {
		println("чило равны друг другу")
	}
	n1, n2, n3 := 12, 46, 25
	if n1 <= n2 && n2 <= n3 {
		fmt.Println(n3, n2, n1)
	} else if n1 <= n3 && n3 <= n2 {
		fmt.Println(n2, n3, n1)
	} else if n2 <= n1 && n1 <= n3 {
		fmt.Println(n3, n1, n2)
	} else if n2 <= n3 && n3 <= n1 {
		fmt.Println(n1, n3, n2)
	} else if n3 <= n1 && n1 <= n2 {
		fmt.Println(n2, n1, n3)
	} else {
		fmt.Println(n1, n2, n3)
	}

	nn := 1
	switch nn {
	case 1:
		fmt.Println("Positiv")
		fallthrough
	case -1:
		fmt.Println("negative")
	case 0:
		fmt.Println("Unsigmet")

	}

	aa := 19
	ab := 26
	ac := 22
	if aa <= ab && ab <= ac {
		fmt.Println(ab + ac)
	} else if aa <= ac && ac <= ab {
		fmt.Println(ac + ab)
	} else if ab <= aa && aa <= ac {
		fmt.Println(aa + ac)
	} else if ab <= ac && ac <= aa {
		fmt.Println(ac + aa)
	} else if ac <= aa && aa <= ab {
		fmt.Println(aa + ab)
	} else {
		fmt.Println(ab + ac)
	}

	ba := 1926

	if ba%4 == 0 && ba%100 == 0 && ba%400 != 0 {
		println("невыcакостный")
	} else if ba%4 == 0 {
		println("высакостный ")
	} else {
		println("невычакостный")
	}

	var ca int
	fmt.Scan(&ca)
	switch ca {
	case 1:
		println("январь имеет 31")
	case 2:
		println("февраль имеет 28")
	case 3:
		println("март имеет 31")
	case 4:
		println("апрель имеет 30")
	case 5:
		println("май имеет 31")
	case 6:
		println("июнь имеет 30")
	case 7:
		println("июль имеет 31")
	case 8:
		println("август имеет 31")
	case 9:
		println("сентябр имеет 30")
	case 10:
		println("октябр имеет 31")
	case 11:
		println("ноябр имеет 30")
	case 12:
		println("декабр имеет 31")
	default:
		println("введите число от 1 до 12")
	}

}
