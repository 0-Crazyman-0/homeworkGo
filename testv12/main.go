package main

import (
	"fmt"
	"strings"
)

func main() {
	//1
	a := []string{"elefent", "baba", "hello", "my brother"}
	c := srez(a)
	fmt.Println(c)
	//2
	b := []int{-10, 0, 5, -3, 8, -2, 4}
	d := neotr(b)
	fmt.Println(d)
	//3
	movies := []Movie{
		{"Inception", 8.8},
		{"The Dark Knight", 9.0},
		{"Interstellar", 8.6},
		{"The Prestige", 8.5},
	}
	e := hello(movies)
	fmt.Printf("Фильм с наивысшим рейтингом: %s, рейтинг: %.1f\n", e.Title, e.Rating)
}

// •	Реализуйте функцию, которая принимает срез строк и возвращает срез, содержащий только те строки, которые содержат букву e.
func srez(s []string) []string {
	var a []string
	for _, b := range s {
		if strings.Contains(b, "e") {
			a = append(a, b)
		}
	}
	return a
}

// •	Напишите функцию, которая принимает срез целых чисел и возвращает новый срез, содержащий только те числа, которые являются неотрицательными.
func neotr(a []int) []int {
	var b []int
	for _, i := range a {
		if i >= 0 {
			b = append(b, i)
		}
	}
	return b
}

// •	Реализуйте функцию, которая принимает срез структур Movie и возвращает структуру Movie, у которой рейтинг наивысший.
type Movie struct {
	Title  string
	Rating float64
}

func hello(a []Movie) Movie {
	if len(a) == 0 {
		return Movie{}
	}
	b := a[0]
	for _, movie := range a[1:] {
		if movie.Rating > b.Rating {
			b = movie
		}
	}
	return b
}
