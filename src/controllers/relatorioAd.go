package controllers

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

//ADScan gera um relatório de hostnames presentes na OU informada.
func ADScan(OU string) string {

	ps := "powershell.exe"
	cm := "-command"
	getAd := "Get-ADComputer"
	ds := "-Filter"
	d1 := "*"
	search := "-SearchBase"
	select1 := "|"
	select2 := "Select-Object"
	arg1 := "-ExpandProperty"
	arg2 := "Name"
	arg3 := "Out-File"
	file := "relatorio/ativos" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + ".txt"

	//executa o comando no cmd
	cmd := exec.Command(ps, cm, getAd, ds, d1, search, OU, select1, select2, arg1, arg2, select1, arg3, file)
	//retorna os dados do cmd
	cmdOut, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return "erro"
	}

	CriaAtivoNoBanco(file)
	cmdReturn := string(cmdOut)

	return cmdReturn
}
