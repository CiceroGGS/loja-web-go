# Loja Web com Go (Arquitetura MVC)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

## 📖 Sobre o Projeto

Este projeto é uma aplicação web completa para o gerenciamento de um catálogo de produtos, desenvolvida inteiramente em Go. O principal objetivo foi aplicar uma arquitetura organizada e escalável, inspirada no padrão **MVC (Model-View-Controller)**, para separar as responsabilidades do código e facilitar a manutenção.

A aplicação se conecta a um banco de dados **PostgreSQL** para persistir os dados dos produtos.

> 🚀 Projeto em desenvolvimento, com funcionalidades de listagem e criação implementadas.

## ✨ Funcionalidades

* **Listagem de Produtos**: A página inicial exibe todos os produtos cadastrados no banco de dados em uma tabela.
* **Criação de Produtos**: Rota e formulário para adicionar novos produtos ao catálogo.
* **(Planejado)** Edição e exclusão de produtos existentes.

## 🏛️ Arquitetura do Projeto

O código foi estruturado para separar as preocupações, seguindo uma abordagem modular:

* `main.go`: Ponto de entrada da aplicação. Inicializa o servidor e carrega as rotas.
* `routes/`: Define todas as rotas da aplicação (ex: `/`, `/new`) e as associa aos seus respectivos *controllers*.
* `controllers/`: Contém a lógica de negócio. Recebe as requisições HTTP, interage com os *models* para buscar ou manipular dados e envia os dados para as *views* (templates).
* `models/`: Responsável por toda a comunicação com o banco de dados. Contém as `structs` que representam os dados e as funções para realizar as operações de CRUD (Create, Read, Update, Delete).
* `db/`: Módulo dedicado para configurar e estabelecer a conexão com o banco de dados PostgreSQL.
* `templates/`: Camada de visualização (View) da aplicação. Contém os arquivos HTML que são renderizados para o usuário.

## 🛠️ Tecnologias Utilizadas

* **Backend**: Go
* **Banco de Dados**: PostgreSQL
* **Frontend**: HTML5, Bootstrap 4

## 🚀 Como Executar o Projeto

Siga os passos abaixo para rodar a aplicação localmente.

### Pré-requisitos

* [Go](https://go.dev/doc/install) instalado.
* [PostgreSQL](https://www.postgresql.org/download/) instalado e em execução.
* Um banco de dados criado para a aplicação (ex: `loja_go`).

### Passos

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/seu-usuario/seu-repositorio.git](https://github.com/seu-usuario/seu-repositorio.git)
    ```

2.  **Navegue até a pasta do projeto:**
    ```bash
    cd seu-repositorio
    ```

3.  **Configure a Conexão com o Banco de Dados:**
    Abra o arquivo `db/db.go` e atualize a string de conexão com as suas credenciais do PostgreSQL (usuário, senha, nome do banco).

4.  **Execute a Aplicação:**
    ```bash
    go run main.go
    ```

5.  Abra seu navegador e acesse `http://localhost:8000`.

## 📁 Estrutura de Arquivos

```
/
├── controllers/
│   └── produtos.go
├── db/
│   └── db.go
├── models/
│   └── produtos.go
├── routes/
│   └── routes.go
├── templates/
│   └── index.html
├── go.mod
├── go.sum
└── main.go
```
## 👨‍💻 Autor

Feito com ❤️ por **[Cicero]**.

[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/cicero-guilherme-a9473a260/)
[![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/CiceroGGS/)
