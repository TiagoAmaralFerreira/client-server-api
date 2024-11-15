package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TiagoAmaralFerreira/client-server-api/db"
	"github.com/TiagoAmaralFerreira/client-server-api/models"
	"gorm.io/gorm"
)

func main() {
	// Inicializa o banco de dados
	database := db.InitDB()

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Timeout para obter cotação
		apiCtx, cancelApi := context.WithTimeout(ctx, 200*time.Millisecond)
		defer cancelApi()

		cotacao, err := fetchCotacao(apiCtx)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao obter cotação: %v", err), http.StatusInternalServerError)
			return
		}

		// Timeout para salvar no banco
		dbCtx, cancelDb := context.WithTimeout(ctx, 10*time.Millisecond)
		defer cancelDb()

		err = saveCotacao(dbCtx, database, cotacao.Bid)
		if err != nil {
			log.Printf("Erro ao salvar no banco: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cotacao)
	})

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fetchCotacao(ctx context.Context) (*models.Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	bid := result["USDBRL"]["bid"]
	return &models.Cotacao{Bid: bid}, nil
}

func saveCotacao(ctx context.Context, db *gorm.DB, bid string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		cotacao := models.Cotacao{Bid: bid}
		return db.WithContext(ctx).Create(&cotacao).Error
	}
}
