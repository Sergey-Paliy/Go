package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}
	filePth := os.Args[1]
	var fileName, fileExt string
	var index1, index2 int
	for i := len(filePth) - 1; i > 0; i-- {
		if filePth[i] == '/' {
			index1 = i
			break
		}
	}
	for i := len(filePth) - 1; i > 0; i-- {
		if filePth[i] == '.' {
			index2 = i
			break
		}
	}
	if index2 == 0 {
		fileName = filePth[index1+1:]
	} else {
		fileName = filePth[index1+1 : index2]
		fileExt = filePth[index2+1:]
	}
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
