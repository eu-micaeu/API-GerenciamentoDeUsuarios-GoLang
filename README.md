# API de Gerenciamento de Usuários - GoLang

Este repositório contém uma API desenvolvida em GoLang (Golang) para o gerenciamento de usuários. A API permite o cadastro, autenticação, atualização e remoção de usuários, além de fornecer funcionalidades para controle de permissões e papéis de usuário.

## Funcionalidades

- **Cadastro de Usuários**: Permite a criação de novos usuários com dados como nome, email e senha.
- **Autenticação**: Suporte a autenticação JWT (JSON Web Token), permitindo login seguro e controle de sessão.
- **Gerenciamento de Perfis**: Atualização e visualização dos perfis dos usuários.
- **Deleção de Usuário**: Permite a remoção de contas de usuários.

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação principal utilizada.
- **Gin**: Framework web utilizado para roteamento e manipulação de requisições HTTP.
- **GORM**: ORM (Object-Relational Mapping) utilizado para interagir com o banco de dados.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenamento de usuários.
- **JWT**: Utilizado para autenticação e controle de sessão.

## Pré-requisitos

- **Go**: Versão 1.16 ou superior.
- **PostgreSQL**: Configurado e em execução.

## Configuração

1. Clone o repositório:

    ```bash
    git clone https://github.com/eu-micaeu/API-GerenciamentoDeUsuarios-GoLang.git
    ```

2. Entre no diretório do projeto:

    ```bash
    cd API-GerenciamentoDeUsuarios-GoLang
    ```

3. Configure as variáveis de ambiente no arquivo `.env`:

    ```bash
    DB_HOST=<seu_host>
    DB_USER=<seu_usuario>
    DB_PASS=<sua_senha>
    DB_NAME=<nome_do_banco_de_dados>
    JWT_SECRET=<seu_segredo_jwt>
    ```

4. Instale as dependências:

    ```bash
    go mod tidy
    ```

5. Execute a aplicação:

    ```bash
    go run main.go
    ```
    
## Endpoints

### Autenticação

- **POST /register**: Cadastro de novos usuários.
- **GET /login**: Login de usuários e geração de token JWT.

### Usuários

- **PUT /update**: Atualiza os dados de um usuário.
- **DELETE /delete**: Deleta um usuário.

## Testes

Para rodar os testes:

```bash
go test ./...
