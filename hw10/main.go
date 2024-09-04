package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//1
	book1 := Book{Title: "1984", Author: "George Orwell"}
	book2 := Book{Title: "To Kill a Mockingbird", Author: "Harper Lee"}
	library := Library{}
	library.AddBook(book1)
	library.AddBook(book2)
	for _, book := range library.Books {
		fmt.Println(book.GetDetails())
	}
	//2
	student := Student{}
	student.AddGrade(60)
	student.AddGrade(70)
	student.AddGrade(80)
	student.AddGrade(90)
	average := student.AverageGrade()
	fmt.Printf("Средняя оценка студента: %.2f\n", average)
	//3
	circle := Circle{Radius: 5}
	area := circle.Area()
	circumference := circle.Circumference()
	fmt.Printf("Площадь круга: %.2f\n", area)
	fmt.Printf("Периметр круга: %.2f\n", circumference)
	//4
	item1 := Item{Name: "Laptop", Price: 1200.50}
	item2 := Item{Name: "Smartphone", Price: 800.00}
	item3 := Item{Name: "Headphones", Price: 150.75}
	inventory := Inventory{}
	inventory.AddItem(item1)
	inventory.AddItem(item2)
	inventory.AddItem(item3)
	totalValue := inventory.TotalValue()
	fmt.Printf("Общая стоимость товаров в инвентаре: $%.2f\n", totalValue)
	//5

	//6
	taskList := TaskList{}
	taskList.AddTask("Complete Go project")
	taskList.AddTask("Read a book")
	taskList.Tasks[0].Completed = true
	fmt.Printf("Completed tasks: %d\n", taskList.CompletedTasks())
	//7
	temp := Temperature{Celsius: 25}
	fmt.Printf("Fahrenheit: %.2f\n", temp.ToFahrenheit())
	fmt.Printf("Kelvin: %.2f\n", temp.ToKelvin())
	//8
	classroom := Classroom{}
	classroom.AddStudent("Alice", 20)
	classroom.AddStudent("Bob", 22)
	fmt.Printf("Average Age: %.2f\n", classroom.AverageAge())
	//9
	warehouse := Warehouse{}
	warehouse.AddProduct("Laptop", 10)
	warehouse.AddProduct("Smartphone", 20)
	fmt.Printf("Total Quantity: %d\n", warehouse.TotalQuantity())
	//10
	notebook := Notebook{}
	notebook.AddNote("Learn Go", []string{"programming", "go"})
	notebook.AddNote("Buy groceries", []string{"shopping"})
	notes := notebook.NotesWithTag("programming")
	fmt.Printf("Notes with tag 'programming': %d\n", len(notes))
	//11
	payroll := Payroll{}
	payroll.AddEmployee("John Doe", 50000)
	payroll.AddEmployee("Jane Smith", 60000)
	fmt.Printf("Average Salary: %.2f\n", payroll.AverageSalary())
	//12
	budget := Budget{Amount: 1000.00}
	if budget.Allocate(200) {
		fmt.Println("Amount allocated successfully.")
	} else {
		fmt.Println("Insufficient budget.")
	}
	fmt.Printf("Remaining budget: %.2f\n", budget.Remaining())
	//13
	orderManager := OrderManager{}
	orderManager.AddOrder(1, 150.00)
	orderManager.AddOrder(2, 250.00)
	fmt.Printf("Total Orders Amount: %.2f\n", orderManager.TotalOrdersAmount())
	//14
	account := Account{Balance: 1000.00}
	account.Deposit(200.00)
	fmt.Printf("Current Balance after deposit: %.2f\n", account.GetBalance())
	if account.Withdraw(500.00) {
		fmt.Println("Withdrawal successful.")
	} else {
		fmt.Println("Insufficient balance.")
	}
	fmt.Printf("Current Balance: %.2f\n", account.GetBalance())
	//15
	text := Text{Content: "Hello, world!"}
	text.AddText(" How are you today?")
	fmt.Println("Updated Content:", text.Content)
	text.ReplaceWord("world", "Go")
	fmt.Println("After Replacement:", text.Content)
	fmt.Println("Length of Content:", text.Length())
}

// 1.	Книга и Автор
// •	Описание: Реализуйте структуру Book с полями title и author, а также метод GetDetails, который возвращает строку с деталями книги. Реализуйте структуру Library с массивом книг и метод AddBook, добавляющий книгу в библиотеку.
// •	Методы:
// •	GetDetails() string для структуры Book
// •	AddBook(book Book) для структуры Library

type Book struct {
	Title  string
	Author string
}

func (b Book) GetDetails() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.Title, b.Author)
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

// 2.	Оценки студента
// •	Описание: Реализуйте структуру Student с полем grades (срез оценок). Реализуйте метод AddGrade, добавляющий оценку, и метод AverageGrade, возвращающий среднее значение оценок.
// •	Методы:
// •	AddGrade(grade int)
// •	AverageGrade() float64

type Student struct {
	Grades []int
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}

func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0.0
	}

	total := 0
	for _, grade := range s.Grades {
		total += grade
	}
	return float64(total) / float64(len(s.Grades))
}

// 3.	Круг и Площадь
// •	Описание: Реализуйте структуру Circle с полем radius. Реализуйте методы Area и Circumference для вычисления площади и периметра круга.
// •	Методы:
// •	Area() float64
// •	Circumference() float64

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// 4.	Контейнер для товаров
// •	Описание: Реализуйте структуру Item с полями name и price. Реализуйте структуру Inventory с срезом товаров и методы для добавления товара и получения общего значения товаров в инвентаре.
// •	Методы:
// •	AddItem(item Item)
// •	TotalValue() float64

type Item struct {
	Name  string
	Price float64
}

type Inventory struct {
	Items []Item
}

func (inv *Inventory) AddItem(item Item) {
	inv.Items = append(inv.Items, item)
}

func (inv Inventory) TotalValue() float64 {
	total := 0.0
	for _, item := range inv.Items {
		total += item.Price
	}
	return total
}

// 5.	Обработка транзакций
// •	Описание: Реализуйте структуру Transaction с полями amount и description. Реализуйте метод Process, который выводит информацию о транзакции.
// •	Методы:
// •	Process()
// •	AddTransaction(amount float64, description string)

// 6.	Управление задачами
// •	Описание: Реализуйте структуру Task с полями title и completed. Реализуйте структуру TaskList с срезом задач и методы для добавления задачи и получения количества завершённых задач.
// •	Методы:
// •	AddTask(title string)
// •	CompletedTasks() int

type Task struct {
	Title     string
	Completed bool
}

type TaskList struct {
	Tasks []Task
}

func (tl *TaskList) AddTask(title string) {
	tl.Tasks = append(tl.Tasks, Task{Title: title, Completed: false})
}

func (tl TaskList) CompletedTasks() int {
	count := 0
	for _, task := range tl.Tasks {
		if task.Completed {
			count++
		}
	}
	return count
}

// 7.	Работа с температурой
// •	Описание: Реализуйте структуру Temperature с полем celsius. Реализуйте методы для преобразования температуры в Фаренгейт и Кельвин.
// •	Методы:
// •	ToFahrenheit() float64
// •	ToKelvin() float64

type Temperature struct {
	Celsius float64
}

func (t Temperature) ToFahrenheit() float64 {
	return t.Celsius*9/5 + 32
}

func (t Temperature) ToKelvin() float64 {
	return t.Celsius + 273.15
}

// 8.	Управление списком студентов
// •	Описание: Реализуйте структуру Student с полем name и age. Реализуйте структуру Classroom с срезом студентов и методы для добавления студента и получения средней возрастной группы.
// •	Методы:
// •	AddStudent(name string, age int)
// •	AverageAge() float64

type Student1 struct {
	Name string
	Age  int
}
type Classroom struct {
	Students []Student1
}

func (c *Classroom) AddStudent(name string, age int) {
	c.Students = append(c.Students, Student1{Name: name, Age: age})
}
func (c Classroom) AverageAge() float64 {
	totalAge := 0
	for _, student := range c.Students {
		totalAge += student.Age
	}
	return float64(totalAge) / float64(len(c.Students))

}

// 9.	Склад товаров
// •	Описание: Реализуйте структуру Product с полями name и quantity. Реализуйте структуру Warehouse с срезом продуктов и методы для добавления продукта и получения общего количества товаров на складе.
// •	Методы:
// •	AddProduct(name string, quantity int)
// •	TotalQuantity() int

type Product struct {
	Name     string
	Quantity int
}
type Warehouse struct {
	Products []Product
}

func (w *Warehouse) AddProduct(name string, quantity int) {
	w.Products = append(w.Products, Product{Name: name, Quantity: quantity})
}
func (w Warehouse) TotalQuantity() int {
	total := 0
	for _, product := range w.Products {
		total += product.Quantity
	}
	return total
}

// 10.	Заметки и метки
// •	Описание: Реализуйте структуру Note с полем content и tags. Реализуйте структуру Notebook с срезом заметок и методы для добавления заметки и получения всех заметок с указанной меткой.
// •	Методы:
// •	AddNote(content string, tags []string)
// •	NotesWithTag(tag string) []Note

type Note struct {
	Content string
	Tags    []string
}
type Notebook struct {
	Notes []Note
}

func (nb *Notebook) AddNote(content string, tags []string) {
	nb.Notes = append(nb.Notes, Note{Content: content, Tags: tags})
}
func (nb Notebook) NotesWithTag(tag string) []Note {
	var notesWithTag []Note
	for _, note := range nb.Notes {
		for _, noteTag := range note.Tags {
			if noteTag == tag {
				notesWithTag = append(notesWithTag, note)
				break
			}
		}
	}
	return notesWithTag
}

// 11.	Управление зарплатами
// •	Описание: Реализуйте структуру Employee с полями name и salary. Реализуйте структуру Payroll с срезом сотрудников и методы для добавления сотрудника и получения средней зарплаты.
// •	Методы:
// •	AddEmployee(name string, salary float64)
// •	AverageSalary() float64

type Employee struct {
	Name   string
	Salary float64
}
type Payroll struct {
	Employees []Employee
}

func (p *Payroll) AddEmployee(name string, salary float64) {
	p.Employees = append(p.Employees, Employee{Name: name, Salary: salary})
}
func (p Payroll) AverageSalary() float64 {
	totalSalary := 0.0
	for _, employee := range p.Employees {
		totalSalary += employee.Salary
	}
	return totalSalary / float64(len(p.Employees))
}

// 12.	Разделение бюджета
// •	Описание: Реализуйте структуру Budget с полем amount. Реализуйте метод Allocate, который выделяет сумму из бюджета, если сумма доступна. Реализуйте метод Remaining для получения оставшегося бюджета.
// •	Методы:
// •	Allocate(amount float64) bool
// •	Remaining() float64

type Budget struct {
	Amount float64
}

func (b *Budget) Allocate(amount float64) bool {
	if amount <= b.Amount {
		b.Amount -= amount
		return true
	}
	return false
}
func (b Budget) Remaining() float64 {
	return b.Amount
}

// 13.	Отслеживание заказов
// •	Описание: Реализуйте структуру Order с полями id и totalAmount. Реализуйте структуру OrderManager с срезом заказов и методы для добавления заказа и получения общей суммы всех заказов.
// •	Методы:
// •	AddOrder(id int, totalAmount float64)
// •	TotalOrdersAmount() float64

type Order struct {
	ID          int
	TotalAmount float64
}
type OrderManager struct {
	Orders []Order
}

func (om *OrderManager) AddOrder(id int, totalAmount float64) {
	om.Orders = append(om.Orders, Order{ID: id, TotalAmount: totalAmount})
}
func (om OrderManager) TotalOrdersAmount() float64 {
	total := 0.0
	for _, order := range om.Orders {
		total += order.TotalAmount
	}
	return total
}

// 14.	Управление аккаунтами
// •	Описание: Реализуйте структуру Account с полем balance. Реализуйте методы Deposit и Withdraw, а также метод GetBalance для получения текущего баланса.
// •	Методы:
// •	Deposit(amount float64)
// •	Withdraw(amount float64) bool
// •	GetBalance() float64

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}
func (a *Account) Withdraw(amount float64) bool {
	if amount <= a.Balance {
		a.Balance -= amount
		return true
	}
	return false
}
func (a Account) GetBalance() float64 {
	return a.Balance
}

// 15.	Операции над строками
// •	Описание: Реализуйте структуру Text с полем content. Реализуйте методы для добавления текста, замены слова и получения длины текста.
// •	Методы:
// •	AddText(text string)
// •	ReplaceWord(oldWord, newWord string)
// •	Length() int

type Text struct {
	Content string
}

func (t *Text) AddText(text string) {
	t.Content += text
}
func (t *Text) ReplaceWord(oldWord, newWord string) {
	t.Content = strings.ReplaceAll(t.Content, oldWord, newWord)
}
func (t Text) Length() int {
	return len(t.Content)
}

// 16.	Управление списком покупок
// •	Описание: Реализуйте структуру ShoppingItem с полями name и price. Реализуйте структуру ShoppingList с срезом покупок и методы для добавления покупки и получения общего значения покупок.
// •	Методы:
// •	AddItem(name string, price float64)
// •	TotalPrice() float64
// 17.	Учет времени
// •	Описание: Реализуйте структуру Event с полями name и date (в формате строк). Реализуйте методы для добавления события и получения всех событий после указанной даты.
// •	Методы:
// •	AddEvent(name string, date string)
// •	EventsAfterDate(date string) []Event
// 18.	Управление заказами в магазине
// •	Описание: Реализуйте структуру Order с полями orderID, product и quantity. Реализуйте структуру Store с срезом заказов и методы для добавления заказа и получения общего количества товаров по каждому продукту.
// •	Методы:
// •	AddOrder(orderID int, product string, quantity int)
// •	TotalQuantityByProduct(product string) int
// 19.	Расчет стоимости поездки
// •	Описание: Реализуйте структуру Trip с полями distance и costPerMile. Реализуйте методы для установки стоимости за милю и расчета общей стоимости поездки.
// •	Методы:
// •	SetCostPerMile(cost float64)
// •	TotalCost() float64
// 20.	Рейтинг фильмов
// •	Описание: Реализуйте структуру Movie с полями title и rating. Реализуйте структуру MovieList с срезом фильмов и методы для добавления фильма и получения среднего рейтинга.
// •	Методы:
// •	AddMovie(title string, rating float64)
// •	AverageRating() float64
// 21.	Система учёта заказов с учетом скидок
// •	Описание: Реализуйте структуру Order с полями id, product, quantity и price. Реализуйте структуру Discount с полем percentage. Реализуйте методы для применения скидки и получения общей суммы заказа с учетом скидки.
// •	Методы:
// •	ApplyDiscount(discount Discount)
// •	TotalAmount() float64
// 22.	Многократные переводы денег
// •	Описание: Реализуйте структуру Account с полями balance и history (срез транзакций). Реализуйте методы для пополнения, снятия и получения истории транзакций.
// •	Методы:
// •	Deposit(amount float64)
// •	Withdraw(amount float64) bool
// •	TransactionHistory() []string
// 23.	Статистика по продажам
// •	Описание: Реализуйте структуру Sale с полями item и amount. Реализуйте структуру SalesReport с методами для добавления продажи и получения статистики по общим продажам.
// •	Методы:
// •	AddSale(item string, amount float64)
// •	TotalSales() float64
// 24.	Учет кредитов и дебетов
// •	Описание: Реализуйте структуру Transaction с полями transactionType (кредит или дебет) и amount. Реализуйте структуру Account с методами для добавления транзакций и получения баланса.
// •	Методы:
// •	AddTransaction(transactionType string, amount float64)
// •	Balance() float64
// 25.	Учет заказов с учетом возвратов
// •	Описание: Реализуйте структуру Order с полями id, product, quantity и price. Реализуйте структуру OrderManager с методами для добавления заказа, обработки возвратов и получения общего значения активных заказов.
// •	Методы:
// •	AddOrder(id int, product string, quantity int, price float64)
// •	ReturnOrder(id int, quantity int)
// •	TotalActiveOrders() float64
// 26.	Управление проектами
// •	Описание: Реализуйте структуру Task с полями title, description и status. Реализуйте структуру Project с срезом задач и методы для добавления задачи и получения количества задач в каждом статусе.
// •	Методы:
// •	AddTask(title, description, status string)
// •	CountTasksByStatus(status string) int
// 27.	Управление пользовательскими данными
// •	Описание: Реализуйте структуру User с полями username, email и age. Реализуйте структуру UserManager с срезом пользователей и методы для добавления пользователя и получения всех пользователей старше определенного возраста.
// •	Методы:
// •	AddUser(username, email string, age int)
// •	UsersOlderThan(age int) []User
// 28.	Управление контрактами
// •	Описание: Реализуйте структуру Contract с полями contractID, client и amount. Реализуйте структуру ContractManager с срезом контрактов и методы для добавления контракта и получения общего объема контрактов для клиента.
// •	Методы:
// •	AddContract(contractID int, client string, amount float64)
// •	TotalAmountForClient(client string) float64
// 29.	Управление списком книг
// •	Описание: Реализуйте структуру Book с полями title, author и year. Реализуйте структуру Library с срезом книг и методы для добавления книги и получения всех книг авторов, опубликованных после определенного года.
// •	Методы:
// •	AddBook(title, author string, year int)
// •	BooksByAuthorsAfterYear(year int) []Book
// 30.	Отслеживание активностей пользователя
// •	Описание: Реализуйте структуру Activity с полями activityType и timestamp. Реализуйте структуру UserActivityTracker с срезом активностей и методы для добавления активности и получения всех активностей после определенного времени.
// •	Методы:
// •	AddActivity(activityType string, timestamp time.Time)
// •	ActivitiesAfterTime(timestamp time.Time) []Activity
