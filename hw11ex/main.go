package main

import (
	"fmt"
	"strings"
)

func main() {
	//1
	sp := StringProcessor{}
	stringsSlice := []string{"hello world", "golang is awesome", "callbacks are useful"}
	upperCaseStrings := sp.Process(stringsSlice, ConvertToUpper)
	fmt.Println("Верхний регистр:", upperCaseStrings)
	//2
	formatter := Formatter{}
	str := "example"
	withPrefix := formatter.Format(str, AddPrefix)
	fmt.Println("С префиксом:", withPrefix)
	//3
	tp := TextProcessor{}
	text1 := "Go is a statically typed, compiled programming language designed at Google"
	wordCount := tp.Process(text1, WordCount)
	fmt.Printf("Количество слов: %d\n", wordCount)
	//4
	sm := StringManipulator{}

	text2 := "Hello World Go Programming"

	invert := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	replaceSpaces := func(s string) string {
		return strings.ReplaceAll(s, " ", "-")
	}

	// Пример 1: Инверсия строки
	invertedText := sm.Manipulate(text2, invert)
	fmt.Printf("Инверсия строки: %s\n", invertedText)

	// Пример 2: Замена пробелов на тире
	replacedText := sm.Manipulate(text2, replaceSpaces)
	fmt.Printf("Замена пробелов на тире: %s\n", replacedText)
	//5
	sf := StringFilter{}

	stringsSlice1 := []string{"apple", "banana", "cherry", "date", "fig", "grape"}

	// Пример 1: Фильтрация строк, содержащих подстроку "ap"
	filteredBySubstring := sf.Filter(stringsSlice1, ContainsSubstring("ap"))
	fmt.Println("Строки, содержащие 'ap':", filteredBySubstring)

	// Пример 2: Фильтрация строк длиной не менее 5 символов
	filteredByLength := sf.Filter(stringsSlice, HasMinLength(5))
	fmt.Println("Строки длиной не менее 5 символов:", filteredByLength)
}

// Задача 1: Обработка строк с использованием методов и callback
// Создайте структуру StringProcessor, которая содержит метод для обработки строки с использованием функции callback. Функция callback должна применяться к каждой строке. Реализуйте несколько функций callback для преобразования строки, таких как конвертация в верхний регистр и удаление пробелов.
type StringProcessor struct{}

func (sp *StringProcessor) Process(stringsSlice []string, callback func(string) string) []string {
	var result []string
	for _, str := range stringsSlice {
		result = append(result, callback(str))
	}
	return result
}

func ConvertToUpper(s string) string {
	return strings.ToUpper(s)
}

func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// Задача 2: Форматирование строки с использованием callback
// Создайте структуру Formatter, которая имеет метод для форматирования строки. Используйте callback для определения способа форматирования строки. Реализуйте различные способы форматирования, такие как добавление префикса и суффикса.
type Formatter struct{}

func (f *Formatter) Format(s string, callback func(string) string) string {
	return callback(s)
}
func AddPrefix(s string) string {
	return "Prefix_" + s
}

func AddSuffix(s string) string {
	return s + "_Suffix"
}

func AddPrefixAndSuffix(s string) string {
	return "Prefix_" + s + "_Suffix"
}

// Задача 3: Обработка текста с использованием методов и callback
// Создайте структуру TextProcessor, которая содержит метод для обработки текста с использованием функции callback. Функция callback должна возвращать количество слов в строке. Реализуйте функцию callback для подсчета количества слов и символов.
type TextProcessor struct{}

func (tp *TextProcessor) Process(text string, callback func(string) int) int {
	return callback(text)
}
func WordCount(s string) int {
	words := strings.Fields(s)
	return len(words)
}
func CharacterCount(s string) int {
	return len(s)
}

// Задача 4: Изменение строки с использованием методов и closure
// Создайте структуру StringManipulator, которая содержит метод для изменения строки с использованием closure. Реализуйте несколько закрытых функций для преобразования строки, таких как инверсия строки и замена пробелов на тире.

type StringManipulator struct{}

func (sm *StringManipulator) Manipulate(s string, closure func(string) string) string {
	return closure(s)
}

// Задача 5: Обработка строки с использованием методов и callback для фильтрации
// Создайте структуру StringFilter, которая имеет метод для фильтрации строки с использованием функции callback. Функция callback должна проверять, удовлетворяет ли строка определенному условию. Реализуйте функции callback для проверки наличия подстроки и длины строки

type StringFilter struct{}

func (sf *StringFilter) Filter(stringsSlice []string, callback func(string) bool) []string {
	var result []string
	for _, str := range stringsSlice {
		if callback(str) {
			result = append(result, str)
		}
	}
	return result
}
func ContainsSubstring(substring string) func(string) bool {
	return func(s string) bool {
		return strings.Contains(s, substring)
	}
}
func HasMinLength(minLength int) func(string) bool {
	return func(s string) bool {
		return len(s) >= minLength
	}
}

// Задача 6: Изменение строки с использованием методов и callback
// Создайте структуру TextEditor, которая содержит метод для изменения текста с использованием функции callback. Реализуйте несколько функций callback для изменения текста, таких как добавление текста в начало и конец строки.

// Задача 7: Обработка строки с использованием методов и closure для статистики
// Создайте структуру Statistics, которая содержит метод для подсчета статистики строки. Используйте closure для подсчета различных метрик, таких как количество пробелов и количество уникальных символов.

// Задача 8: Замена строки с использованием методов и callback
// Создайте структуру Replacer, которая содержит метод для замены строки с использованием функции callback. Реализуйте функции callback для замены всех вхождений подстроки и удаления всех цифр.

// Задача 9: Форматирование строки с использованием методов и callback
// Создайте структуру TextFormatter, которая содержит метод для форматирования строки с использованием callback. Реализуйте разные функции форматирования, такие как подчеркивание и замена пробелов на подчеркивания.

// Задача 10: Анализ строки с использованием методов и callback
// Создайте структуру StringAnalyzer, которая содержит метод для анализа строки с использованием функции callback. Реализуйте функции callback для подсчета количества гласных и согласных букв.
