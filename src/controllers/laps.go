package controllers

import (
	"fmt"
	"os/exec"
	"strings"
)

func BuscaSenha(ativo string) string {
	//comando para buscar informações do ativo (hostname)
	ps := "powershell.exe"
	cm := "-command"
	ad := "Get-ADComputer -Identity"
	ad2 := "-Properties *"

	//executa o comando no cmd
	cmd := exec.Command(ps, cm, ad, ativo, ad2)
	//retorna os dados do cmd
	cmdOut, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return "erro"
	}
	cmdReturn := string(cmdOut)

	//filtra onde estará a informação
	var senhaWithoutSpace = ""
	if strings.Contains(cmdReturn, "ms-Mcs-AdmPwd") {
		SenhaSliceInitial := strings.Index(cmdReturn, "ms-Mcs-AdmPwd")
		SenhaSliceEnding := strings.Index(cmdReturn, `ms-Mcs-AdmPwdExpirationTime`)
		//fatia os dados com a informação recebida
		senha := string(cmdReturn)[SenhaSliceInitial+38 : SenhaSliceEnding]
		//Formata os dados
		senhaWithoutSpace = strings.TrimSpace(senha)
	} else {
		senhaWithoutSpace = "Nenhuma senha encontrada"
	}

	return senhaWithoutSpace

}
