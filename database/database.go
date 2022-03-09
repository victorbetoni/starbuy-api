package database

import (
	"starbuy/util"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db sqlx.DB

// Connect vai criar uma conexão com o banco utilizando as variáveis de ambiente definidas na config
func Connect() (err error) {

	config := util.GrabConfig()

	//Mantenha o SSLMode ativado, caso contrario ele ficara direcionando para a localhost
	/*dataSource := fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=require", config.HostAddress, config.Port, config.Username, config.Password, config.Schema)*/

	var database *sqlx.DB
	if database, err = sqlx.Open(config.Driver, "postgres://ntmqwvedocrexr:36074513dd76c22fc09a0eea3dc2cd12dd6a1434fc346e45a005709c0a60ffe8@ec2-34-236-87-247.compute-1.amazonaws.com:5432/d324nv8ovevvpk"); err != nil {
		return err
	}
	db = *database
	return
}

func GrabDB() *sqlx.DB {
	if err := db.Ping(); err != nil {
		Connect()
	}
	return &db
}
