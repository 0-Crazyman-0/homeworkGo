package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("exemple.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024)
	totalCount := 0
	for {
		n, err := file.Read(buf)
		if err != nil && err.Error() != "EOF" {
			return 0, err
		}
		totalCount += n
		if n == 0 {
			break
		}
	}
	fmt.Println(countCharacters(file.Name()))
}
func countCharacters(fileName string) (int, error) {
	c, err := io.ReadAll()
	if err != nil {
		return 0, err
	}
	return len(co)
}

// 1.	Подсчет символов в файле
// •	Описание: Напишите программу, которая считывает файл и подсчитывает количество символов в нём.
// •	Методы или функции:
// •	func countCharacters(fileName string) (int, error)
// 2.	Подсчет строк в файле
// •	Описание: Напишите программу, которая считает количество строк в текстовом файле.
// •	Методы или функции:
// •	func countLines(fileName string) (int, error)
// 3.	Счетчик слов в строке
// •	Описание: Напишите функцию, которая считает количество слов в строке.
// •	Методы или функции:
// •	func countWords(s string) int
// 4.	Запись строки в файл
// •	Описание: Напишите программу, которая записывает заданную строку в файл.
// •	Методы или функции:
// •	func writeStringToFile(fileName, content string) error
// 5.	Чтение первой строки файла
// •	Описание: Напишите программу, которая читает первую строку из текстового файла.
// •	Методы или функции:
// •	func readFirstLine(fileName string) (string, error)
// 6.	Копирование содержимого одного файла в другой
// •	Описание: Напишите программу, которая копирует содержимое одного файла в другой.
// •	Методы или функции:
// •	func copyFile(src, dst string) error
// 7.	Чтение строки с консоли и запись в файл
// •	Описание: Напишите программу, которая читает строку с консоли и записывает её в файл.
// •	Методы или функции:
// •	func readAndWriteToFile(fileName string) error
// 8.	Обратное чтение файла
// •	Описание: Напишите программу, которая читает файл с конца в начало и выводит его содержимое на экран.
// •	Методы или функции:
// •	func reverseReadFile(fileName string) (string, error)
// 9.	Объединение содержимого двух файлов
// •	Описание: Напишите программу, которая объединяет содержимое двух файлов и записывает результат в новый файл.
// •	Методы или функции:
// •	func concatenateFiles(file1, file2, outputFile string) error
// 10.	Проверка существования файла
// •	Описание: Напишите функцию, которая проверяет, существует ли файл с заданным именем.
// •	Методы или функции:
// •	func fileExists(fileName string) bool
// 11.	Чтение и подсчет уникальных слов в файле
// •	Описание: Напишите программу, которая читает файл и подсчитывает количество уникальных слов.
// •	Методы или функции:
// •	func countUniqueWords(fileName string) (int, error)
// 12.	Поиск и замена слова в файле
// •	Описание: Напишите программу, которая заменяет все вхождения определенного слова в файле на другое слово.
// •	Методы или функции:
// •	func replaceWordInFile(fileName, oldWord, newWord string) error
// 13.	Подсчет частоты слов в файле
// •	Описание: Напишите программу, которая подсчитывает частоту каждого слова в файле.
// •	Методы или функции:
// •	func wordFrequency(fileName string) (map[string]int, error)
// 14.	Переворачивание строк в файле
// •	Описание: Напишите программу, которая переворачивает строки в файле и записывает их в новый файл.
// •	Методы или функции:
// •	func reverseLines(fileName, outputFile string) error
// 15.	Удаление пустых строк из файла
// •	Описание: Напишите программу, которая удаляет все пустые строки из файла.
// •	Методы или функции:
// •	func removeEmptyLines(fileName, outputFile string) error
// 16.	Сравнение двух файлов на идентичность
// •	Описание: Напишите программу, которая сравнивает два файла и определяет, идентичны ли они.
// •	Методы или функции:
// •	func compareFiles(file1, file2 string) (bool, error)
// 17.	Чтение файла с конца
// •	Описание: Напишите программу, которая читает файл с конца и выводит его содержимое на экран.
// •	Методы или функции:
// •	func readFileReverse(fileName string) error
// 18.	Подсчет количества строк с определенным словом
// •	Описание: Напишите программу, которая считает количество строк в файле, содержащих определенное слово.
// •	Методы или функции:
// •	func countLinesWithWord(fileName, word string) (int, error)
// 19.	Генерация файла с повторяющимися строками
// •	Описание: Напишите программу, которая генерирует файл, содержащий заданную строку, повторенную указанное количество раз.
// •	Методы или функции:
// •	func generateRepeatedLinesFile(fileName, line string, count int) error
// 20.	Подсчет количества символов в каждой строке файла
// •	Описание: Напишите программу, которая подсчитывает количество символов в каждой строке файла и выводит их на экран.
// •	Методы или функции:
// •	func countCharactersPerLine(fileName string) error
// 21.	Поиск самого длинного слова в файле
// •	Описание: Напишите программу, которая находит самое длинное слово в файле и выводит его на экран.
// •	Методы или функции:
// •	func findLongestWord(fileName string) (string, error)
// 22.	Подсчет частоты встречаемости символов в файле
// •	Описание: Напишите программу, которая подсчитывает частоту встречаемости каждого символа в файле.
// •	Методы или функции:
// •	func charFrequency(fileName string) (map[rune]int, error)
// 23.	Реверсирование слов в каждой строке файла
// •	Описание: Напишите программу, которая реверсирует слова в каждой строке файла и записывает результат в новый файл.
// •	Методы или функции:
// •	func reverseWordsInFile(fileName, outputFile string) error
// 24.	Подсчет количества слов в каждой строке файла
// •	Описание: Напишите программу, которая подсчитывает количество слов в каждой строке файла и выводит результат на экран.
// •	Методы или функции:
// •	func countWordsPerLine(fileName string) error
// 25.	Найти и заменить слово в файле, с сохранением регистра
// •	Описание: Напишите программу, которая находит и заменяет все вхождения слова в файле, сохраняя регистр (с учетом заглавных и строчных букв).
// •	Методы или функции:
// •	func replaceWordWithCase(fileName, oldWord, newWord string) error
// 26.	Поиск самого короткого слова в файле
// •	Описание: Напишите программу, которая находит самое короткое слово в файле и выводит его на экран.
// •	Методы или функции:
// •	func findShortestWord(fileName string) (string, error)
// 27.	Чтение и запись файла по частям
// •	Описание: Напишите программу, которая читает файл по частям (например, по 1024 байта) и записывает его содержимое в другой файл.
// •	Методы или функции:
// •	func copyFileInChunks(srcFileName, dstFileName string) error
// 28.	Подсчет количества символов, слов и строк в файле
// •	Описание: Напишите программу, которая подсчитывает количество символов, слов и строк в файле и выводит результат на экран.
// •	Методы или функции:
// •	func countFileStats(fileName string) (int, int, int, error)
// 29.	Поиск и удаление строки в файле
// •	Описание: Напишите программу, которая находит и удаляет строку с определенным содержимым из файла.
// •	Методы или функции:
// •	func deleteLineFromFile(fileName, lineToDelete string) error
// 30.	Сортировка строк в файле в алфавитном порядке
// •	Описание: Напишите программу, которая сортирует строки в файле в алфавитном порядке и записывает результат в новый файл.
// •	Методы или функции:
// •	func sortFileLines(fileName, outputFile string) error
