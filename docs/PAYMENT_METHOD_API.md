# API de Formas de Pagamento (Payment Methods)

## Descrição

Módulo para gerenciamento de formas de pagamento (payment methods) com operações CRUD completas.

---

## Endpoints

### 1. Criar Forma de Pagamento
**POST** `/api/v1/payment-method`

Cria uma nova forma de pagamento.

**Request Body:**
```json
{
  "description": "PIX"
}
```

**Campos:**
- `description` (string, obrigatório): Descrição da forma de pagamento

**Response:**
- **201 Created**: Forma de pagamento criada com sucesso
- **400 Bad Request**: Dados inválidos

**Exemplo:**
```bash
curl -X POST http://localhost:8080/api/v1/payment-method \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Cartão de Crédito"
  }'
```

---

### 2. Listar Todas as Formas de Pagamento
**GET** `/api/v1/payment-method`

Retorna todas as formas de pagamento cadastradas, ordenadas por descrição.

**Response:**
- **200 OK**: Lista de formas de pagamento

**Exemplo:**
```bash
curl -X GET http://localhost:8080/api/v1/payment-method
```

**Response Body:**
```json
[
  {
    "id": 1,
    "description": "Dinheiro"
  },
  {
    "id": 2,
    "description": "PIX"
  },
  {
    "id": 3,
    "description": "Cartão de Crédito"
  }
]
```

---

### 3. Buscar Forma de Pagamento por ID
**GET** `/api/v1/payment-method/{id}`

Retorna uma forma de pagamento específica pelo ID.

**Path Parameters:**
- `id` (integer): ID da forma de pagamento

**Response:**
- **200 OK**: Forma de pagamento encontrada
- **404 Not Found**: Forma de pagamento não encontrada
- **400 Bad Request**: ID inválido

**Exemplo:**
```bash
curl -X GET http://localhost:8080/api/v1/payment-method/1
```

**Response Body:**
```json
{
  "id": 1,
  "description": "Dinheiro"
}
```

---

### 4. Atualizar Forma de Pagamento
**PUT** `/api/v1/payment-method/{id}`

Atualiza uma forma de pagamento existente.

**Path Parameters:**
- `id` (integer): ID da forma de pagamento

**Request Body:**
```json
{
  "description": "Cartão de Crédito Visa"
}
```

**Response:**
- **200 OK**: Forma de pagamento atualizada com sucesso
- **400 Bad Request**: Dados inválidos
- **404 Not Found**: Forma de pagamento não encontrada

**Exemplo:**
```bash
curl -X PUT http://localhost:8080/api/v1/payment-method/3 \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Cartão de Crédito Mastercard"
  }'
```

---

### 5. Excluir Forma de Pagamento
**DELETE** `/api/v1/payment-method/{id}`

Exclui uma forma de pagamento pelo ID.

**Path Parameters:**
- `id` (integer): ID da forma de pagamento a ser excluída

**Response:**
- **200 OK**: Forma de pagamento excluída com sucesso
- **400 Bad Request**: ID inválido
- **404 Not Found**: Forma de pagamento não encontrada
- **500 Internal Server Error**: Erro ao excluir

**Exemplo:**
```bash
curl -X DELETE http://localhost:8080/api/v1/payment-method/5
```

**Response Body:**
```json
{
  "message": "Payment method deleted successfully"
}
```

---

## Estrutura da Entidade

```go
type PaymentMethod struct {
    Id          int64  `json:"id"`
    Name string `json:"name"`
}
```

---

## Validações

- O campo `description` é obrigatório e não pode ser vazio
- O `id` deve ser um número inteiro positivo (para operações que requerem ID)

---

## Banco de Dados

### Tabela: `forma_pagamento`

| Coluna      | Tipo         | Descrição                      |
|-------------|--------------|--------------------------------|
| id          | BIGINT       | Chave primária (auto increment)|
| descricao   | VARCHAR(255) | Descrição da forma de pagamento|
| created_at  | TIMESTAMP    | Data de criação                |
| updated_at  | TIMESTAMP    | Data de atualização            |

### Script de Migração

Execute o script para criar a tabela:
```bash
mysql -u seu_usuario -p seu_banco < scripts/007_create_payment_method_table.sql
```

Este script cria a tabela `forma_pagamento` e insere formas de pagamento padrão:
- Dinheiro
- PIX
- Cartão de Crédito
- Cartão de Débito
- Boleto Bancário
- Transferência Bancária
- Cheque

---

## Códigos de Status HTTP

- **200 OK**: Requisição bem-sucedida
- **201 Created**: Recurso criado com sucesso
- **400 Bad Request**: Dados da requisição inválidos
- **404 Not Found**: Recurso não encontrado
- **500 Internal Server Error**: Erro no servidor

---

## Exemplos de Uso Completo

### Criar uma nova forma de pagamento
```bash
curl -X POST http://localhost:8080/api/v1/payment-method \
  -H "Content-Type: application/json" \
  -d '{"description": "Criptomoeda"}'
```

### Listar todas as formas
```bash
curl -X GET http://localhost:8080/api/v1/payment-method
```

### Buscar uma forma específica
```bash
curl -X GET http://localhost:8080/api/v1/payment-method/2
```

### Atualizar uma forma
```bash
curl -X PUT http://localhost:8080/api/v1/payment-method/2 \
  -H "Content-Type: application/json" \
  -d '{"description": "PIX QR Code"}'
```

### Excluir uma forma
```bash
curl -X DELETE http://localhost:8080/api/v1/payment-method/8
```

---

## Mensagens de Erro

### Descrição vazia:
```json
{
  "error": "description is required"
}
```

### ID inválido:
```json
{
  "error": "invalid id"
}
```

### Forma de pagamento não encontrada:
```json
{
  "error": "Payment method not found"
}
```

---

## Arquivos do Módulo

- **Entidade**: `internal/domain/entities/payment_method.go`
- **Repository Interface**: `internal/domain/contract/repository/repository.go`
- **Repository Implementation**: `internal/infrastructure/database/repository/payment_method_repository.go`
- **Service**: `internal/services/payment_method_service.go`
- **Controller**: `api/controllers/payment_method_controller.go`
- **Routes**: `api/routes/payment_method_routes.go`
- **Migration**: `scripts/007_create_payment_method_table.sql`
