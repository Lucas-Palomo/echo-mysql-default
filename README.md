# README

Este projeto é um CRUD de uma API Rest para manipulação produtos


> O projeto iniciaria o banco sozinho
>
> É obrigatório renomear o .env-example para .env, somente assim o projeto pode iniciar
>
> Preencha o .env com as informações necessárias para o ambiente iniciar

## Requisitos

- Docker
- Docker Compose
- Golang
- Postman (Opcional)

---

## Executando o Projeto

Para executar o projeto primeiro inicie o docker

```shell
docker-compose up -d --force-recreate
```

Enquanto o container inicia, você pode baixar as dependecias do projeto

```shell
go mod tidy
```

Agora basta esperar o servidor mysql iniciar e executar o projeto

```shell
go run cmd/echo_mysql.go
```

---

# Mais informações

Pensando em um mini ecommerce, decidi criar a tabela de produtos no seguinte modelo

```sql
CREATE TABLE IF NOT EXISTS `products`
(
    id          serial primary key,
    name        varchar(100)    not null,
    description text            not null,
    price       bigint unsigned not null,
    online      tinyint         not null default 0,
    created_at  datetime        not null default now(),
    modified_at datetime                 default null
);
```

Deste modo por regra de négocio ficou definido:

- Que todos os produtos começam desabilitados (offline)
- Que o preço é definido em centavos (para evitar divergências de números flutuantes entre banco e aplicação)

## Rotas

Todas as rotas podem ser testadas pelo postman, basta importar o arquivo `Commerce.postman_collection.json`

---

GET `/products/`

Rota responsável pela listagem de produtos

---

- POST `/products/`

Rota utilizada para inserção de um novo produto

```json
{
  "name": "notebook",
  "description": "",
  "price": 2000
}
```

---

- DELETE `/products/:id/`

Rota dedicada para exclusão de um único produto, o id deve corresponder a um produto existente

---

- PUT `/products/:id/`

Rota dedicada para atualização de um único produto, o id deve corresponder a um produto existente

```json
{
  "name": "notebook gamer",
  "description": "notebook gamer",
  "price": 2000,
  "online": true
}
```

---

## Informações Gerais

### Logs

O projeto gera logs dois tipos de log

- Até o echo iniciar todo log é disparado para o std
- Após o echo iniciar todo log é guardado em arquivo

### mysql.cnf

É um arquivo para otimização do mysql, nele podemos definir, por exemplo, o máximo de conexões permitidas

### .env

Arquivo com variáveis de ambiente

- SERVER_ADDR
    - Define o endereço da API
- LOG_FILE
    - Define a saida do log no formato de arquivo
- MYSQL_ADDR
    - Define o endereço do banco de dados
- MYSQL_USER
    - Define o usuário do banco de dados
- MYSQL_PASSWORD
    - Define a senha do banco de dados
- MYSQL_DATABASE
    - Define o banco de dados a ser usado
- MYSQL_RANDOM_ROOT_PASSWORD
    - Define a senha do usuário do root como aleatória, como não usamos este usuário definimos esse acesso como desconhecido.

```text
SERVER_ADDR=localhost:3000
LOG_FILE=./service.log

MYSQL_ADDR=localhost:3306
MYSQL_USER=toor
MYSQL_PASSWORD=toor
MYSQL_DATABASE=commerce
MYSQL_RANDOM_ROOT_PASSWORD=true
```

### PS
Utilizei o podman, uma alternativa ao docker.