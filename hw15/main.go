package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	//1
	person := Person{
		Name:  "Karimov Nekruz",
		Age:   20,
		Email: "NekruzRa@gmail.com",
	}
	jsonD, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error serializing to JSON:", err)
		return
	}
	fmt.Println(string(jsonD))
	//2
	jsonString1 := `{"name":"Nekruz Radjabov","age":20,"email":"NekruzRa@gmail.com"}`

	var person1 Person1

	err1 := json.Unmarshal([]byte(jsonString1), &person1)
	if err1 != nil {
		fmt.Println("Error deserializing JSON:", err1)
		return
	}
	fmt.Printf("Name: %s\n", person1.Name)
	fmt.Printf("Age: %d\n", person1.Age)
	fmt.Printf("Email: %s\n", person1.Email)
	//3
	data1 := map[string]int{
		"apples":  5,
		"oranges": 10,
	}
	jsonD1, err2 := json.Marshal(data1)
	if err2 != nil {
		fmt.Println("Error serializing map to JSON:", err2)
		return
	}
	fmt.Println(string(jsonD1))
	//4
	jsonString2 := `{
        "apples": 5,
        "oranges": 10
    }`
	data2 := make(map[string]int)
	err3 := json.Unmarshal([]byte(jsonString2), &data2)
	if err3 != nil {
		fmt.Println("Error deserializing JSON:", err3)
		return
	}
	for k, v := range data2 {
		fmt.Printf("%s: %d\n", k, v)
	}
	//5
	product := Product{
		Name:    "Hellokiti",
		Price:   10,
		InStock: true,
	}
	jsonD2, err4 := json.Marshal(product)
	if err4 != nil {
		fmt.Println("Error serializing structure to JSON:", err4)
		return
	}
	fmt.Println(string(jsonD2))
	//6
	person2 := Person2{
		Name:  "Nani",
		Age:   19,
		Email: "nanioumigot@gmail.com",
		Address: Address{
			Street: "33 Sino Street",
			City:   "Dushanbe",
		},
	}
	jsonD3, err5 := json.Marshal(person2)
	if err5 != nil {
		fmt.Println("Error serializing structure to JSON:", err5)
		return
	}
	fmt.Println("Вывод сериализованных данных")
	fmt.Println(string(jsonD3))

	var deserializedPerson Person2
	err5 = json.Unmarshal(jsonD3, &deserializedPerson)
	if err5 != nil {
		fmt.Println("Error deserializing JSON:", err5)
		return
	}
	fmt.Println("\nВывод десериализованных данных:")
	fmt.Printf("Name: %s\n", deserializedPerson.Name)
	fmt.Printf("Age: %d\n", deserializedPerson.Age)
	fmt.Printf("Email: %s\n", deserializedPerson.Email)
	fmt.Printf("Address: %s, %s\n", deserializedPerson.Address.Street, deserializedPerson.Address.City)
	//7
	data3 := map[string]interface{}{
		"name":  "John",
		"age":   30,
		"email": "john@example.com",
	}
	jsonD4, err6 := json.Marshal(data3)
	if err6 != nil {
		fmt.Println("Error serializing map to JSON:", err6)
		return
	}
	fmt.Println("Вывод сериализованных данных:")
	fmt.Println(string(jsonD4))
	var deserializedData map[string]interface{}
	err6 = json.Unmarshal(jsonD4, &deserializedData)
	if err6 != nil {
		fmt.Println("Error deserializing JSON:", err6)
		return
	}
	fmt.Println("\nВывод десериализованных данных:")
	for k, v := range deserializedData {
		fmt.Printf("%s: %v\n", k, v)
	}
	//8
	file, err7 := os.Open("person.json")
	if err7 != nil {
		fmt.Println("Error opening file:", err7)
		return
	}
	defer file.Close()
	data4, err8 := io.ReadAll(file)
	if err8 != nil {
		fmt.Println("Error reading file:", err8)
		return
	}
	var person3 Person3
	err8 = json.Unmarshal(data4, &person3)
	if err8 != nil {
		fmt.Println("Error deserializing JSON:", err8)
		return
	}
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Email: %s\n", person.Email)
	//9
	product1 := Product1{
		Name:  "Kartoshka",
		Price: 6.99,
	}
	jsonData, err9 := json.Marshal(product1)
	if err9 != nil {
		fmt.Println("Error serializing structure to JSON:", err9)
		return
	}
	file1, err10 := os.Create("output.json")
	if err10 != nil {
		fmt.Println("Error creating file:", err10)
		return
	}
	defer file1.Close()
	_, err10 = file1.Write(jsonData)
	if err10 != nil {
		fmt.Println("Error writing JSON to file:", err10)
		return
	}
	fmt.Println("JSON data successfully written to output.json")
}

// 1. Сериализация структуры в JSON
// Описание: Напишите программу, которая сериализует структуру Person в формат
// JSON и выводит результат на экран.
// Структура:
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// 2. Десериализация JSON в структуру
// Описание: Напишите программу, которая десериализует JSON-строку в структуру
// Person и выводит результат на экран.
// Структура:
type Person1 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// 3. Сериализация карты в JSON
// Описание: Напишите программу, которая сериализует карту map[string]int в JSON
// и выводит результат на экран.
// Пример данных:
//
//	data := map[string]int{
//	 "apples": 5,
//	 "oranges": 10,
//	}
//
// 4. Десериализация JSON в карту
// Описание: Напишите программу, которая десериализует JSON-строку в карту
// map[string]int и выводит результат на экран.
// Пример данных:
//
//	{
//	 "apples": 5,
//	 "oranges": 10
//	}
//
// 5. Пропуск полей при сериализации
// Описание: Напишите программу, которая сериализует структуру, пропуская поля
// с пустыми значениями.
// Структура:
type Product struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price,omitempty"`
	InStock bool    `json:"in_stock,omitempty"`
}

// 6. Работа с вложенными структурами
// Описание: Напишите программу, которая сериализует и десериализует структуру
// с вложенной структурой.
// Структура:
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}
type Person2 struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

// 7. Использование JSON с интерфейсами
// Описание: Напишите программу, которая сериализует и десериализует карту
// map[string]interface{} и выводит результат на экран.
// Пример данных:
//
//	data := map[string]interface{}{
//	 "name": "John",
//	 "age": 30,
//	 "email": "john@example.com",
//	}
//
// 8. Чтение JSON из файла
// Описание: Напишите программу, которая читает JSON из файла и десериализует
// его в структуру.
// Файл: person.json
// Структура:
type Person3 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// 9. Запись JSON в файл
// Описание: Напишите программу, которая сериализует структуру в JSON и
// записывает результат в файл.
// Файл: output.json
// Структура:
type Product1 struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// 10. Пропуск неизвестных полей при десериализации
// Описание: Напишите программу, которая десериализует JSON-строку в структуру,
// игнорируя неизвестные поля.
// Структура:
// type Person struct {
//  Name string `json:"name"`
//  Age int `json:"age"`
// }
// 11. Форматирование JSON-вывода
// Описание: Напишите программу, которая сериализует структуру в JSON с
// отступами для улучшения читаемости.
// Метод: Используйте json.MarshalIndent.
// Структура:
// type Person struct {
//  Name string `json:"name"`
//  Age int `json:"age"`
//  Email string `json:"email"`
// }
// 12. Обновление значения в JSON
// Описание: Напишите программу, которая десериализует JSON в структуру,
// изменяет одно из значений, а затем сериализует структуру обратно в JSON.
// Структура:
// type Person struct {
//  Name string `json:"name"`
//  Age int `json:"age"`
//  Email string `json:"email"`
// }
// 13. Парсинг JSON с разными типами данных
// Описание: Напишите программу, которая десериализует JSON-строку,
// содержащую данные разных типов (строки, числа, массивы, объекты), в
// map[string]interface{}.
// Пример данных:
// {
//  "name": "John",
//  "age": 30,
//  "scores": [100, 98, 95],
//  "address": {"city": "New York", "street": "5th Avenue"}
// }
// 14. Массив структур в JSON
// Описание: Напишите программу, которая сериализует массив структур Person в
// JSON и выводит результат на экран.
// Структура:
// type Person struct {
//  Name string `json:"name"`
//  Age int `json:"age"`
//  Email string `json:"email"`
// }
// 15. Добавление нового объекта в JSON-массив
// Описание: Напишите программу, которая читает JSON-массив из файла,
// добавляет новый объект и записывает обновленный массив обратно в файл.
// Файл: people.json
// Структура:
// type Person struct {
//  Name string `json:"name"`
//  Age int `json:"age"`
//  Email string `json:"email"`
// }
// 16. Удаление объекта из JSON-массива
// Описание: Напишите программу, которая читает JSON-массив из файла, удаляет
// объект с определенным значением поля и записывает обновленный массив
// обратно в файл.
// Файл: people.json
// Структура:
// type Person struct {
//  Name string `json:"name"`
//  Age int `json:"age"`
//  Email string `json:"email"`
// }
// 17. Поиск объекта в JSON-массиве
// Описание: Напишите программу, которая читает JSON-массив из файла и находит
// объект с определенным значением поля.
// Файл: people.json
// Структура:
// type Person struct {
//  Name string `json:"name"`
//  Age int `json:"age"`
//  Email string `json:"email"`
// }
// 18. Подсчет объектов в JSON-массиве
// Описание: Напишите программу, которая читает JSON-массив из файла и
// подсчитывает количество объектов в массиве.
// Файл: people.json
// Структура:
// type Person struct {
//  Name string `json:"name"`
//  Age int `json:"age"`
//  Email string `json:"email"`
// }
// 19. Десериализация JSON с произвольными полями
// Описание: Напишите программу, которая десериализует JSON-строку с
// произвольным набором полей в map[string]interface{}.
// Пример данных:
// {
//  "name": "John",
//  "age": 30,
//  "country": "USA",
//  "job": "Engineer"
// }
// 20. Сериализация и десериализация времени в JSON
// Описание: Напишите программу, которая сериализует и десериализует структуру
// с полем времени (time.Time) в JSON.
// Структура:
// type Event struct {
//  Name string `json:"name"`
//  Timestamp time.Time `json:"timestamp"`
// }
