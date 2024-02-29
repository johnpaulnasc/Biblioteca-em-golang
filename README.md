# Sistema de Gerenciamento de Biblioteca

Este projeto é um sistema de gerenciamento de biblioteca simples, escrito em Go, que permite aos usuários adicionar, atualizar e excluir livros de uma base de dados SQLite.

## Tecnologias Utilizadas

- **Go**: A linguagem de programação usada para desenvolver o backend do sistema.
- **SQLite**: O banco de dados usado para armazenar informações sobre os livros.
- **Gorilla Mux**: Um roteador e manipulador de URL para Go.

## Como Configurar

1. **Clone o Repositório**:
git clone https://github.com/johnpaulnasc/biblioteca-golang.git 
cd biblioteca-golang

2. **Instale as Dependências**:

go mod download

3. **Inicie o Servidor**:

go run ./cmd/api/main.go

## Como Usar

Após iniciar o servidor, você pode acessar a aplicação no navegador navegando até `http://localhost:8080`. Aqui estão algumas operações que você pode realizar:

- **Visualizar Livros**: Acesse `http://localhost:8080/books` para ver a lista de livros.
- **Adicionar Livro**: Use o formulário na página inicial para adicionar um novo livro.
- **Atualizar Livro**: Envie uma requisição PUT para `/books/{id}` com os detalhes do livro atualizados.
- **Excluir Livro**: Envie uma requisição DELETE para `/books/{id}` para excluir um livro.


