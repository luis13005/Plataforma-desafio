# 🚀 Plataforma Desafio - gRPC User Management

[![Go](https://skillicons.dev/icons?i=go,docker,)](https://skillicons.dev)

> Um projeto moderno de gerenciamento de usuários utilizando gRPC, construído com Go e containerizado com Docker.

## 📋 Sobre o Projeto

Este projeto implementa um sistema de gerenciamento de usuários usando **gRPC** (Google Remote Procedure Call), permitindo operações de cadastro e atualização de usuários de forma eficiente e escalável.

## ✨ Principais Funcionalidades

- 🔐 **Cadastro de Usuários** - Registro seguro com validação de dados
- 📝 **Atualização de Perfil** - Modificação de informações do usuário
- 🔒 **Autenticação** - Sistema de verificação de senha com bcrypt
- 📡 **API gRPC** - Comunicação eficiente com Protocol Buffers

## 🛠️ Tecnologias Utilizadas

| Tecnologia | Descrição |
|------------|-----------|
| **Go** | Linguagem principal do projeto |
| **gRPC** | Framework de comunicação RPC |
| **Protocol Buffers** | Serialização de dados |
| **Docker** | Containerização da aplicação |
| **Makefile** | Automação de tarefas |
| **Evans** | Cliente gRPC para testes |
| **GORM** | ORM para Go |
| **bcrypt** | Criptografia de senhas |

## 🚀 Como Executar

### Pré-requisitos

- Go 1.19+
- Docker & Docker Compose
- Make

### Instalação

1. **Clone o repositório**
```bash
git clone https://github.com/luis13005/plataforma-desafio.git
cd plataforma-desafio
```
2. **Execute o Docker**
```bash
docker-compose up
```
3. **Rode o GO**
   ```bash
   go run framework/cmd/server/server.go -port 50051
   ```
### Testando a API

**Usando Evans (Cliente gRPC):**
```bash
make evans
# Cadastrar usuário
call CreateUser

# Atualizar usuário
call UpdateUser
```
## 🌟 Funcionalidades Implementadas

### Cadastro de Usuário
- Validação de dados de entrada
- Criptografia de senha com bcrypt
- Geração de tokens únicos
- Timestamps automáticos

  ### Atualização de Usuário
- Verificação de senha atual
- Validação de token

### Segurança
- Senhas criptografadas com bcrypt
- Validação de tokens
- Sanitização de dados de entrada

## ⚠️ Observações Importantes

- **Porta do Servidor**: O Go e o Evans devem rodar na **mesma porta** (padrão: 50051)
- **Alteração de Porta**: Se quiser alterar a porta, lembre-se de:
  - Modificar o comando: `go run framework/cmd/server/server.go -port NOVA_PORTA`
  - Atualizar a porta do Evans no **Makefile**
