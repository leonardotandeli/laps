package database

import (
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
	stringDeConexao := "laps:laps@1010@tcp(localhost:3306)/laps?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(stringDeConexao))
	if err != nil {
		log.Panic("erro ao conectar no db")
	}
	DB.AutoMigrate(&models.Ativo{})
}
