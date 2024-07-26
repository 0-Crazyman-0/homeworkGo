package main

import "fmt"

func main() {
	println(Temp(20))
	println(Agge(20))
	println(speedtest(110))
	println(Scr(20))
	println(bmii(16.5))
	var a UnaryOperation
	a = sqr
	println(a(5))
	var b UnaryOperation
	b = dabl
	println(b(10))
	var c Check
	c = chc
	println(c(4))
	var d Check
	d = pol
	println(d(1))
	var x Operation
	x[0] = Operpl
	x[1] = opermin
	x[2] = operum
	println(Operpl(12, 10), opermin(15, 20), operum(2, 5))
	obratno(60)
	println(Battery(50))
	println(weig(50))
	println(Grd(60))
	println(AldB(22.5, 115.0))
	product := Product{"katroshka", 99}
	opis(product)
	Hello := Person{"Nekruz", "Karimov", 19}
	hello(Hello)
	Product1 := Product{name: "kartoshka", price: 20}
	Product2 := Product{name: "savzi", price: 15}
	e := opis1(Product1, Product2)
	fmt.Println("Более дорогой продукт:", e.name, e.price)

}

// 1.	Проверка температуры
// Определите тип Temperature на основе float64. Напишите функцию, которая принимает температуру и возвращает сообщение о том, является ли она ниже, выше или равной нулю.
type Temperature float64

func Temp(a Temperature) string {
	if a < 0 {
		return "ниже нуля"
	} else {
		return "выше нуля или ноль"
	}
}

// 2.	Проверка возраста
// Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том, является ли человек подростком (возраст от 13 до 19 лет включительно) или нет.
type Age int

func Agge(a Age) string {
	if a >= 13 && a <= 19 {
		return "вы евляетесь подростком "
	} else if a > 19 {
		return "вы не  евляетесь подростком"
	}
	return ""
}

// 3.	Проверка скорости
// Определите тип Speed на основе float64. Напишите функцию, которая принимает скорость и возвращает сообщение о том, является ли она допустимой (от 0 до 120 км/ч включительно) или нет.
type speed float64

func speedtest(a speed) string {
	if a >= 0 && a <= 120 {
		return "скорость допустима "
	} else {
		return "скорость не  допустима"
	}
}

// 4.	Проверка счета
// Определите тип Score на основе int. Напишите функцию, которая принимает счет и возвращает сообщение о том, является ли он положительным, отрицательным или нулевым.
type Score int

func Scr(a Score) string {
	if a > 0 {
		return "является положительным"
	} else {
		return " является отрицательным или нулевым"
	}
}

// 5.	Классификация BMI
// Определите тип BMI на основе float64. Напишите функцию, которая принимает значение BMI и возвращает сообщение о том, в какой категории оно находится: "Underweight", "Normal weight", "Overweight", или "Obesity".
type BMI float64

func bmii(a BMI) string {
	if a < 18.5 {
		return "Underweight"
	} else if a >= 18.5 && a < 24.9 {
		return "Normal weight"
	} else if a >= 25 && a < 29.9 {
		return "Overweight"
	} else {
		return "Obesity"
	}
}

// 6.	Возведение в квадрат
// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функцию для возведения числа в квадрат и используйте тип UnaryOperation для вызова этой функции.
type UnaryOperation func(int) int

func sqr(a int) int {
	return a * a
}

// 7.	Удвоение числа
// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функцию для удвоения числа и используйте тип UnaryOperation для вызова этой функции.

func dabl(a int) int {
	return 2 * a
}

// 8.	Проверка четности
// Определите тип функции Check, которая принимает int и возвращает bool. Напишите функцию для проверки четности числа и используйте тип Check для вызова этой функции.
type Check func(int) bool

func chc(a int) bool {
	if a%2 == 1 {
		return false
	} else {
		return true
	}
}

// 9.	Проверка положительности
// Определите тип функции Check, которая принимает int и возвращает bool. Напишите функцию для проверки, является ли число положительным, и используйте тип Check для вызова этой функции.

func pol(a int) bool {
	if a >= 0 {
		return true
	} else {
		return false
	}
}

// 10.	Комбинирование функций
// Определите тип функции Operation, которая принимает два int и возвращает int. Напишите функции для сложения, вычитания и умножения. Напишите функцию, которая принимает два int и массив функций Operation, и возвращает массив результатов применения этих функций.
type Operation [3]func(int, int) int

func Operpl(a, b int) int {
	return a + b
}
func opermin(a, b int) int {
	return a - b
}
func operum(a, b int) int {
	return a * b
}

// 11.	Обратный отсчет
// Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает Count и выводит обратный отсчет до нуля.
type Count = int

func obratno(a Count) {
	for i := a; i >= 0; i-- {
		println(i)
	}
}

// 12.	Проверка уровня батареи
// Создайте псевдоним для типа int под названием BatteryLevel. Напишите функцию, которая принимает BatteryLevel и возвращает сообщение о том, низкий, средний или высокий уровень заряда.

type BatteryLevel = int

func Battery(a BatteryLevel) string {
	if a > 0 && a < 20 {
		return "низкий уровень заряд"
	} else if a >= 20 && a < 60 {
		return "средний уровень заряд"
	} else if a >= 60 && a <= 100 {
		return "высокий уровень заряда"
	}
	return ""
}

// 13.	Определение категории веса
// Создайте псевдоним для типа float64 под названием Weight. Напишите функцию, которая принимает Weight и возвращает сообщение о том, в какой категории веса находится объект: "Light", "Medium" или "Heavy".
type Weight float64

func weig(a Weight) string {
	if a < 50 {
		return "Light"
	} else if a >= 50 && a < 100 {
		return "Medium"
	} else {
		return "Heavy"
	}
}

// 14.	Оценка успеваемости
// Создайте псевдоним для типа int под названием Grade. Напишите функцию, которая принимает Grade и возвращает сообщение о том, является ли оценка удовлетворительной (50 и выше) или нет.
type Grade = int

func Grd(a Grade) string {
	if a >= 50 {
		return "оценка является удовлетворительной"
	} else {
		return "оценка не  является удовлетворительной"
	}
}

// 15.	Оценка состояния здоровья
// Создайте псевдонимы для типов float64 под названиями BMI и BloodPressure. Напишите функцию, которая принимает BMI и BloodPressure, и возвращает сообщение о состоянии здоровья: "Healthy", "At risk" или "Unhealthy".
type BloodPressure = float64

func AldB(a BMI, b BloodPressure) string {
	if a < 18.5 || a >= 30 || b < 90 || b > 140 {
		return "Unhealthy"
	} else if a >= 25 && a < 30 || b >= 120 && b <= 140 {
		return "At risk"
	} else {
		return "Unhealthy"
	}
}

// 16.	Описание продукта
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает Product и выводит его описание.
type Product struct {
	name  string
	price int
}

func opis(a Product) {
	fmt.Printf("Product: %s,Price:%d\n", a.name, a.price)
}

// 17.	Вывод информации о человеке
// Создайте структуру Person с полями FirstName, LastName и Age. Напишите функцию, которая принимает Person и выводит его полное имя и возраст.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func hello(a Person) {
	fmt.Println(a.FirstName, a.LastName, a.Age)
}

// 18.	Сравнение продуктов
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает два Product и возвращает более дорогой продукт.

func opis1(a, b Product) Product {
	if a.price > b.price {
		return a
	} else {
		return b
	}
}

// 19.	Управление инвентарем
// Создайте структуру Item с полями Name и Quantity. Напишите функцию, которая принимает массив Item и возвращает суммарное количество всех предметов.

type Item struct {
	name     string
	Quantity int
}

// 20.	Управление библиотекой
// Создайте структуру Library с полями Name и Books (массив структур Book, где Book имеет поля Title и Author). Напишите функцию, которая принимает Library и выводит информацию о библиотеке и всех книгах в ней.

type Library struct {
	name string
	book [2]string
}

func librr(a Library) {
	println("", a.name("Порно"))
}

// 21.	Обновление цены продукта
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает указатель на Product и обновляет его цену.

// 22.	Увеличение возраста
// Создайте структуру Person с полями Name и Age. Напишите функцию, которая принимает указатель на Person и увеличивает его возраст на 1.

// 23.	Обновление информации о продукте
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает указатель на Product и обновляет его название и цену.

// 24.	Увеличение количества предметов
// Создайте структуру Item с полями Name и Quantity. Напишите функцию, которая принимает указатель на Item и увеличивает количество на заданное значение.

// 25.	Управление списком задач
// Создайте структуру Task с полями Title и Completed. Напишите функцию, которая принимает указатель на массив Task и обновляет статус выполнения задачи по заданному индексу.

// 26.	Адрес человека
// Создайте структуру Address с полями Street и City. Создайте структуру Person с полями Name и Address. Напишите функцию, которая принимает Person и выводит его имя и адрес.

// 27.	Описание компании
// Создайте структуру Company с полями Name и Location (структура Address). Напишите функцию, которая принимает Company и выводит информацию о компании.

// 28.	Описание курса
// Создайте структуру Course с полями Title и Instructor (структура Person). Напишите функцию, которая принимает Course и выводит информацию о курсе.

// 29.	Описание события
// Создайте структуру Event с полями Title и Location (структура Address). Напишите функцию, которая принимает Event и выводит информацию о событии.

// 30.	Управление проектом
// Создайте структуру Project с полями Name и Manager (структура Person, где Person имеет поле Name и Address (структура Address)). Напишите функцию, которая принимает Project и выводит информацию о проекте и менеджере.

// 31.	Простой узел
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел. Напишите функцию, которая создает два связанных узла и выводит их значения.

// 32.	Три связанных узла
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел. Напишите функцию, которая создает три связанных узла и выводит их значения.

// 33.	Обход связного списка
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел. Напишите функцию, которая принимает указатель на первый узел и обходит связный список, выводя значения всех узлов.

// 34.	Добавление узла в конец списка
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел. Напишите функцию, которая принимает указатель на первый узел и добавляет новый узел в конец списка.

// 35.	Удаление узла из списка
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел. Напишите функцию, которая принимает указатель на первый узел и значение, которое нужно удалить, и удаляет узел с этим значением из списка.

// 36.	Анонимный тип продукта
// Создайте переменную анонимного типа, которая содержит Name и Price для продукта. Напишите функцию, которая принимает эту переменную и выводит информацию о продукте.

// 37.	Анонимный тип человека
// Создайте переменную анонимного типа, которая содержит FirstName, LastName и Age для человека. Напишите функцию, которая принимает эту переменную и выводит полное имя и возраст человека.

// 38.	Анонимный тип автомобиля
// Создайте переменную анонимного типа, которая содержит Make, Model и Year для автомобиля. Напишите функцию, которая принимает эту переменную и выводит информацию о автомобиле.

// 39.	Анонимный тип книги
// Создайте переменную анонимного типа, которая содержит Title и Author для книги. Напишите функцию, которая принимает эту переменную и выводит информацию о книге.

// 40.	Анонимный тип события
// Создайте переменную анонимного типа, которая содержит Title, Date и Location (анонимная структура с полями Street и City) для события. Напишите функцию, которая принимает эту переменную и выводит информацию о событии.
