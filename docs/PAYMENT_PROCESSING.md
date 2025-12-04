# Processamento de Pagamentos - API Finance

## Descrição

Funcionalidade para processar pagamentos de lançamentos financeiros, permitindo registrar o valor pago, data de realização, forma de pagamento e opcionalmente lançar a diferença como um novo registro.

---

## Endpoint

### Processar Pagamento
**PUT** `/api/v1/finance/{id}/payment`

Processa o pagamento de um lançamento financeiro.

**Path Parameters:**
- `id` (integer): ID do lançamento financeiro

**Request Body:**
```json
{
  "valorPago": 500.00,
  "dataRealizacao": "2025-01-20T10:30:00Z",
  "formaPagamentoId": 2,
  "lancarDiferenca": true
}
```

**Campos:**
- `valorPago` (float64, obrigatório): Valor efetivamente pago
- `dataRealizacao` (datetime, obrigatório): Data em que o pagamento foi realizado
- `formaPagamentoId` (int64, obrigatório): ID da forma de pagamento utilizada
- `lancarDiferenca` (boolean, opcional): Se true, cria um novo lançamento com a diferença entre o valor original e o valor pago

**Response:**
- **200 OK**: Pagamento processado com sucesso
- **400 Bad Request**: Dados inválidos ou pagamento já processado
- **404 Not Found**: Lançamento não encontrado

---

## Exemplos de Uso

### Exemplo 1: Pagamento Exato
Pagamento do valor total, sem diferença:

**Request:**
```bash
curl -X PUT http://localhost:8080/api/v1/finance/123/payment \
  -H "Content-Type: application/json" \
  -d '{
    "valorPago": 500.00,
    "dataRealizacao": "2025-01-20T10:30:00Z",
    "formaPagamentoId": 2,
    "lancarDiferenca": false
  }'
```

**Resultado:**
- Lançamento ID 123 marcado como pago
- `valorPago`: 500.00
- `dataRealizacao`: 2025-01-20
- `formaPagamentoId`: 2 (PIX, por exemplo)
- `realizado`: true

---

### Exemplo 2: Pagamento com Desconto
Pagou menos que o valor original (desconto):

**Cenário:**
- Valor original: R$ 1.000,00
- Valor pago: R$ 950,00
- Diferença: -R$ 50,00 (desconto)

**Request:**
```bash
curl -X PUT http://localhost:8080/api/v1/finance/456/payment \
  -H "Content-Type: application/json" \
  -d '{
    "valorPago": 950.00,
    "dataRealizacao": "2025-01-20T10:30:00Z",
    "formaPagamentoId": 1,
    "lancarDiferenca": true
  }'
```

**Resultado:**
1. **Lançamento Original (ID 456):**
   - `valorPago`: 950.00
   - `realizado`: true
   
2. **Novo Lançamento (Ajuste):**
   - `valor`: 50.00
   - `origem`: "ajuste_pagamento"
   - `origemId`: 456
   - `numeroDocumento`: "{original}-AJUSTE"
   - `observacao`: "Ajuste de pagamento - Ref: {doc_original} (Diferença: -50.00) - Desconto"
   - `realizado`: true

---

### Exemplo 3: Pagamento com Acréscimo
Pagou mais que o valor original (juros, multa, etc):

**Cenário:**
- Valor original: R$ 1.000,00
- Valor pago: R$ 1.050,00
- Diferença: +R$ 50,00 (acréscimo)

**Request:**
```bash
curl -X PUT http://localhost:8080/api/v1/finance/789/payment \
  -H "Content-Type: application/json" \
  -d '{
    "valorPago": 1050.00,
    "dataRealizacao": "2025-01-20T10:30:00Z",
    "formaPagamentoId": 3,
    "lancarDiferenca": true
  }'
```

**Resultado:**
1. **Lançamento Original (ID 789):**
   - `valorPago`: 1050.00
   - `realizado`: true
   
2. **Novo Lançamento (Ajuste):**
   - `valor`: 50.00
   - `origem`: "ajuste_pagamento"
   - `origemId`: 789
   - `numeroDocumento`: "{original}-AJUSTE"
   - `observacao`: "Ajuste de pagamento - Ref: {doc_original} (Diferença: 50.00) - Acréscimo"
   - `realizado`: true

---

## Campos Adicionados à Entidade Finance

```go
type Finance struct {
    // ... campos existentes ...
    
    FormaPagamentoId *int64     `json:"formaPagamentoId"` // ID da forma de pagamento (null se não pago)
    ValorPago        *float64   `json:"valorPago"`        // Valor efetivamente pago (null se não pago)
    Realizado        bool       `json:"realizado"`        // Se o pagamento foi realizado
}
```

---

## Validações

1. **Lançamento deve existir**: O ID fornecido deve corresponder a um lançamento existente
2. **Não pode estar pago**: O lançamento não pode ter sido processado anteriormente (`realizado` = false)
3. **ValorPago deve ser maior que zero**: Não é permitido valor zero ou negativo
4. **FormaPagamentoId é obrigatório**: Deve ser um ID válido de forma de pagamento

---

## Regras de Negócio

### Processamento do Pagamento Principal
1. Atualiza o lançamento original com:
   - `valorPago` = valor informado
   - `dataRealizacao` = data informada
   - `formaPagamentoId` = forma de pagamento informada
   - `realizado` = true

### Lançamento de Diferença (quando `lancarDiferenca` = true)

**Se houver diferença:**
- Calcula: `diferenca = valorPago - valorOriginal`

**Se diferença > 0 (pagou mais):**
- Cria lançamento de ACRÉSCIMO
- Mesmo tipo de categoria do original
- Valor positivo

**Se diferença < 0 (pagou menos):**
- Cria lançamento de DESCONTO
- Valor convertido para positivo
- Observação indica desconto

**Campos do lançamento de ajuste:**
- `pessoaId`: mesmo do original
- `categoriaId`: mesmo do original
- `origemId`: ID do lançamento original
- `origem`: "ajuste_pagamento"
- `valor`: valor absoluto da diferença
- `numeroDocumento`: "{documento_original}-AJUSTE"
- `dataCompetencia`: data da realização
- `dataVencimento`: data da realização
- `dataRealizacao`: data da realização
- `formaPagamentoId`: mesma do pagamento
- `realizado`: true
- `observacao`: descrição detalhada com referência e tipo

---

## Mensagens de Erro

### Lançamento não encontrado:
```json
{
  "error": "finance record not found"
}
```

### Pagamento já processado:
```json
{
  "error": "payment already processed"
}
```

### Valor pago inválido:
```json
{
  "error": "valorPago must be greater than zero"
}
```

### Forma de pagamento não informada:
```json
{
  "error": "formaPagamentoId is required"
}
```

### ID inválido:
```json
{
  "error": "Invalid ID"
}
```

---

## Migração do Banco de Dados

Execute o script para adicionar as novas colunas:
```bash
mysql -u seu_usuario -p seu_banco < scripts/008_add_payment_fields_to_financeiro.sql
```

**Colunas adicionadas:**
- `forma_pagamento_id` (BIGINT, NULL) - FK para forma_pagamento
- `valor_pago` (DECIMAL(10,2), NULL) - Valor efetivamente pago
- `realizado` (BOOLEAN, DEFAULT FALSE) - Status do pagamento

**Índices criados:**
- `idx_financeiro_realizado` - Para consultas por status
- `idx_financeiro_forma_pagamento` - Para consultas por forma de pagamento

---

## Fluxo Completo

```
1. Cliente envia PUT /api/v1/finance/{id}/payment
   ↓
2. Controller valida o ID e decodifica o body
   ↓
3. Service ProcessPayment:
   a. Busca o lançamento original
   b. Valida se não foi pago
   c. Valida os dados (valorPago, formaPagamentoId)
   d. Atualiza o lançamento com os dados do pagamento
   e. Se lancarDiferenca = true E houver diferença:
      - Calcula a diferença
      - Cria novo lançamento de ajuste
   ↓
4. Retorna sucesso ou erro
```

---

## DTO Utilizado

```go
type PaymentRequest struct {
    ValorPago         float64   `json:"valorPago"`
    DataRealizacao    time.Time `json:"dataRealizacao"`
    FormaPagamentoId  int64     `json:"formaPagamentoId"`
    LancarDiferenca   bool      `json:"lancarDiferenca"`
}
```

---

## Considerações

1. **Transações**: Atualmente o processamento não é atômico. Se houver erro ao criar o lançamento de diferença, o pagamento principal já terá sido registrado. Considere implementar transações.

2. **Reversão**: Não há funcionalidade para reverter um pagamento. Uma vez processado, o status `realizado` permanece true.

3. **Auditoria**: O campo `origem` = "ajuste_pagamento" e `origemId` permitem rastrear ajustes relacionados ao lançamento original.

4. **Categoria**: Os lançamentos de ajuste herdam a mesma categoria do lançamento original, mas isso pode ser customizado se necessário.
