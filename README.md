[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

# civi-backend-challenge
civi backend challenge é um repositório criado com o intuito de resolver o desafio de backend da civi app, utilizando golang.

## Requisitos

* [GoLang](https://golang.org/) - compilador da linguagem Go
* [Gin](https://github.com/gin-gonic/gin) - framework web para golang

## Pré Configuração

* Configurar client git para baixar repositórios internos via ssh:

```bash 
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

### Antes de tudo

``` bash
# Crie o arquivo .env com base no arquivo de exemplo para carregar as variaveis de ambiente

# Baixe as dependências do projeto
$ make deps
```

### Se for utilizar docker
``` bash
# Inicie o serviço do docker
$ systemctl start docker

# Liste os containers disponíveis
$ sudo docker ps -a

# Inicie o container específico
$ docker build . -t go-dock
$ docker run -d -p 8000:8000 go-dock
```

### Inicializando
``` bash
# Comando para buildar a aplicação e garantir que nada está quebrado
$ make build

# Comando para subir a aplicação
$ make run
```

### Testando
```bash
# Executa o comando para executar os testes unitários
make run-tests
```

### Swagger
```bash
# http://localhost:8090/swagger/index.html
```

### Estrutura do projeto
```
├── data/
│   ├── points.json
├── handlers/
│   ├── points.go
├── models/
│   ├── error.go
│   ├── points.go
├── routers/
│   ├── router.go
├── services/
│   ├── points.go
│
└── main.go
└── Makefile
└── README.md
```