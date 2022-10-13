package models

import "time"

// Ativo define a estrutura de um ativo.
type Ativo struct {
	ID                     uint
	HOSTNAME               string `gorm:"uniqueIndex"`
	SenhaAtual             string
	DataSenhaAtual         time.Time
	PenultimaSenha         string
	DataPenultimaSenha     time.Time
	AntePenultimaSenha     string
	DataAntePenultimaSenha time.Time
}
