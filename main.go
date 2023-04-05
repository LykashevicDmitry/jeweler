package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Выбираем файл с доменами
	filePath := selectFile()
	if filePath == "" {
		fmt.Println("Не удалось выбрать файл")
		os.Exit(1)
	}

	// Открываем файл с доменами
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Не удалось открыть файл %s\n", filePath)
		os.Exit(1)
	}
	defer file.Close()

	// Считываем домены из файла и находим для каждого IP-адрес
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		ipAddr, err := net.LookupIP(domain)
		if err != nil {
			fmt.Printf("%s - Не удалось найти IP-адрес\n", domain)
		} else {
			fmt.Printf("%s - %s\n", domain, ipAddr[0].String())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при считывании файла")
		os.Exit(1)
	}
}

// Функция выбора файла с доменами
func selectFile() string {
	var filePath string
	fmt.Print("Введите путь к файлу с доменами: ")
	fmt.Scan(&filePath)
	_, err := os.Stat(filePath)
	if err != nil {
		return ""
	}
	return filePath
}
