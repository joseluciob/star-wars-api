# API Star Wars

## Indíce
* [Sobre o projeto](#about-problem)
    * [Desafio](#about-problem)
    * [Tecnologias](#about-techs)
* [Requisitos](#requirements)
* [Instalação](#instalation)
* [Funcionamento](#running)
* [Testes](#tests)
* [TODO](#todo)

<a name="about-problem"></a>
## Desafio
Criar uma API para obter informações do serviço público https://swapi.dev. Salvar nome, clima e terreno de cada planeta. Os filmes de cada planeta também devem ser salvos.

#### Funcionalidades desejadas:
* Carregar um planeta da API através do Id
* Listar os planetas
* Buscar planeta por nome
* Buscar por ID,
* Remover planeta


<a name="about-techs"></a>
## Tecnologias
* Golang 
    * [Fast HTTP](https://github.com/valyala/fasthttp) - Cliente HTTP
    * [Gorm](https://github.com/go-gorm/gorm) - ORM
    * [Cobra](https://github.com/spf13/cobra) - CLI
    * [Viper](https://github.com/spf13/viper) - Configs 
    * [Gin](https://github.com/gin-gonic/gin) - Web Framework
    * [Testify](https://github.com/stretchr/testify) - Testes 
    * [Zap](https://github.com/uber-go/zap) - Logs
* Postgresql
* Docker
* Docker compose
* Swagger

<a name="requirements"></a>
## Requisitos
- Git
- Docker
- Docker composer
- Golang >= v1.19


<a name="instalation"></a>
## Instalação

Na raiz do projeto existe um arquivo denominado `.env.example`, é necessário renomea-lo para `.env`. Pode ser feito através do comando:
```sh
cp .env.example .env
```

Inicie o container docker:
```sh 
docker-compose up -d
```

Três containers serão iniciados: 
* **Importer** - Faz a importação de planetas e filmes. Foi concebido para funcionar como uma cron.
* **App** - Aplicação
* **Postgresdb** - Banco de dados

A aplicação estará disponível através do endereço: http://localhost:8190/api/v1/docs/index.html


<a name="running"></a>
## API - Funcionamento (endpoints)
```
GET /planets
GET /planets?name=Zolan
GET /planets/{id}
DELETE /planets/{id}
GET /docs/index.html
```

Para facilitar o consumo dos recursos, existe uma coleção do postman na pasta `/docs/postman`. 

<a name="running-api-docs"></a>
## Tests

Para executar os testes, rode o comando:
```docker-compose exec app go test -failfast -vet=off -race -timeout=1m ./...```

Para exibir a cobertura dos testes, acrescente o parametro **`--cover`** no comando acima.

<a name="running-api-docs"></a>
## TODO
- [ ] Aumentar cobertura dos testes unitários;
- [ ] Melhorar log de erros e eventos;
- [ ] Adicionar recurso para importação de apenas um planeta.