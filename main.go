package main

import (
	"fmt"
	"laps/src/controllers"

	"github.com/jasonlvhit/gocron"
)

//task chama a função de busca da senhas.
func task() {
	controllers.LerOU()
}

//Cronjob define as configurações do gocron para executar as tarefas a cada 24 horas
func Cronjob() {
	// executa as tarefas a cada dia
	gocron.Every(1).Day().At("12:00").Do(task)

	// inicia todas as tarefas
	<-gocron.Start()
}

func main() {

	cm := "-command"
	// importModule := "import-module"
	// ad := " ActiveDirectory;"
	getAd := "Get-ADComputer"
	ds := "-Filter"
	d1 := "*"
	search := "-SearchBase"
	select1 := "|"
	select2 := "Select-Object"
	arg1 := "-ExpandProperty"
	arg2 := "Name"
	arg3 := "Out-File"

	//executa o comando net user no cmd
	fmt.Println(cm, getAd, ds, d1, search, "t", select1, select2, arg1, arg2, select1, arg3, "txt")

	// database.Dbconnection()
	// // controllers.ReadTXTFileOU()
	// go Cronjob()
}
