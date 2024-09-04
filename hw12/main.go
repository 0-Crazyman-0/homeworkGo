package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// 1.	Создание и вывод map
	// Описание: Создайте map для хранения строк и их длин, добавьте несколько элементов и выведите содержимое.
	//1
	strilen1 := make(map[string]int)
	strings := []string{"sonse", "luna", "nebosvod"}
	for _, s := range strings {
		strilen1[s] = len(s)
	}
	for str, length := range strilen1 {
		fmt.Printf("Строка: %s, Длина: %d\n", str, length)
	}
	//2
	strilen2 := map[string]int{
		"more":   4,
		"zemlia": 6,
		"kosmos": 6,
	}
	prover1 := []string{"more", "slon"}
	for _, key := range prover1 {
		if proverkey(strilen2, key) {
			fmt.Printf("Ключ '%s' существует в map.\n", key)
		} else {
			fmt.Printf("Ключ '%s' не найден в map.\n", key)
		}
	}
	//3
	strilen3 := map[string]int{
		"krisa":    5,
		"loshad":   6,
		"krakadil": 7,
	}
	obevlen1(strilen3, "loshad", 7)
	for str, length := range strilen3 {
		fmt.Printf("Строка: %s, Длина: %d\n", str, length)
	}
	//4
	mMap1 := map[string]int{
		"odin": 1,
		"dwa":  2,
		"tri":  3,
	}
	fmt.Println("До удаления:", mMap1)
	udalit1(mMap1, "tri")
	fmt.Println("После удаления:", mMap1)
	//5
	text1 := "сдараствуйте сегоня мы идем рыбачить "
	fq1 := podsjotjas1(text1)
	fmt.Println(fq1)
	//6
	mMap2 := map[string]int{
		"telefon":  2,
		"kandisha": 3,
		"telok":    5,
	}
	sort1 := sorlirkey(mMap2)
	fmt.Println("Отсортированные ключи:", sort1)
	//7
	mMap3 := map[string]int{
		"raz": 1,
		"dva": 2,
	}
	emap := map[string]int{}
	fmt.Println("mMap3 пустой?", pustottest(mMap3))
	fmt.Println("emap пустой?", pustottest(emap))
	//8
	mMap4 := map[string]int{
		"raz": 1,
		"dva": 2,
		"tri": 3,
	}
	invMap1 := inverting1(mMap4)
	fmt.Println("Исходная map:", mMap4)
	fmt.Println("Инвертированная map:", invMap1)
	//9
	mMap5 := map[string]int{
		"a": 1,
		"b": 2,
		"v": 3,
		"g": 2,
	}
	fmt.Println(proverdubl(mMap5))
	//10
	mMap6 := map[string]int{
		"a": 1,
		"b": 2,
		"v": 3,
		"g": 4,
	}
	fmt.Println(poluchznagh(mMap6))
	//11
	slice := []string{"apple", "banana", "apple", "orange", "banana", "grape"}
	fmt.Println(podsunikznak(slice))
	//12
	mMap7 := map[string]int{
		"a": 1,
		"b": 2,
		"v": 3,
		"g": 4,
	}
	con1 := func(v int) bool {
		return v%2 == 0
	}
	fmt.Println(poiskznagh(mMap7, con1))
	//13
	mMap8 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	mMap9 := map[string]int{
		"b": 4,
		"c": 5,
		"d": 6,
	}
	merMap1 := obeddvamap(mMap8, mMap9)
	fmt.Println(merMap1)
	//14
	peopl1 := map[string]int{
		"Nek": 1,
		"bax": 2,
		"Sum": 3,
		"Ali": 4,
	}
	key1 := poiskpoznachen(peopl1, 4)
	if key1 != "" {
		fmt.Printf("Ключ для значения 4: %s\n", key1)
	} else {
		fmt.Println("Значение не найдено")
	}
	//15
	mMap10 := map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}

	mMap11 := map[string]int{
		"Tom": 1,
		"Bob": 3,
		"Sam": 4,
		"Eve": 8,
	}

	result := peresechen(mMap10, mMap11)
	fmt.Println("Пересечение двух map:", result)
}

// 2.	Проверка наличия ключа
// Описание: Создайте map с несколькими элементами и напишите функцию, которая проверяет наличие заданного ключа.

func proverkey(m map[string]int, key string) bool {
	_, exsist := m[key]
	return exsist
}

// 3.	Обновление значения по ключу
// Описание: Напишите функцию для обновления значения в map по заданному ключу.

func obevlen1(m map[string]int, key string, newval int) {
	m[key] = newval
}

// 4.	Удаление элемента из map
// Описание: Создайте функцию, которая удаляет элемент из map по заданному ключу.

func udalit1(m map[string]int, key string) {
	delete(m, key)
}

// 5.	Подсчет частоты слов
// Описание: Напишите функцию, которая подсчитывает частоту слов в строке и возвращает map с результатами.

func podsjotjas1(m string) map[string]int {
	word := strings.Fields(m)
	fmap := make(map[string]int)
	for _, w := range word {
		fmap[w]++
	}
	return fmap
}

// 6.	Сортировка ключей в map
// Описание: Напишите функцию, которая сортирует ключи в map и возвращает отсортированные ключи.

func sorlirkey(m map[string]int) []string {
	key := make([]string, 0, len(m))
	for k := range m {
		key = append(key, k)
	}
	sort.Strings(key)
	return key
}

// 7.	Проверка пустого map
// Описание: Напишите функцию, которая проверяет, пустой ли map.

func pustottest(m map[string]int) bool {
	return len(m) == 0
}

// 8.	Инвертирование ключей и значений
// Описание: Напишите функцию, которая инвертирует ключи и значения в map.

func inverting1(m map[string]int) map[int]string {
	inver := make(map[int]string)
	for k, v := range m {
		inver[v] = k
	}
	return inver
}

// 9.	Проверка дубликатов в map
// Описание: Напишите функцию, которая проверяет, есть ли дубликаты значений в map.

func proverdubl(m map[string]int) bool {
	val := make(map[int]int)
	for _, v := range m {
		val[v]++
		if val[v] > 1 {
			return true
		}
	}
	return false
}

// 10.	Получение всех значений
// Описание: Напишите функцию, которая возвращает все значения из map в срезе.

func poluchznagh(m map[string]int) []int {
	val := make([]int, 0, len(m))
	for _, v := range m {
		val = append(val, v)
	}
	return val
}

// 11.	Подсчет уникальных значений в срезе строк
// Описание: Напишите функцию, которая подсчитывает уникальные значения в срезе строк с использованием map.

func podsunikznak(m []string) int {
	a := make(map[string]bool)
	for _, str := range m {
		a[str] = true
	}
	return len(a)
}

// 12.	Поиск значений по маске
// Описание: Напишите функцию, которая возвращает все ключи, чьи значения удовлетворяют заданному условию.

func poiskznagh(m map[string]int, a func(int) bool) []string {
	key := []string{}
	for k, v := range m {
		if a(v) {
			key = append(key, k)
		}
	}
	return key
}

// 13.	Объединение двух map
// Описание: Напишите функцию для объединения двух map. Если ключи совпадают, значения из второго map должны заменять значения из первого.

func obeddvamap(m1, m2 map[string]int) map[string]int {
	mMap := make(map[string]int)
	for k, v := range m1 {
		mMap[k] = v
	}
	for k, v := range m2 {
		mMap[k] = v
	}
	return mMap
}

// 14.	Поиск по значению и получение ключа
// Описание: Напишите функцию, которая возвращает ключ по значению в map. Если значение не найдено, верните пустую строку.

func poiskpoznachen(m map[string]int, value int) string {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return ""
}

// 15.	Пересечение двух map
// Описание: Напишите функцию для нахождения пересечения двух map на основе ключей и значений.

func peresechen(m1, m2 map[string]int) map[string]int {
	i := make(map[string]int)
	for k, v := range m1 {
		if val, ok := m2[k]; ok && val == v {
			i[k] = v
		}
	}
	return i
}

// 16.	Создание среза ключей из map
// Описание: Напишите функцию, которая возвращает срез всех ключей из map.
// 17.	Упаковка значений map в строки
// Описание: Напишите функцию, которая упаковывает все значения map в одну строку, разделенную запятой.
// 18.	Копирование map
// Описание: Напишите функцию, которая создает копию map.
// 19.	Преобразование map в срез пар ключ-значение
// Описание: Напишите функцию, которая преобразует map в срез пар ключ-значение.
// 20.	Обновление значений map на основе другой функции
// Описание: Напишите функцию, которая обновляет значения в map, применяя функцию к каждому значению.
// 21.	Кэширование с использованием map и sync.RWMutex
// Описание: Реализуйте простой кэш с использованием map и sync.RWMutex.
// 22.	Сложный фильтр по нескольким условиям
// Описание: Напишите функцию, которая фильтрует map по нескольким условиям.
// 23.	Слияние и агрегация значений из нескольких map
// Описание: Реализуйте функцию, которая сливает несколько map в один, суммируя значения для одинаковых ключей.
// 24.	Итерирование по ключам и значениям с использованием функции обратного вызова
// Описание: Напишите функцию, которая итерирует по map, вызывая функцию обратного вызова для каждого ключа и значения.
// 25.	Создание и использование карты с указателями на структуры
// Описание: Создайте map с указателями на структуры в качестве значений.
// 26.	Сопоставление map и массивов
// Описание: Напишите функцию, которая преобразует map в массив структур и наоборот.
// 27.	Реализация метода для структуры с map
// Описание: Реализуйте метод для структуры, который работает с map в этой структуре.
// 28.	Использование map для реализации простого хранилища данных
// Описание: Реализуйте простое хранилище данных с map для управления записями.
// 29.	Реализация кэширования с лимитом на количество элементов
// Описание: Реализуйте кэш с использованием map и ограничением на количество элементов.
// 30.	Реализация map для хранения срезов строк
// Описание: Реализуйте map, где ключи – это строки, а значения – срезы строк. Добавьте функции для добавления, удаления и получения срезов.
