package main

import (
	"laps/config"
	"laps/src/controllers"
	"laps/src/database"

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
	config.CarregarConfiguracao()
	database.Dbconnection()

	go Cronjob()
}
