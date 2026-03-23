# Devbook API

A **Devbook API** é o motor por trás de uma rede social para desenvolvedores, gerenciando autenticação, perfis de usuários e um sistema de publicações.

## 🛠️ Endpoints da API

Abaixo estão detalhadas as rotas disponíveis, organizadas por módulo.

### 🔐 Autenticação
*   **POST `/login`**: Permite a autenticação de usuários no sistema.

### 👤 Usuários
| Método | Rota | Descrição | Requer Auth |
| :--- | :--- | :--- | :--- |
| **POST** | `/usuarios` | Realiza o cadastro de um novo usuário. | Não |
| **GET** | `/usuarios` | Busca usuários cadastrados no sistema. | Sim |
| **GET** | `/usuarios/{id}` | Retorna as informações de um usuário específico. | Sim |
| **PUT** | `/usuarios/{id}` | Atualiza os dados de um usuário. | Sim |
| **DELETE** | `/usuarios/{id}` | Remove permanentemente a conta de um usuário. | Sim |
| **POST** | `/usuarios/{id}/seguir` | Permite seguir outro usuário da rede. | Sim |
| **POST** | `/usuarios/{id}/parar-seguir` | Permite deixar de seguir um usuário. | Sim |
| **GET** | `/usuarios/{id}/seguidores` | Lista todos os seguidores de um usuário. | Sim |
| **GET** | `/usuarios/{id}/seguindo` | Lista quem o usuário está seguindo. | Sim |
| **POST** | `/usuarios/{id}/atualizar-senha` | Permite que o usuário altere sua senha de acesso. | Sim |

### 📝 Publicações
| Método | Rota | Descrição | Requer Auth |
| :--- | :--- | :--- | :--- |
| **POST** | `/publicacoes` | Cria uma nova publicação no feed. | Sim |
| **GET** | `/publicacoes` | Retorna as publicações do feed (próprias e de quem segue). | Sim |
| **GET** | `/publicacoes/{publicacaoId}` | Busca uma publicação específica por seu ID. | Sim |
| **PUT** | `/publicacoes/{publicacaoId}` | Edita o conteúdo de uma publicação existente. | Sim |
| **DELETE** | `/publicacoes/{publicacaoId}` | Remove uma publicação do sistema. | Sim |
| **GET** | `/usuarios/{usuarioId}/publicacoes` | Lista todas as publicações de um usuário específico. | Sim |
| **POST** | `/publicacoes/{publicacaoId}/curtir` | Adiciona uma curtida a uma publicação. | Sim |
| **DELETE** | `/publicacoes/{publicacaoId}/curtir` | Remove a curtida de uma publicação. | Sim |

---

## 🛠️ Guia de Configuração e Instalação

Este guia descreve como preparar o ambiente para executar a Devbook API. Note que esta aplicação utiliza o banco de dados **MySQL**.

### 1. Configuração das Variáveis de Ambiente

A aplicação utiliza variáveis de ambiente para gerenciar credenciais sensíveis e configurações de porta. 

Crie um arquivo chamado `.env` na raiz do projeto e adicione as seguintes chaves (preencha os valores de acordo com seu ambiente local):

```env
DB_USER=seu_usuario_banco
DB_PASSWORD=sua_senha_do_banco
DB_NAME=nome_do_seu_banco
PORT=5000 (ou outra qualquer)
SECRET_KEY=uma_chave_segura_para_o_jwt
```

---

## 2. Instalação das Dependências
Certifique-se de ter o Go instalado em sua máquina. Para instalar todas as dependências necessárias. 
Para o funcionamento das rotas de usuários, publicações e login, execute os seguintes comandos no terminal:
### Inicialize o módulo Go (caso ainda não tenha feito)
```go mod init devbook-api```

### Instalação dos drivers e utilitários
```bash
go get github.com/go-sql-driver/mysql          # Driver de conexão com MySQL
go get github.com/gorilla/mux                   # Roteador HTTP
go get github.com/joho/godotenv                 # Gerenciamento de arquivo .env
go get github.com/dgrijalva/jwt-go              # Implementação de tokens JWT
go get golang.org/x/crypto/bcrypt               # Criptografia de senhas
go get github.com/badoux/checkmail              # Validação de e-mails
go get filippo.io/edwards25519                  # Dependência de criptografia/matemática
```
Após executar os comandos acima, rode o comando abaixo para garantir que todas as dependências estejam sincronizadas:
```go mod tidy ```


---
## 3. Banco de Dados
Esta aplicação foi projetada para funcionar com MySQL. Certifique-se de que o serviço do MySQL está ativo e que você criou o banco de dados especificado na variável DB_NAME antes de iniciar a API.
## 4. Executando a Aplicação
Com as variáveis configuradas e as dependências instaladas, você pode iniciar o servidor com o comando:
go run main.go

## 🚀 Como Executar
1. Clone este repositório.
2. Configure as variáveis de ambiente para conexão com o banco de dados.
3. Execute `go run main.go`.

## 🔑 Autenticação
Para as rotas que exigem autenticação, é necessário enviar o token gerado no login através do header:
`Authorization: Bearer <seu-token-aqui>`