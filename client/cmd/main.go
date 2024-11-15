package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	link := "http://localhost:8080/cotacao"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	var cotacao Cotacao
	err = json.NewDecoder(resp.Body).Decode(&cotacao)
	if err != nil {
		log.Fatalf("Erro ao decodificar resposta: %v", err)
	}

	content := fmt.Sprintf("Dólar: %s\n", cotacao.Bid)
	err = ioutil.WriteFile("cotacao.txt", []byte(content), 0644)
	if err != nil {
		log.Fatalf("Erro ao salvar arquivo: %v", err)
	}

	fmt.Println("Cotação salva em cotacao.txt")
}
