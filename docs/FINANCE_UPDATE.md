# Atualização de Lançamentos Financeiros

## Endpoint

**PUT** `/api/v1/finance/{id}`

Atualiza um lançamento financeiro existente.

## Parâmetros

**Path Parameters:**
- `id` (integer, obrigatório): ID do lançamento financeiro

**Request Body:**
```json
{
  "pessoaId": 10,
  "categoriaId": 5,
  "origemId": null,
  "origem": "manual",
  "valor": 1500.00,
  "valorParcela": 1500.00,
  "numeroParcela": 1,
  "totalParcelas": 1,
  "numeroDocumento": "NF-12345",
  "dataCompetencia": "2025-01-15T00:00:00Z",
  "dataVencimento": "2025-02-01T00:00:00Z",
  "dataRealizacao": null,
  "observacao": "Lançamento atualizado",
  "formaPagamentoId": null,
  "valorPago": null
}
```

## Validações

- **ID válido**: O ID deve ser maior que zero
- **PessoaId obrigatório**: Deve ser maior que zero
- **CategoriaId obrigatório**: Deve ser maior que zero
- **Valor obrigatório**: Valor ou ValorParcela deve ser maior que zero
- **Registro deve existir**: O lançamento com o ID informado deve existir no banco

## Respostas

### Sucesso (200 OK)
```json
{
  "message": "Finance record updated successfully"
}
```

### Erro - ID Inválido (400 Bad Request)
```json
{
  "error": "Invalid ID"
}
```

### Erro - Body Inválido (400 Bad Request)
```json
{
  "error": "Invalid request body"
}
```

### Erro - Validação (400 Bad Request)
```json
{
  "error": "pessoaId is required"
}
```

ou

```json
{
  "error": "categoriaId is required"
}
```

ou

```json
{
  "error": "valor or valorParcela must be greater than zero"
}
```

### Erro - Registro Não Encontrado (400 Bad Request)
```json
{
  "error": "finance record not found"
}
```

## Exemplos de Uso

### Atualizar um lançamento
```bash
curl -X PUT http://localhost:8080/api/v1/finance/123 \
  -H "Content-Type: application/json" \
  -d '{
    "pessoaId": 10,
    "categoriaId": 5,
    "origem": "manual",
    "valor": 2000.00,
    "valorParcela": 2000.00,
    "numeroParcela": 1,
    "totalParcelas": 1,
    "numeroDocumento": "NF-12345-ALT",
    "dataCompetencia": "2025-01-15T00:00:00Z",
    "dataVencimento": "2025-02-15T00:00:00Z",
    "observacao": "Valor atualizado"
  }'
```

### Atualizar apenas alguns campos
```bash
curl -X PUT http://localhost:8080/api/v1/finance/123 \
  -H "Content-Type: application/json" \
  -d '{
    "pessoaId": 10,
    "categoriaId": 5,
    "origem": "manual",
    "valorParcela": 1800.00,
    "dataVencimento": "2025-02-20T00:00:00Z",
    "observacao": "Data de vencimento alterada"
  }'
```

## Notas Importantes

1. **ID na URL prevalece**: O ID informado na URL sempre será usado, mesmo que um ID diferente seja enviado no body
2. **Campos obrigatórios**: PessoaId, CategoriaId e Valor/ValorParcela são sempre obrigatórios
3. **Campos nullable**: OrigemId, FormaPagamentoId, ValorPago e DataRealizacao podem ser null
4. **Validação de existência**: O sistema verifica se o registro existe antes de atualizar
5. **Não altera pagamentos**: Para processar pagamentos, use o endpoint específico `PUT /api/v1/finance/{id}/payment`

## Diferença entre Update e ProcessPayment

- **PUT /finance/{id}**: Atualiza dados do lançamento (valor, vencimento, pessoa, etc)
- **PUT /finance/{id}/payment**: Processa o pagamento (marca como pago, registra valor pago, forma de pagamento)

Use o endpoint apropriado para cada caso!
