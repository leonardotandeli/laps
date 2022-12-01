package controllers

import (
	"bufio"
	"fmt"
	"laps/src/database"
	"laps/src/models"
	"os"
	"time"
)

//CriaAtivoNoBanco armazena o hostname e senha no banco de dados
func CriaAtivoNoBanco(fileDir string) {

	f, err := os.Open(fileDir)

	if err != nil {
		fmt.Println("erro ao ler o arquivo")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	for scanner.Scan() {

		if len(scanner.Text()) >= 17 {

			char1 := scanner.Text()[1:2]
			char2 := scanner.Text()[3:4]
			char3 := scanner.Text()[5:6]
			char4 := scanner.Text()[7:8]
			char5 := scanner.Text()[9:10]
			char6 := scanner.Text()[11:12]
			char7 := scanner.Text()[13:14]

			var ativo models.Ativo

			ativo.HOSTNAME = char1 + char2 + char3 + char4 + char5 + char6 + char7

			ativo.SenhaAtual = BuscaSenha(`"` + ativo.HOSTNAME + `"`)
			ativo.DataSenhaAtual = time.Now()
			// ativo.SenhaAtual = "Pendente"
			database.DB.Create(&ativo)

			ChecaSenha(ativo.HOSTNAME, ativo.SenhaAtual)

		} else {
			fmt.Println("fim")
		}

	}

}

//ChecaSenha faz a verificação de senhas e só armazena uma nova senha se a senha atual for diferente da última senha armazenada
func ChecaSenha(hostname string, senhaAtual string) {
	var ativo models.Ativo
	database.DB.Where("hostname = ?", hostname).First(&ativo)

	if ativo.SenhaAtual == senhaAtual {
		//fmt.Println("senhas são iguais, nada será feito...")
	} else {

		ativo.AntePenultimaSenha = ativo.PenultimaSenha
		ativo.DataAntePenultimaSenha = ativo.DataPenultimaSenha
		ativo.PenultimaSenha = ativo.SenhaAtual
		ativo.DataPenultimaSenha = ativo.DataSenhaAtual
		ativo.SenhaAtual = senhaAtual
		ativo.DataSenhaAtual = time.Now()

		database.DB.Save(&ativo)
	}

}
