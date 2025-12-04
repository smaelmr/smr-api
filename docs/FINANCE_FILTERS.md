# API de Lançamentos Financeiros - Filtros por Tipo, Mês e Ano

## Alterações Realizadas

Os métodos `GetAll`, `GetReceipts` e `GetPayments` foram atualizados para aceitar parâmetros de filtro por **tipo de categoria**, **mês** e **ano** da data de lançamento.

---

## Endpoints Atualizados

### 1. Buscar Todos os Lançamentos (com filtros)
**GET** `/api/v1/finance?type={R|D}&month={1-12}&year={YYYY}`

Retorna todos os lançamentos financeiros filtrados por tipo de categoria, mês e ano.

**Query Parameters (todos obrigatórios):**
- `type` (string): Tipo de categoria
  - `"R"` para Receitas
  - `"D"` para Despesas
- `month` (integer): Mês do lançamento (1-12)
- `year` (integer): Ano do lançamento (> 1900)

**Response:**
- **200 OK**: Lista de lançamentos
- **400 Bad Request**: Parâmetros inválidos ou faltando

**Exemplo - Buscar Despesas de Janeiro/2025:**
```bash
curl -X GET "http://localhost:8080/api/v1/finance?type=D&month=1&year=2025"
```

**Exemplo - Buscar Receitas de Dezembro/2024:**
```bash
curl -X GET "http://localhost:8080/api/v1/finance?type=R&month=12&year=2024"
```

**Response Body:**
```json
[
  {
    "id": 1,
    "pessoaId": 10,
    "categoriaId": 5,
    "lancamentoId": 100,
    "lancamentoTipo": "abastecimento",
    "valor": 500.50,
    "numeroParcela": 1,
    "numeroDocumento": "DOC-001",
    "dataLancamento": "2025-01-15T10:30:00Z",
    "dataVencimento": "2025-01-20T00:00:00Z",
    "dataRealizacao": "2025-01-18T00:00:00Z",
    "observacao": "Combustível",
    "realizado": true,
    "createdAt": "2025-01-15T10:30:00Z",
    "updatedAt": "2025-01-15T10:30:00Z"
  }
]
```

---

### 2. Buscar Receitas
**GET** `/api/v1/finance/receipts?month={1-12}&year={YYYY}`

Retorna receitas (tipo R) filtradas por mês e ano.

**Query Parameters (todos obrigatórios):**
- `month` (integer): Mês do lançamento (1-12)
- `year` (integer): Ano do lançamento (> 1900)

**Response:**
- **200 OK**: Lista de receitas
- **400 Bad Request**: Parâmetros inválidos ou faltando

**Exemplo:**
```bash
curl -X GET "http://localhost:8080/api/v1/finance/receipts?month=3&year=2025"
```

---

### 3. Buscar Despesas
**GET** `/api/v1/finance/payments?month={1-12}&year={YYYY}`

Retorna despesas (tipo D) filtradas por mês e ano.

**Query Parameters (todos obrigatórios):**
- `month` (integer): Mês do lançamento (1-12)
- `year` (integer): Ano do lançamento (> 1900)

**Response:**
- **200 OK**: Lista de despesas
- **400 Bad Request**: Parâmetros inválidos ou faltando

**Exemplo:**
```bash
curl -X GET "http://localhost:8080/api/v1/finance/payments?month=3&year=2025"
```

---

## Validações Implementadas

### No Service Layer:
- **Tipo**: Deve ser exatamente `"R"` ou `"D"`
- **Mês**: Deve estar entre 1 e 12
- **Ano**: Deve ser maior que 1900

### No Controller Layer:
- Todos os parâmetros são obrigatórios
- Validação de formato numérico para mês e ano
- Mensagens de erro descritivas

---

## Query SQL Executada

A query no repositório foi atualizada para:

```sql
SELECT 
    f.id, f.pessoa_id, f.valor, f.numero_documento, f.data_lancamento, 
    f.data_vencimento, f.data_realizacao, f.origem, f.origem_id, f.observacao, 
    f.realizado, f.created_at, f.updated_at
FROM financeiro f
INNER JOIN categoria c ON f.categoria_id = c.id
WHERE c.tipo = ?
    AND MONTH(f.data_lancamento) = ?
    AND YEAR(f.data_lancamento) = ?
ORDER BY f.data_lancamento DESC
```

**Nota:** A query faz `INNER JOIN` com a tabela `categoria` para filtrar pelo tipo, portanto é necessário que:
1. A tabela `categoria` exista no banco de dados
2. O campo `categoria_id` na tabela `financeiro` esteja preenchido e seja uma FK válida

---

## Mensagens de Erro

### Parâmetros Faltando:
```json
{
  "error": "type, month and year are required"
}
```

### Mês Inválido:
```json
{
  "error": "invalid month"
}
```

ou

```json
{
  "error": "month must be between 1 and 12"
}
```

### Ano Inválido:
```json
{
  "error": "invalid year"
}
```

ou

```json
{
  "error": "year must be greater than 1900"
}
```

### Tipo Inválido:
```json
{
  "error": "type must be 'R' (receita) or 'D' (despesa)"
}
```

---

## Exemplos de Uso

### Buscar todas as despesas de Março/2025:
```bash
curl -X GET "http://localhost:8080/api/v1/finance?type=D&month=3&year=2025"
```

### Buscar todas as receitas de Dezembro/2024:
```bash
curl -X GET "http://localhost:8080/api/v1/finance?type=R&month=12&year=2024"
```

### Usar endpoint específico para despesas:
```bash
curl -X GET "http://localhost:8080/api/v1/finance/payments?month=6&year=2025"
```

### Usar endpoint específico para receitas:
```bash
curl -X GET "http://localhost:8080/api/v1/finance/receipts?month=6&year=2025"
```

---

## Requisitos do Banco de Dados

Para que os filtros funcionem corretamente, certifique-se de que:

1. A tabela `categoria` existe e contém categorias com tipos 'R' ou 'D'
2. A tabela `financeiro` possui o campo `categoria_id` como chave estrangeira
3. Execute o script de migração se necessário: `scripts/006_create_categoria_table.sql`

---

## Alterações nos Arquivos

### Arquivos Modificados:
1. `internal/domain/contract/repository/repository.go` - Interface `FinanceRepository`
2. `internal/infrastructure/database/repository/finance_repository.go` - Implementação do `GetAll`
3. `internal/services/finance_service.go` - Métodos `GetAll`, `GetReceipts` e `GetPayments`
4. `api/controllers/finance_controller.go` - Handlers com validação de parâmetros
