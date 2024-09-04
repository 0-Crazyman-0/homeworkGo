package main

import (
	"fmt"
	"strings"
)

func main() {
	//1
	str1 := MyString{Text: "Hello, my brather "}
	fmt.Println("Length of string:", str1.Length())
	fmt.Println("Word count:", str1.WordCount())
	//2
	for1 := MyFormatter{Text: "Hello,Brather"}
	fmt.Println("Original string:", for1.Text)
	fmt.Println("Upper case:", for1.ToUpper())
	fmt.Println("Lower case:", for1.ToLower())
	//3
	num1 := 10
	intP1 := IntPointer{Value: &num1}
	fmt.Println("Initial value:", *intP1.Value)
	intP1.Increment()
	fmt.Println("After Increment:", *intP1.Value)
	intP1.Decrement()
	fmt.Println("After Decrement:", *intP1.Value)
	//4
	clen1 := TextCleaner{Text: "Hello,brather go to run with me"}
	fmt.Println("Original string:", clen1.Text)
	fmt.Println("After TrimSpaces:", clen1.TrimSpaces())
	fmt.Println("After RemoveSpaces:", clen1.RemoveSpaces())
	//5
	joi1 := StringJoiner{Strings: []string{"Hello", "My", "Brodher", "you", "are", "Ronaldo"}}
	fmt.Println("Concatenated string:", joi1.Concat())
	fmt.Println("Joined with '-':", joi1.Join("-"))
	//6
	extr1 := TextExtractor{Text: "Hello,Ronaldo"}
	fmt.Println("Original string:", extr1.Text)
	fmt.Println("Substr(7, 5):", extr1.Substr(7, 5))
	fmt.Println("Suffix(6):", extr1.Suffix(6))
	//7
	rev1 := TextReverser{Text: "Hello, World How are you"}
	fmt.Println("Original string:", rev1.Text)
	fmt.Println("Reversed string:", rev1.Reverse())
	fmt.Println("Reversed words:", rev1.ReverseWords())
	//8
	inser1 := TextInserter{Text: "How are you Ronaldo "}
	fmt.Println("Original string:", inser1.Text)
	fmt.Println("InsertAt(7, ' Gamer '):", inser1.InsertAt(7, " Gamer "))
	fmt.Println("InsertAfterWord(' Ronaldo', ', Gemer'):", inser1.InsertAfterWord("Ronaldo", ", Gamer"))

}

// 1.  Длина строки и Количество слов
// •	Описание: Реализуйте интерфейс StringProcessor, который будет содержать методы Length() и WordCount(). Реализуйте структуру MyString, которая будет работать с текстом и реализует этот интерфейс.
// •	Методы:
// •	Length() int для получения длины строки.
// •	WordCount() int для подсчета количества слов

type StringProcessor interface {
	Length() int
	WordCount() int
}
type MyString struct {
	Text string
}

func (ms MyString) Length() int {
	return len(ms.Text)
}
func (ms MyString) WordCount() int {
	words := strings.Fields(ms.Text)
	return len(words)
}

// 2. Форматирование строки
// •	Описание: Реализуйте интерфейс Formatter с методами ToUpper() и ToLower(). Реализуйте структуру MyFormatter, которая будет работать со строкой и реализует этот интерфейс.
// •	Методы:
// •	ToUpper() string для преобразования строки в верхний регистр.
// •	ToLower() string для преобразования строки в нижний регистр.

type Formatter interface {
	ToUpper() string
	ToLower() string
}
type MyFormatter struct {
	Text string
}

func (mf MyFormatter) ToUpper() string {
	return strings.ToUpper(mf.Text)
}
func (mf MyFormatter) ToLower() string {
	return strings.ToLower(mf.Text)
}

// 3. Работа с указателями на числа
// •	Описание: Реализуйте интерфейс PointerOperations с методами Increment() и Decrement(). Реализуйте структуру IntPointer с указателем на число, которая реализует этот интерфейс.
// •	Методы:
// •	Increment() для увеличения значения числа на 1.
// •	Decrement() для уменьшения значения числа на 1.

type PointerOperations interface {
	Increment()
	Decrement()
}
type IntPointer struct {
	Value *int
}

func (ip *IntPointer) Increment() {
	*ip.Value++
}
func (ip *IntPointer) Decrement() {
	*ip.Value--
}

// 4. Удаление пробелов в строке
// •	Описание: Реализуйте интерфейс StringCleaner с методами TrimSpaces() и RemoveSpaces(). Реализуйте структуру TextCleaner, которая будет работать со строками и реализует этот интерфейс.
// •	Методы:
// •	TrimSpaces() string для удаления пробелов с начала и конца строки.
// •	RemoveSpaces() string для удаления всех пробелов из строки.

type StringCleaner interface {
	TrimSpaces() string
	RemoveSpaces() string
}
type TextCleaner struct {
	Text string
}

func (tc TextCleaner) TrimSpaces() string {
	return strings.TrimSpace(tc.Text)
}
func (tc TextCleaner) RemoveSpaces() string {
	return strings.ReplaceAll(tc.Text, " ", "")
}

// 5. Конкатенация строк
// •	Описание: Реализуйте интерфейс StringConcatenator с методами Concat() и Join(). Реализуйте структуру StringJoiner, которая будет работать с массивами строк и реализует этот интерфейс.
// •	Методы:
// •	Concat() string для конкатенации всех строк в массиве.
// •	Join(separator string) string для объединения всех строк с заданным разделителем.

type StringConcatenator interface {
	Concat() string
	Join(separator string) string
}
type StringJoiner struct {
	Strings []string
}

func (sj StringJoiner) Concat() string {
	return strings.Join(sj.Strings, "")
}
func (sj StringJoiner) Join(separator string) string {
	return strings.Join(sj.Strings, separator)
}

// 6. Извлечение подстроки
// •	Описание: Реализуйте интерфейс SubstringExtractor с методами Substr(start, length int) string и Suffix(length int) string. Реализуйте структуру TextExtractor, которая работает со строками и реализует этот интерфейс.
// •	Методы:
// •	Substr(start, length int) string для извлечения подстроки заданной длины с указанной позиции.
// •	Suffix(length int) string для извлечения последних n символов строки.

type SubstringExtractor interface {
	Substr(start, length int) string
	Suffix(length int) string
}
type TextExtractor struct {
	Text string
}

func (te TextExtractor) Substr(start, length int) string {
	if start < 0 || start >= len(te.Text) || length <= 0 {
		return ""
	}

	end := start + length
	if end > len(te.Text) {
		end = len(te.Text)
	}

	return te.Text[start:end]
}
func (te TextExtractor) Suffix(length int) string {
	if length <= 0 || length > len(te.Text) {
		return te.Text
	}

	return te.Text[len(te.Text)-length:]
}

// 7. Переворот строки
// •	Описание: Реализуйте интерфейс StringReverser с методами Reverse() string и ReverseWords() string. Реализуйте структуру TextReverser, которая будет работать со строками и реализует этот интерфейс.
// •	Методы:
// •	Reverse() string для переворота строки.
// •	ReverseWords() string для переворота слов в строке.

type StringReverser interface {
	Reverse() string
	ReverseWords() string
}
type TextReverser struct {
	Text string
}

func (tr TextReverser) Reverse() string {
	runes := []rune(tr.Text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func (tr TextReverser) ReverseWords() string {
	words := strings.Fields(tr.Text)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// 8. Вставка подстроки
// •	Описание: Реализуйте интерфейс StringInserter с методами InsertAt(index int, substring string) string и InsertAfterWord(word, substring string) string. Реализуйте структуру TextInserter, которая работает со строками и реализует этот интерфейс.
// •	Методы:
// •	InsertAt(index int, substring string) string для вставки подстроки в указанную позицию.
// •	InsertAfterWord(word, substring string) string для вставки подстроки после указанного слова.

type StringInserter interface {
	InsertAt(index int, substring string) string
	InsertAfterWord(word, substring string) string
}
type TextInserter struct {
	Text string
}

func (ti TextInserter) InsertAt(index int, substring string) string {
	if index < 0 || index > len(ti.Text) {
		return ti.Text
	}

	return ti.Text[:index] + substring + ti.Text[index:]
}
func (ti TextInserter) InsertAfterWord(word, substring string) string {
	wordIndex := strings.Index(ti.Text, word)
	if wordIndex == -1 {
		return ti.Text
	}

	insertPos := wordIndex + len(word)
	return ti.Text[:insertPos] + substring + ti.Text[insertPos:]
}

// 9. Проверка на палиндром
// •	Описание: Реализуйте интерфейс PalindromeChecker с методами IsPalindrome() bool и IsWordPalindrome() bool. Реализуйте структуру TextPalindromeChecker, которая будет работать со строками и реализует этот интерфейс.
// •	Методы:
// •	IsPalindrome() bool для проверки, является ли строка палиндромом.
// •	IsWordPalindrome() bool для проверки, является ли строка палиндромом по словам.

type PalindromeChecker interface {
	IsPalindrome() bool
	IsWordPalindrome() bool
}
type TextPalindromeChecker struct {
	Text string
}

// func (pr TextPalindromeChecker)IsPalindrome()bool{

// }

// 10. Удаление повторяющихся слов
// •	Описание: Реализуйте интерфейс DuplicateRemover с методами RemoveDuplicates() и RemoveDuplicatesCaseInsensitive(). Реализуйте структуру TextDuplicateRemover, которая работает со строками и реализует этот интерфейс.
// •	Методы:
// •	RemoveDuplicates() string для удаления дубликатов слов в строке.
// •	RemoveDuplicatesCaseInsensitive() string для удаления дубликатов слов без учета регистра.
// 11. Бюджетирование
// •	Описание: Реализуйте интерфейс Budget с методами AddIncome(amount float64) и AddExpense(amount float64). Реализуйте структуру MonthlyBudget, которая будет работать с месячным бюджетом и реализует этот интерфейс.
// •	Методы:
// •	AddIncome(amount float64) для добавления дохода.
// •	AddExpense(amount float64) для добавления расхода.
// •	GetBalance() float64 для получения текущего баланса.
// 12. Валютообменник
// •	Описание: Реализуйте интерфейс CurrencyConverter с методами ToUSD(amount float64) и ToEUR(amount float64). Реализуйте структуру Exchange, которая будет работать с конвертацией валют и реализует этот интерфейс.
// •	Методы:
// •	ToUSD(amount float64) float64 для конвертации в доллары США.
// •	ToEUR(amount float64) float64 для конвертации в евро.
// 13. Банковский счет
// •	Описание: Реализуйте интерфейс BankAccount с методами Deposit(amount float64) и Withdraw(amount float64). Реализуйте структуру Account, которая будет работать с банковским счетом и реализует этот интерфейс.
// •	Методы:
// •	Deposit(amount float64) для депозита средств на счет.
// •	Withdraw(amount float64) для снятия средств со счета.
// •	GetBalance() float64 для получения текущего баланса.
// 14. Налоговый расчет
// •	Описание: Реализуйте интерфейс TaxCalculator с методами CalculateIncomeTax(income float64) и CalculateVAT(amount float64). Реализуйте структуру SimpleTaxCalculator, которая будет работать с расчетом налогов и реализует этот интерфейс.
// •	Методы:
// •	CalculateIncomeTax(income float64) float64 для расчета налога на доход.
// •	CalculateVAT(amount float64) float64 для расчета НДС.
// 15. Ведение счета-фактуры
// •	Описание: Реализуйте интерфейс Invoice с методами AddItem(name string, price float64) и Total() float64. Реализуйте структуру SimpleInvoice, которая будет работать со счетами и реализует этот интерфейс.
// •	Методы:
// •	AddItem(name string, price float64) для добавления товара в счет.
// •	Total() float64 для получения общей суммы счета.
// 16. Управление портфелем акций
// •	Описание: Реализуйте интерфейс StockPortfolio с методами Buy(symbol string, quantity int) и Sell(symbol string, quantity int). Реализуйте структуру Portfolio, которая будет работать с управлением акциями и реализует этот интерфейс.
// •	Методы:
// •	Buy(symbol string, quantity int) для покупки акций.
// •	Sell(symbol string, quantity int) для продажи акций.
// •	GetHoldings() map[string]int для получения текущих владений.
// 17. Банкомат
// •	Описание: Реализуйте интерфейс ATM с методами Deposit(amount float64) и Withdraw(amount float64). Реализуйте структуру ATMSystem, которая будет работать с банкоматом и реализует этот интерфейс.
// •	Методы:
// •	Deposit(amount float64) для депозита средств в банкомат.
// •	Withdraw(amount float64) для снятия средств из банкомата.
// •	GetBalance() float64 для получения текущего баланса.
// 18. Управление зарплатой сотрудников
// •	Описание: Реализуйте интерфейс Payroll с методами AddEmployee(name string, salary float64) и GetSalary(name string) float64. Реализуйте структуру CompanyPayroll, которая будет работать с зарплатами сотрудников и реализует этот интерфейс.
// •	Методы:
// •	AddEmployee(name string, salary float64) для добавления сотрудника и его зарплаты.
// •	GetSalary(name string) float64 для получения зарплаты сотрудника.
// •	TotalPayroll() float64 для получения общей суммы зарплат.
// 19. Управление библиотекой
// •	Описание: Реализуйте интерфейс Library с методами AddBook(title, author string) и GetBooks() []string. Реализуйте структуру BookLibrary, которая будет работать с библиотекой книг и реализует этот интерфейс.
// •	Методы:
// •	AddBook(title, author string) для добавления книги в библиотеку.
// •	GetBooks() []string для получения списка книг в библиотеке.
// 20. Управление заказами в интернет-магазине
// •	Описание: Реализуйте интерфейс OrderManager с методами AddOrder(orderID string, amount float64) и GetTotalSales() float64. Реализуйте структуру OnlineStore, которая будет работать с заказами и реализует этот интерфейс.
// •	Методы:
// •	AddOrder(orderID string, amount float64) для добавления заказа.
// •	GetTotalSales() float64 для получения общей суммы продаж.
// 21. Управление банковскими транзакциями
// •	Описание: Реализуйте интерфейс TransactionManager с методами AddTransaction(amount float64) и GetStatement() []string. Реализуйте структуру BankAccount, которая будет работать с банковскими транзакциями и реализует этот интерфейс.
// •	Методы:
// •	AddTransaction(amount float64) для добавления транзакции.
// •	GetStatement() []string для получения выписки по счету.
// 22. Управление каталогом продуктов
// •	Описание: Реализуйте интерфейс ProductCatalog с методами AddProduct(name string, price float64) и GetProductPrice(name string) float64. Реализуйте структуру Catalog, которая будет работать с каталогом продуктов и реализует этот интерфейс.
// •	Методы:
// •	AddProduct(name string, price float64) для добавления продукта в каталог.
// •	GetProductPrice(name string) float64 для получения цены продукта.
// •	GetProducts() map[string]float64 для получения всех продуктов и их цен.
// 23. Управление заказами в ресторане
// •	Описание: Реализуйте интерфейс RestaurantOrderManager с методами AddOrder(orderID string, items []string) и GetOrderDetails(orderID string) []string. Реализуйте структуру Restaurant, которая будет работать с заказами и реализует этот интерфейс.
// •	Методы:
// •	AddOrder(orderID string, items []string) для добавления заказа.
// •	GetOrderDetails(orderID string) []string для получения деталей заказа.
// •	GetAllOrders() map[string][]string для получения всех заказов.
// 24. Расчет ипотечного кредита
// •	Описание: Реализуйте интерфейс MortgageCalculator с методами CalculateMonthlyPayment(principal, annualRate float64, years int) и CalculateTotalPayment(principal, annualRate float64, years int). Реализуйте структуру Mortgage, которая будет работать с расчетом ипотечного кредита и реализует этот интерфейс.
// •	Методы:
// •	CalculateMonthlyPayment(principal, annualRate float64, years int) float64 для расчета ежемесячного платежа.
// •	CalculateTotalPayment(principal, annualRate float64, years int) float64 для расчета общей суммы выплат.
// 25. Управление клиентами банка
// •	Описание: Реализуйте интерфейс CustomerManager с методами AddCustomer(name string, accountNumber string) и GetCustomerDetails(accountNumber string) (string, string). Реализуйте структуру Bank, которая будет работать с клиентами банка и реализует этот интерфейс.
// •	Методы:
// •	AddCustomer(name string, accountNumber string) для добавления клиента.
// •	GetCustomerDetails(accountNumber string) (string, string) для получения деталей клиента.
// •	GetAllCustomers() map[string]string для получения всех клиентов.
// 26. Учет рабочего времени сотрудников
// •	Описание: Реализуйте интерфейс TimeTracker с методами ClockIn(employeeID string) и ClockOut(employeeID string). Реализуйте структуру WorkTime, которая будет работать с учетом рабочего времени сотрудников и реализует этот интерфейс.
// •	Методы:
// •	ClockIn(employeeID string) для регистрации начала работы.
// •	ClockOut(employeeID string) для регистрации окончания работы.
// •	GetWorkingHours(employeeID string) float64 для получения отработанных часов.
// 27. Система учета студентов
// •	Описание: Реализуйте интерфейс StudentManager с методами AddStudent(name string, studentID string) и GetStudentDetails(studentID string) (string, string). Реализуйте структуру University, которая будет работать с учетной записью студентов и реализует этот интерфейс.
// •	Методы:
// •	AddStudent(name string, studentID string) для добавления студента.
// •	GetStudentDetails(studentID string) (string, string) для получения данных студента.
// •	GetAllStudents() map[string]string для получения списка всех студентов.
// 28. Управление расходами предприятия
// •	Описание: Реализуйте интерфейс ExpenseManager с методами AddExpense(category string, amount float64) и GetTotalExpenses(category string) float64. Реализуйте структуру Enterprise, которая будет работать с расходами предприятия и реализует этот интерфейс.
// •	Методы:
// •	AddExpense(category string, amount float64) для добавления расхода.
// •	GetTotalExpenses(category string) float64 для получения общей суммы расходов по категории.
// •	GetAllExpenses() map[string]float64 для получения всех расходов.
// 29. Управление учебными курсами
// •	Описание: Реализуйте интерфейс CourseManager с методами AddCourse(courseName string, credits int) и GetCourseCredits(courseName string) int. Реализуйте структуру EducationInstitution, которая будет работать с учебными курсами и реализует этот интерфейс.
// •	Методы:
// •	AddCourse(courseName string, credits int) для добавления курса.
// •	GetCourseCredits(courseName string) int для получения количества кредитов курса.
// •	GetAllCourses() map[string]int для получения всех курсов.
// 30. Учет счетов за коммунальные услуги
// •	Описание: Реализуйте интерфейс UtilityBillManager с методами AddBill(utilityType string, amount float64) и GetTotalBills(utilityType string) float64. Реализуйте структуру Household, которая будет работать с учетом коммунальных услуг и реализует этот интерфейс.
// •	Методы:
// •	AddBill(utilityType string, amount float64) для добавления счета за коммунальные услуги.
// •	GetTotalBills(utilityType string) float64 для получения общей суммы по типу коммунальной услуги.
// •	GetAllBills() map[string]float64 для получения всех счетов.
