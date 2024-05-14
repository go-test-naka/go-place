# go-place

Projeto de teste que utiliza:
- Endpoint HTTP GET/POST para consulta e cadastro
- Mysql para armazenar dados.
- Docker-compose para executar a aplicação

## Tecnologias
- Golang
- Mysql
- Docker

## Mysql

```sh
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mysql:latest
```
## Migration

```sql
CREATE TABLE sys.Person (
	id BIGINT UNSIGNED auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	birthday DATETIME NOT NULL,
	country varchar(100) NOT NULL,
	CONSTRAINT Person_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

```

## Environment variables

```sh
source set_environment.sh
```

Um script bash é executado em um subshell, e não no atual.
Para executar o script dentro do shell atual, utilizarmos o `source`.

Quando criamos uma variável, ela existe apenas dentro do shell atual.
Quando executamos uma aplicação, ela é executada em um processo filho, e não possui as variáveis definidas no processo pai.
Para permitir que as variáveis estejam disponíveis nos processos filhos, utilizamos o `export`

