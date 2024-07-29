package main

import (
	"fmt"
)

func main() {
	// var x int = 10
	// var ptr *int
	// ptr = &x
	// fmt.Println(x)
	// fmt.Println(&x)
	// fmt.Println(ptr)
	// fmt.Println(*ptr)

	//1
	// var x int = 5
	// updateValue(&x)
	// fmt.Println(x)

	//2
	// 	x, y := 10, 20
	// 	swap(&x, &y)
	// 	fmt.Println(x, y)

	// 3

	//4
	var math, physics, chemestry int
	mathPtr := &math
	physicsPtr := &physics
	chemestryPtr := &chemestry
	addGrade(mathPtr, 85)
	addGrade(physicsPtr, 90)
	addGrade(chemestryPtr, 78)
	fmt.Printf("средний бал по предсетам %.2f\n", printAverageGrade(mathPtr, physicsPtr, chemestryPtr))

}

// 1 напишите функцию updateValue  которая принимает указатель на целое число и увеличивает его значение на 10
func updateValue(a *int) {
	*a = *a + 10
}

// 2 напишите функцию swap которая меняет местами значения двух переменных с использованием переменных
func swap(x, y *int) {
	*x, *y = *y, *x
}

//3 Напишите программу для учета голосов на выборах. Реализуйте функции для подсчета голосов за каждого кандидата и определения победителя

func vote(candidarePtr *string, votePtr *int) {
	*votePtr++
}

func winner(cadidate1VotesPtr *int, cadidate2VotesPtr *int) string {
	if *cadidate1VotesPtr > *cadidate2VotesPtr {
		return "Кандидат 1"
	} else if *cadidate2VotesPtr > *cadidate1VotesPtr {
		return "Кандидат 2"
	} else {
		return "ничья"
	}
}

//4 Напишите программу для учета оценок по предметам студентов. Реализуйте функции для добавления оценки по предмету и вывода среднего балла.

func addGrade(mathPtr *int, grade int) {
	*mathPtr = grade
}
func printAverageGrade(mathPtr *int, physicsPtr *int, chemestryPtr *int) float64 {
	sum := *mathPtr + *physicsPtr + *chemestryPtr
	return float64(sum) / 3.0
}
