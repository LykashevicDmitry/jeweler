package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Получаем имя файла с доменами из аргументов командной строки
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Использование: go run main.go <имя_файла>")
		os.Exit(1)
	}
	filename := args[1]

	// Открываем файл
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Создаем карту для хранения доменов, сгруппированных по IP-адресу
	domains := make(map[string][]string)

	// Считываем каждую строку файла и получаем IP-адреса для каждого домена
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()

		// Получаем IP-адреса для домена
		ips, err := net.LookupIP(domain)
		if err != nil {
			fmt.Println("Ошибка при получении IP-адреса для домена", domain, err)
			continue
		}

		// Добавляем домен в карту, группируя по IP-адресу
		for _, ip := range ips {
			domains[ip.String()] = append(domains[ip.String()], domain)
		}
	}

	// Выводим домены, сгруппированные по IP-адресу
	for ip, domains := range domains {
		fmt.Println(ip)
		for _, domain := range domains {
			fmt.Println("  ", domain)
		}
	}
}
