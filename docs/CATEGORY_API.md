# API de Categorias Financeiras

## Endpoints

### 1. Criar Categoria
**POST** `/api/v1/category`

Cria uma nova categoria financeira.

**Request Body:**
```json
{
  "description": "Nome da Categoria",
  "type": "R"
}
```

**Campos:**
- `description` (string, obrigatório): Descrição da categoria
- `type` (string, obrigatório): Tipo da categoria
  - `"R"` para Receita
  - `"D"` para Despesa

**Response:**
- **201 Created**: Categoria criada com sucesso
- **400 Bad Request**: Dados inválidos

**Exemplo:**
```bash
curl -X POST http://localhost:8080/api/v1/category \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Venda de Produtos",
    "type": "R"
  }'
```

---

### 2. Listar Todas as Categorias
**GET** `/api/v1/category`

Retorna todas as categorias cadastradas.

**Response:**
- **200 OK**: Lista de categorias

**Exemplo:**
```bash
curl -X GET http://localhost:8080/api/v1/category
```

**Response Body:**
```json
[
  {
    "id": 1,
    "description": "Venda de Mercadorias",
    "type": "R"
  },
  {
    "id": 2,
    "description": "Combustível",
    "type": "D"
  }
]
```

---

### 3. Buscar Categorias por Tipo
**GET** `/api/v1/category?type=R` ou `/api/v1/category?type=D`

Retorna categorias filtradas por tipo (Receita ou Despesa).

**Query Parameters:**
- `type` (string): Tipo de categoria
  - `"R"` para Receitas
  - `"D"` para Despesas

**Response:**
- **200 OK**: Lista de categorias do tipo especificado
- **400 Bad Request**: Tipo inválido

**Exemplo - Buscar Receitas:**
```bash
curl -X GET "http://localhost:8080/api/v1/category?type=R"
```

**Exemplo - Buscar Despesas:**
```bash
curl -X GET "http://localhost:8080/api/v1/category?type=D"
```

**Response Body:**
```json
[
  {
    "id": 1,
    "description": "Venda de Mercadorias",
    "type": "R"
  },
  {
    "id": 3,
    "description": "Prestação de Serviços",
    "type": "R"
  }
]
```

---

### 4. Buscar Categoria por ID
**GET** `/api/v1/category/{id}`

Retorna uma categoria específica pelo ID.

**Path Parameters:**
- `id` (integer): ID da categoria

**Response:**
- **200 OK**: Categoria encontrada
- **404 Not Found**: Categoria não encontrada
- **400 Bad Request**: ID inválido

**Exemplo:**
```bash
curl -X GET http://localhost:8080/api/v1/category/1
```

**Response Body:**
```json
{
  "id": 1,
  "description": "Venda de Mercadorias",
  "type": "R"
}
```

---

### 5. Excluir Categoria
**DELETE** `/api/v1/category/{id}`

Exclui uma categoria pelo ID.

**Path Parameters:**
- `id` (integer): ID da categoria a ser excluída

**Response:**
- **200 OK**: Categoria excluída com sucesso
- **400 Bad Request**: ID inválido
- **500 Internal Server Error**: Erro ao excluir

**Exemplo:**
```bash
curl -X DELETE http://localhost:8080/api/v1/category/5
```

**Response Body:**
```json
{
  "message": "Category deleted successfully"
}
```

---

## Tipos de Categoria

- **R** - Receita: Entradas financeiras (vendas, serviços, juros recebidos, etc.)
- **D** - Despesa: Saídas financeiras (combustível, manutenção, salários, impostos, etc.)

## Códigos de Status HTTP

- **200 OK**: Requisição bem-sucedida
- **201 Created**: Recurso criado com sucesso
- **400 Bad Request**: Dados da requisição inválidos
- **404 Not Found**: Recurso não encontrado
- **500 Internal Server Error**: Erro no servidor

## Validações

- O campo `description` é obrigatório e não pode ser vazio
- O campo `type` deve ser exatamente `"R"` ou `"D"`
- O `id` deve ser um número inteiro positivo

## Banco de Dados

Execute o script de migração para criar a tabela:
```bash
mysql -u seu_usuario -p seu_banco < scripts/006_create_categoria_table.sql
```

Este script cria a tabela `categoria` e insere categorias padrão de exemplo.
