package database

import (
	"laps/config"
	"laps/src/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

//Dbconnection inicia a conexão com o banco de dados.
func Dbconnection() {

	DB, err = gorm.Open(mysql.Open(config.StringBanco))
	if err != nil {
		log.Panic("Erro ao conectar no banco de dados")
	}
	DB.AutoMigrate(&models.Ativo{})
}
