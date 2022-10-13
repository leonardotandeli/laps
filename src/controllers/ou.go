package controllers

import (
	"bufio"
	"fmt"
	"os"
)

//LerOU faz a leitura de todas as linhas do arquivo ou.txt
func LerOU() {

	f, err := os.Open("ou/ou.txt")

	if err != nil {
		fmt.Println("erro ao ler a OU")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ADScan(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		return
	}

}
