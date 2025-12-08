# Status Dinâmico de Lançamentos Financeiros

## Descrição

O campo `status` é uma propriedade dinâmica calculada automaticamente para cada lançamento financeiro retornado pela API. Este campo não é armazenado no banco de dados, mas é gerado em tempo de execução baseado em regras de negócio.

---

## Regras de Status

### 1. **Pago/Recebido**
Quando o lançamento possui `dataRealizacao` preenchida (não null).

**Condição:**
```
dataRealizacao != null
```

**Exemplo:**
```json
{
  "id": 123,
  "valorParcela": 1000.00,
  "dataVencimento": "2025-01-15T00:00:00Z",
  "dataRealizacao": "2025-01-10T10:30:00Z",
  "status": "Pago/Recebido"
}
```

---

### 2. **Em Atraso**
Quando o lançamento não possui `dataRealizacao` e a `dataVencimento` já passou.

**Condição:**
```
dataRealizacao == null AND dataVencimento < hoje
```

**Exemplo:**
```json
{
  "id": 124,
  "valorParcela": 1000.00,
  "dataVencimento": "2024-12-01T00:00:00Z",
  "dataRealizacao": null,
  "status": "Em Atraso"
}
```

---

### 3. **Em Aberto**
Quando o lançamento não possui `dataRealizacao` e a `dataVencimento` ainda não passou.

**Condição:**
```
dataRealizacao == null AND dataVencimento >= hoje
```

**Exemplo:**
```json
{
  "id": 125,
  "valorParcela": 1000.00,
  "dataVencimento": "2025-02-15T00:00:00Z",
  "dataRealizacao": null,
  "status": "EM_ABERTO"
}
```

---

## Implementação Técnica

### Método GetStatus()
```go
func (f *Finance) GetStatus() string {
    // Se tem data de realização, está pago/recebido
    if f.DataRealizacao != nil {
        return "Pago/Recebido"
    }

    // Se não tem data de realização, verificar vencimento
    hoje := time.Now()
    
    // Normalizar as datas para comparação apenas de dia (sem hora)
    hojeNormalizado := time.Date(hoje.Year(), hoje.Month(), hoje.Day(), 0, 0, 0, 0, hoje.Location())
    vencimentoNormalizado := time.Date(f.DataVencimento.Year(), f.DataVencimento.Month(), f.DataVencimento.Day(), 0, 0, 0, 0, f.DataVencimento.Location())

    // Se já passou do vencimento
    if hojeNormalizado.After(vencimentoNormalizado) {
        return "Em Atraso"
    }

    // Caso contrário, está em aberto dentro do prazo
    return "Em Aberto"
}
```

### Serialização JSON Customizada
A struct `Finance` implementa o método `MarshalJSON()` para incluir automaticamente o campo `status` na resposta JSON:

```go
func (f *Finance) MarshalJSON() ([]byte, error) {
    type Alias Finance
    return json.Marshal(&struct {
        *Alias
        Status string `json:"status"`
    }{
        Alias:  (*Alias)(f),
        Status: f.GetStatus(),
    })
}
```

---

## Uso na API

O campo `status` é automaticamente incluído em **todas** as respostas que retornam objetos `Finance`:

### GET /api/v1/finance?type=D&month=1&year=2025
```json
[
  {
    "id": 1,
    "pessoaId": 10,
    "categoriaId": 5,
    "valorParcela": 500.00,
    "dataVencimento": "2025-01-15T00:00:00Z",
    "dataRealizacao": null,
    "status": "Em Aberto"
  },
  {
    "id": 2,
    "pessoaId": 12,
    "categoriaId": 5,
    "valorParcela": 800.00,
    "dataVencimento": "2024-12-20T00:00:00Z",
    "dataRealizacao": null,
    "status": "Em Atraso"
  },
  {
    "id": 3,
    "pessoaId": 15,
    "categoriaId": 3,
    "valorParcela": 1200.00,
    "dataVencimento": "2025-01-10T00:00:00Z",
    "dataRealizacao": "2025-01-10T14:30:00Z",
    "status": "Pago/Recebido"
  }
]
```

### GET /api/v1/finance/receipts?month=1&year=2025
```json
[
  {
    "id": 4,
    "pessoaId": 20,
    "categoriaId": 1,
    "valorParcela": 2500.00,
    "dataVencimento": "2025-01-25T00:00:00Z",
    "dataRealizacao": null,
    "status": "Em Aberto"
  }
]
```

### GET /api/v1/finance/payments?month=1&year=2025
Mesmo comportamento para pagamentos.

---

## Vantagens da Implementação

1. **Dinâmico**: O status é sempre calculado em tempo real, refletindo o estado atual
2. **Automático**: Não precisa ser calculado manualmente no frontend
3. **Consistente**: A mesma lógica é aplicada em todos os endpoints
4. **Manutenível**: Centralizado na entidade `Finance`
5. **Sem persistência**: Não ocupa espaço no banco de dados
6. **Performático**: Cálculo simples e rápido

---

## Comparação de Datas

A comparação de datas é feita **normalizando para o início do dia** (00:00:00), ou seja:
- Ignora horas, minutos e segundos
- Compara apenas dia, mês e ano

**Exemplo:**
- Hoje: `2025-01-15 14:30:45`
- Vencimento: `2025-01-15 00:00:00`
- Resultado: **Em Aberto** (mesmo dia)

- Hoje: `2025-01-16 00:00:01`
- Vencimento: `2025-01-15 23:59:59`
- Resultado: **Em Atraso** (dia seguinte)

---

## Filtros no Frontend

Com o campo `status`, o frontend pode facilmente:

### Filtrar por status
```javascript
// Apenas em aberto
const emAberto = lancamentos.filter(l => l.status === "Em Aberto");

// Apenas em atraso
const emAtraso = lancamentos.filter(l => l.status === "Em Atraso");

// Apenas pagos
const pagos = lancamentos.filter(l => l.status === "Pago/Recebido");
```

### Exibir badges/cores
```javascript
function getStatusColor(status) {
  switch(status) {
    case "Em Aberto": return "blue";
    case "Em Atraso": return "red";
    case "Pago/Recebido": return "green";
    default: return "gray";
  }
}
```

### Contadores
```javascript
const totalEmAberto = lancamentos.filter(l => l.status === "Em Aberto").length;
const totalEmAtraso = lancamentos.filter(l => l.status === "Em Atraso").length;
const totalPagos = lancamentos.filter(l => l.status === "Pago/Recebido").length;
```

---

## Notas Importantes

1. **Somente Leitura**: O campo `status` não pode ser enviado em requisições POST/PUT - ele é calculado
2. **Timezone**: A comparação de datas usa o timezone local do servidor
3. **Cache**: Se você usar cache, lembre-se que o status pode mudar quando a data passar
4. **Consultas**: Para consultas complexas baseadas em status, considere adicionar índices em `dataRealizacao` e `dataVencimento`

---

## Exemplos de Cenários

### Cenário 1: Boleto com vencimento futuro
```json
{
  "dataVencimento": "2025-02-01",
  "dataRealizacao": null,
  "status": "Em Aberto"
}
```

### Cenário 2: Boleto vencido há 5 dias
```json
{
  "dataVencimento": "2025-01-02",
  "dataRealizacao": null,
  "status": "Em Atraso"
}
```

### Cenário 3: Boleto pago antes do vencimento
```json
{
  "dataVencimento": "2025-02-01",
  "dataRealizacao": "2025-01-25T10:30:00Z",
  "status": "Pago/Recebido"
}
```

### Cenário 4: Boleto pago após o vencimento
```json
{
  "dataVencimento": "2025-01-01",
  "dataRealizacao": "2025-01-10T15:00:00Z",
  "status": "Pago/Recebido"
}
```
*Nota: Mesmo pago com atraso, o status é "Pago/Recebido" porque possui dataRealizacao*
