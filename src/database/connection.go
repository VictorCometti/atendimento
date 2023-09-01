package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func getStringConnection() (stringConnection string, erro error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("src/config")

	if erro = viper.ReadInConfig(); erro != nil {
		log.Fatalf("Erro ao tentar ler o árquivo de configuração: Erro: %v", erro)
	}

	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	hostname := viper.GetString("database.hostname")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database_name")

	stringConnection = "host=" + hostname + " port=" + port + " user=" + username + " password=" + password + " dbname=" + database + " sslmode=disable"

	return

}

func GetConnection() (db *sql.DB, erro error) {
	stringConnection, erro := getStringConnection()
	if erro != nil {
		log.Fatalf("Erro ao tentar pegar a string de conexão. Erro: %v", erro)
	}

	db, erro = sql.Open("postgres", stringConnection)
	if erro != nil {
		log.Fatalf("Erro ao tentar abrir conexao com banco de dados: Erro: %v", erro)
	}

	if erro = db.Ping(); erro != nil {
		log.Fatalf("Erro ao tentar consultar o ping da conexão: Erro: %v", erro)
	}

	return
}
