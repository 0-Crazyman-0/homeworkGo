package main

func main() {

	a := "Hello"

	d := "Hello"

	x := 5

	b := 10

	z := 25

	v := 40
	PrintGreeting()

	PrintWeather()

	PrintTrafficLight()

	GetGrade()

	GetDiscount()

	GetTemperatureDescription()

	CheckNumber(5)

	CheckAge(20)

	CheckPassword("1234")

	println(add(x, b))

	println(CompareStrings(a, d))

	println(Max(x, b))

	operation := func(z, v int) int {
		c := z - v
		if c < 0 {
			c = -c
		}
		return c

	}
	println(operation(z, v))

	concat := func(a, b string) string {
		c := a + b
		return c
	}
	println(concat("hello \n", "My freand"))

	multiply := func(a, b int) int {
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		return a + b
	}
	println(multiply(25, -40))

	println(ApplyOperation(20, 50, appo))

	CheckCondition(5, chco)

	FormatAndPrint("helo ", foap)

	q := CreateMultiplier(10)
	println(q(10))
	n := CreateGreeter("Nekruz")
	println(n("Nekruz"))
	c := CreateValidator("1234")
	println(c("1234"))
}

func PrintGreeting() {
	dayType := "утро"
	switch dayType {
	case "утро":
		println("Дорое утро")
	case "день":
		println("Добрый день")
	case "вечер":
		println("Жобрый вечер")
	default:
		println("спокойной ночи ")
	}
}

func PrintWeather() {
	weatherType := "облачно"
	switch weatherType {
	case "солнечно":
		println("Погода солнечная")
	case "облачно":
		println("погода облачная")
	case "дождливо":
		println("погода дождливая")
	default:
		println("погода заснеженная ")
	}
}

func PrintTrafficLight() {
	lightColor := "жолтый"
	switch lightColor {
	case "крастный":
		println("стоп")
	case "жолтый":
		println("внимание ")
	case "зеленый":
		println("идите")
	default:
		println("в данный моментр светофор не работает ")
	}
}

func GetGrade() int {
	score := 75
	if score >= 90 {
		println("ваша оценка А")
	} else if score >= 75 {
		println("ваша оценка В")
	} else if score >= 50 {
		println("вашша оценка С")
	} else {
		println("ваша оценка F")
	}
	return score
}

func GetDiscount() int {
	amount := 750 //$
	if amount > 1000 {
		a := amount * 10 / 100
		println("ваша скидка составляет \n", a, "$")
	} else if amount > 500 && amount <= 1000 {
		a := amount * 5 / 100
		println("ваша скидка составляет \n", a, "$")
	} else if amount <= 500 {
		a := amount * 0 / 100
		println("ваша скидка составляет \n", a, "$")
	} else {
		println("у вашей покупки нет скидок")
	}
	return amount
}

func GetTemperatureDescription() int {
	temperature := 40
	if temperature < 10 {
		println("Холодно")
	} else if temperature >= 10 && temperature <= 25 {
		println("Тепло")
	} else {
		println("Жарко")
	}
	return temperature
}

func CheckNumber(a int) {
	if a > 0 {
		println("Положительное")
	} else if a < 0 {
		println("Отрицательное")
	} else {
		println("Ноль")
	}
}

func CheckAge(a uint) {
	if a >= 18 {
		println("Совершеннолетний")
	} else {
		println("Несовершеннолетний")
	}
}

func CheckPassword(Pas string) {
	if Pas == "1234" {
		println("Пароль верный")
	} else {
		println("Пароль неверный")
	}
}

func add(x, b int) int {
	if x < 0 {
		x = -x
	}
	if b < 0 {
		b = -b
	}
	c := x + b
	return c
}

func CompareStrings(a, d string) string {
	if a == d {
		return "равны"
	} else {
		return "не равны"
	}
}

func Max(x, b int) int {
	if x > b {
		return x
	} else if x < b {
		return b
	} else {
		return 0
	}
}

func ApplyOperation(a, b int, c func(a, b int) int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return c(a, b)
}
func appo(a, b int) int {
	c := a + b
	return c
}

func CheckCondition(a int, b func() bool) {
	if a >= 0 {
		println("условие выполнено")
	} else {
		println("уловие не выполнено")
	}
}
func chco() bool {
	return true
}

func FormatAndPrint(a string, b func() string) {
	a = "helo my freand"
	if a == "" {
		println("Пустая строка")
	} else {
		println(a)
	}
}
func foap() string {
	return "helo"
}

func CreateMultiplier(a int) func(b int) int {
	return func(a int) int {
		if a < 0 {
			return a * -a
		} else {
			return a * a
		}

	}
}

func CreateGreeter(a string) func(b string) string {
	return func(b string) string {
		b = "Добро пожаловать "
		if a == "" {
			return b
		} else {
			return b + a
		}
	}
}

func CreateValidator(a string) func(string) bool {
	return func(s string) bool {
		if s == a {
			return true
		} else {
			return false
		}

	}
}
