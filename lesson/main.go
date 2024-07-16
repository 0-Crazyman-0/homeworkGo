package main

func StringRepeater(n int) func(d string) {
	return func(d string) {
		for i := 0; i < n; i++ {
			println(d)
		}
	}
}

func main() {
	f := StringRepeater(5)
	f("Hello")
}
