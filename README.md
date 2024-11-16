# Desafio Go Expert - Client-Server-API

## Proposta

Olá dev, tudo bem?
 
Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.
 
Você precisará nos entregar dois sistemas em Go:
- client.go
- server.go
 
Os requisitos para cumprir este desafio são:
 
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
 
O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

## Executando o sistema

O projeto contem dois diretórios, um para a aplicação do servidor e outro para a aplicação do cliente.

Esteja na raiz do diretório `client-server-api`, abra dois terminais.

### 1. Servidor

Entre no diretório do servidor:

```
cd server/
```

Execute o servidor com o seguinte comando abaixo:

```
go run cmd/main.go
```

Caso resulte em succeso, nosso servidor rodando em http://localhost:8080.

### 2. Cliente

Entre no diretório do cliente:

```
cd client/
```

Execute o cliente com o seguinte comando abaixo:

```
go run cmd/main.go
```

Caso resulte em sucesso, será printado no console a seguinte informação:

```
Cotação salva em cotacao.txt
```

Para visualizar o resultado gravado em arquivo, execute o comando abaixo:
> O arquivo está salvo no seguinte path: _client-server-api/client/cotacao.txt_

```
cat cotacao.txt
```
