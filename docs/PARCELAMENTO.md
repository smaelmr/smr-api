# Funcionalidade de Parcelamento de Lançamentos Financeiros

## Descrição

O sistema agora suporta o parcelamento automático de lançamentos financeiros. Quando um lançamento é criado com mais de uma parcela, o sistema automaticamente cria múltiplos registros no banco de dados, um para cada parcela.

---

## Como Funciona

### Campo Adicionado
- **`totalParcelas`** (int32): Indica o número total de parcelas desejado

### Comportamento

1. **Parcela Única (totalParcelas <= 1)**
   - Cria apenas 1 registro
   - `numeroParcela` é definido como 1
   - `totalParcelas` é definido como 1

2. **Múltiplas Parcelas (totalParcelas > 1)**
   - Cria N registros (onde N = totalParcelas)
   - Cada registro representa uma parcela
   - O valor total é dividido igualmente entre as parcelas
   - A última parcela recebe ajuste para compensar arredondamentos
   - O vencimento é incrementado mensalmente para cada parcela
   - O número do documento recebe o sufixo "-X/N" (X = parcela atual, N = total)

---

## Lógica de Divisão

### Valor das Parcelas
- **Parcelas 1 a N-1**: `valor_total / total_parcelas`
- **Última Parcela (N)**: `valor_total - soma_parcelas_anteriores`
  - Isso garante que não haja perda de centavos por arredondamento

**Exemplo:**
- Valor total: R$ 1.000,00
- Total de parcelas: 3
- Parcela 1: R$ 333,33
- Parcela 2: R$ 333,33
- Parcela 3: R$ 333,34 (ajustada)

### Data de Vencimento
Cada parcela tem o vencimento incrementado mensalmente:
- Parcela 1: Data informada
- Parcela 2: Data informada + 1 mês
- Parcela 3: Data informada + 2 meses
- E assim por diante...

### Número do Documento
Se informado, recebe o sufixo com o número da parcela:
- Formato: `{numeroDocumento}-{parcela}/{total}`
- Exemplo: `DOC-001-1/3`, `DOC-001-2/3`, `DOC-001-3/3`

---

## Exemplos de Uso

### Exemplo 1: Parcelamento em 3x

**Request Body:**
```json
{
  "pessoaId": 10,
  "categoriaId": 5,
  "valor": 1500.00,
  "totalParcelas": 3,
  "numeroDocumento": "NF-12345",
  "dataLancamento": "2025-01-15T10:00:00Z",
  "dataVencimento": "2025-02-01T00:00:00Z",
  "origem": "manual",
  "observacao": "Compra parcelada",
  "realizado": false
}
```

**Registros Criados:**

| Parcela | Valor    | Vencimento | Documento       |
|---------|----------|------------|-----------------|
| 1       | 500.00   | 2025-02-01 | NF-12345-1/3   |
| 2       | 500.00   | 2025-03-01 | NF-12345-2/3   |
| 3       | 500.00   | 2025-04-01 | NF-12345-3/3   |

---

### Exemplo 2: Parcelamento com Arredondamento

**Request Body:**
```json
{
  "pessoaId": 15,
  "categoriaId": 8,
  "valor": 100.00,
  "totalParcelas": 3,
  "numeroDocumento": "BOL-789",
  "dataLancamento": "2025-01-10T10:00:00Z",
  "dataVencimento": "2025-01-20T00:00:00Z",
  "origem": "manual",
  "realizado": false
}
```

**Registros Criados:**

| Parcela | Valor    | Vencimento | Documento     |
|---------|----------|------------|---------------|
| 1       | 33.33    | 2025-01-20 | BOL-789-1/3  |
| 2       | 33.33    | 2025-02-20 | BOL-789-2/3  |
| 3       | 33.34    | 2025-03-20 | BOL-789-3/3  |

*Nota: A última parcela recebe R$ 33,34 para compensar o arredondamento*

---

### Exemplo 3: Sem Parcelamento (Padrão)

**Request Body:**
```json
{
  "pessoaId": 20,
  "categoriaId": 3,
  "valor": 250.00,
  "numeroDocumento": "DOC-001",
  "dataLancamento": "2025-01-15T10:00:00Z",
  "dataVencimento": "2025-01-30T00:00:00Z",
  "origem": "manual",
  "realizado": false
}
```

*Nota: Se `totalParcelas` não for informado ou for 1, cria apenas 1 registro*

**Registro Criado:**

| Parcela | Valor  | Vencimento | Documento |
|---------|--------|------------|-----------|
| 1       | 250.00 | 2025-01-30 | DOC-001   |

---

## Campos da Entidade Finance

```go
type Finance struct {
    Id              int64      `json:"id"`
    PessoaId        int64      `json:"pessoaId"`
    CategoriaId     int64      `json:"categoriaId"`
    OrigemId        *int64     `json:"lancamentoId"`
    Origem          string     `json:"lancamentoTipo"`
    Valor           float64    `json:"valor"`              // Valor da parcela
    NumeroParcela   int32      `json:"numeroParcela"`      // Número desta parcela (1, 2, 3...)
    TotalParcelas   int32      `json:"totalParcelas"`      // Total de parcelas
    NumeroDocumento string     `json:"numeroDocumento"`
    DataLancamento  time.Time  `json:"dataLancamento"`
    DataVencimento  time.Time  `json:"dataVencimento"`     // Vencimento incrementado mensalmente
    DataRealizacao  *time.Time `json:"dataRealizacao"`
    Observacao      string     `json:"observacao"`
    Realizado       bool       `json:"realizado"`
    CreatedAt       time.Time  `json:"createdAt"`
    UpdatedAt       time.Time  `json:"updatedAt"`
}
```

---

## Regras de Negócio

1. **totalParcelas <= 1**: Cria apenas 1 registro (comportamento padrão)
2. **totalParcelas > 1**: Cria N registros, onde N = totalParcelas
3. **Valor**: Dividido igualmente, com ajuste na última parcela
4. **Vencimento**: Incrementado mensalmente (mês a mês)
5. **Número do Documento**: Formatado como `{original}-{parcela}/{total}`
6. **NumeroParcela**: Vai de 1 até totalParcelas
7. **Todos os campos**: São replicados em todas as parcelas, exceto:
   - `valor` (dividido)
   - `numeroParcela` (incrementado)
   - `dataVencimento` (incrementado mensalmente)
   - `numeroDocumento` (com sufixo da parcela)

---

## Endpoint

**POST** `/api/v1/finance`

O endpoint permanece o mesmo, apenas adicione o campo `totalParcelas` no JSON.

---

## Considerações Importantes

1. **Transações**: Atualmente, se houver erro ao criar uma das parcelas, as anteriores já terão sido salvas. Considere implementar transações para garantir atomicidade.

2. **Performance**: Para muitas parcelas (ex: 12x, 24x), serão feitas múltiplas inserções no banco. Isso é aceitável para operações OLTP normais.

3. **Edição**: Ao editar um lançamento parcelado, você editará apenas uma parcela por vez. Não há funcionalidade automática para editar todas as parcelas de uma vez.

4. **Exclusão**: Similar à edição, a exclusão afeta apenas uma parcela. Se quiser excluir todas, será necessário excluir cada registro individualmente.

---

## Melhorias Futuras

- [ ] Implementar transações para garantir que todas as parcelas sejam criadas ou nenhuma
- [ ] Adicionar endpoint para buscar todas as parcelas de um mesmo lançamento
- [ ] Permitir edição em lote de todas as parcelas
- [ ] Adicionar campo para agrupar parcelas do mesmo lançamento
- [ ] Permitir diferentes intervalos (quinzenal, semanal) além de mensal
