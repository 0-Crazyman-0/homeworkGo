package main

// "fmt"
// "os"
// "io"

func main() {
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }
	// defer file.Close()
	// fmt.Println("File created successfully")

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()

	// data, _ := io.ReadAll(file)
	// fmt.Println(string(data))
	// buf := make([]byte, 1024)

	// n, err := file.Read(buf)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, buf[:n])

	// _, err = file.Write([]byte("Hello, Go!"))
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return
	// }
	// fmt.Println("Data written successfully")

	// file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()
	// fmt.Println("File opened successfully")

	// buf := make([]byte, 1024)
	// n, err := file.Read(buf)
	// if err != nil && err.Error() != "EOF" {
	// 	fmt.Println("Eror")
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, buf[:n])

	// _, err = file.Write([]byte("Hello, Go!"))
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return
	// }
	// fmt.Println("Data written successfully")

	// 	file, err := os.Open("largefile.txt")
	// if err != nil {
	// fmt.Println("Error opening file:", err)
	// return
	// }
	// defer file.Close()
	// buffer := make([]byte, 14)
	// for {
	// n, err := file.Read(buffer)
	// if err == io.EOF {
	// break
	// }
	// if err != nil {
	// fmt.Println("Error reading file:", err)
	// return
	// }

	// fmt.Print(string(buffer[:n]))
	// }
}
