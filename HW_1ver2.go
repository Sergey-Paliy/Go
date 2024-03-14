package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	var fileName, fileExt string
	var index1, index2 int
	index1 = strings.LastIndexByte(filePth, '/')
	index2 = strings.LastIndexByte(filePth, '.')
	if index2 == -1 {
		fileName = filePth[index1+1:]
	} else {
		fileName = filePth[index1+1 : index2]
		fileExt = filePth[index2+1:]
		// Напишите код, который выведет следующее
		// filename: <name>
		// extension: <extension>

		// Подсказка. Возможно вам понадобится функция strings.LastIndex
		// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
		// ) Они могут помочь для проверки решения
	}
	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
