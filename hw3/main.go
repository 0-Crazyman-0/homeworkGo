package main

import "strings"

func main() {
	PrintGreeting()

	DisplaySeparator()

	NotifyStart()

	println(GetWelcomeMessage())

	println(GetPiValue())

	println(IsServerActive())

	PrintNumber(26)

	GreetUser("Nikita")

	ToggleLight(false)

	println(add(5, 10))

	println(Concat("Hellow ", "My freand"))

	println(IsEven(50))

	adder := func(a, b int) int {
		c := a + b
		return c
	}
	println(adder)

	concatenator := func(a, b string) string {
		c := a + b
		return c
	}
	println(concatenator("Hi", "Haw are you"))

	isPositive := func(a int) bool {
		if a > 0 {
			return true
		} else {
			return false
		}
	}
	println(isPositive(50))

	println(Calculate(10, 20, add))

	Execute(false, exe2)

	println(Apply(25, aly))

	c := Multiplier(5)
	println(c(50))

	p := StringRepeater("Hello", 5)
	println(p())

	r := ConditionalPrinter(true)
	r()
}

func PrintGreeting() {
	println("Hello, Word")
}
func DisplaySeparator() {
	println("Helo my name is Jonson")
}
func NotifyStart() {
	println("Program started")
}
func GetWelcomeMessage() string {
	return ("Welcome")
}
func GetPiValue() float32 {
	return 3.14
}
func IsServerActive() bool {
	return true
}
func PrintNumber(a int) {
	println(a)
}
func GreetUser(username string) {
	println("Добро пожавловать", username)
}
func ToggleLight(tl bool) {
	if tl {
		println("Light on")
	} else {
		println("Light of")
	}
}
func add(a, b int) int {
	c := a + b
	return c
}
func Concat(a, b string) string {
	return a + b
}
func IsEven(ie int) bool {
	if ie%2 == 0 {
		return true
	} else {
		return false
	}
}
func Calculate(a, b int, add func(a, b int) int) int {
	d := add(a, b)
	return d
}
func Execute(a bool, c func(b bool)) {
	c(a)
}
func exe2(b bool) {
	println(b)
}
func Apply(a int, c func(b int) int) int {
	return c(a)
}
func aly(b int) int {
	return b
}
func Multiplier(factor int) func(a int) int {
	return func(a int) int {
		return a * factor
	}
}

func StringRepeater(a string, n int) func() string {
	return func() string {
		return strings.Repeat(a, n)
	}
}
func ConditionalPrinter(condition bool) func() {
	return func() {
		if condition {
			println("принята")
		}

	}
}
