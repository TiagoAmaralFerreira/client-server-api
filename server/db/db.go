package db

import (
	"log"

	"github.com/TiagoAmaralFerreira/client-server-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Conexão com o banco de dados SQLite.
func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("cotacoes.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = db.AutoMigrate(&models.Cotacao{})
	if err != nil {
		log.Fatalf("Erro ao migrar banco de dados: %v", err)
	}

	return db
}
