# Loja Web com Go (Arquitetura MVC)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)

## 📖 Sobre o Projeto

Este projeto é uma aplicação web desenvolvida em **Go** para o gerenciamento de um catálogo de produtos. O principal objetivo foi aplicar e consolidar conhecimentos em desenvolvimento backend com Go, utilizando uma arquitetura organizada e escalável inspirada no padrão **MVC (Model-View-Controller)**.

A aplicação implementa um CRUD completo e se conecta a um banco de dados **PostgreSQL** para persistir os dados dos produtos.

## 🖼️ Visualização

![Screenshot da Aplicação](https://i.imgur.com/rS2T41B.png)

## ✨ Funcionalidades

O projeto implementa todas as operações essenciais de um sistema de gerenciamento (CRUD):

* **Listagem de Produtos (Read)**: A página inicial exibe todos os produtos cadastrados no banco de dados.
* **Criação de Produtos (Create)**: Rota e formulário para adicionar novos produtos ao catálogo.
* **Edição de Produtos (Update)**: Formulário para editar os dados de um produto existente.
* **Exclusão de Produtos (Delete)**: Funcionalidade para remover um produto do catálogo.

## 🏛️ Arquitetura do Projeto

O código foi estruturado para separar as preocupações, seguindo uma abordagem modular:

* `main.go`: Ponto de entrada da aplicação. Inicializa o servidor e carrega as rotas.
* `routes/`: Define todas as rotas da aplicação (ex: `/`, `/new`, `/edit`, `/delete`) e as associa aos seus respectivos *controllers*.
* `controllers/`: Contém a lógica de negócio. Recebe as requisições HTTP, interage com os *models* e envia os dados para as *views*.
* `models/`: Responsável por toda a comunicação com o banco de dados. Contém as `structs` e as funções para as operações de CRUD.
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
* Um banco de dados criado para a aplicação (ex: `loja_go`) com a tabela `produtos` devidamente configurada.

### Passos

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/CiceroGGS/loja-web-go.git](https://github.com/CiceroGGS/loja-web-go.git)
    ```
    *(**Nota:** Substitua a URL pelo link correto do seu repositório)*

2.  **Navegue até a pasta do projeto:**
    ```bash
    cd loja-web-go
    ```

3.  **Configure a Conexão com o Banco de Dados:**
    Abra o arquivo `db/db.go` e atualize a string de conexão com as suas credenciais do PostgreSQL (usuário, senha, nome do banco).

4.  **Execute a Aplicação:**
    ```bash
    go run main.go
    ```

5.  Abra seu navegador e acesse `http://localhost:8000`.

## 👨‍💻 Autor

**Cícero Guilherme**.

[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/cicero-guilherme-a9473a260/)
[![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/CiceroGGS/)