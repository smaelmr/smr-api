# Campos ValorArla e ValorDiversos no Abastecimento

## Descrição

Foram adicionados dois novos campos ao cadastro de abastecimento:
- **valorArla**: Salvo no banco de dados e editável
- **valorDiversos**: Usado apenas para cálculo do lançamento financeiro (não salvo no banco)

---

## valorArla (Salvo no Banco)

Campo que armazena o valor gasto com Arla no abastecimento.

### Características:
- **Tipo**: `DECIMAL(10,2)`
- **Obrigatório**: Sim (default 0.00)
- **Editável**: Sim
- **Salvo no banco**: Sim
- **Retornado nas consultas**: Sim

### Uso:
```json
POST /api/v1/fueling
{
  "veiculoId": 5,
  "postoId": 10,
  "dataAbastecimento": "2025-12-09T10:30:00Z",
  "tipoCombustivel": "Diesel_S10",
  "litros": 150.5,
  "valorUnitario": 5.80,
  "valorTotal": 872.90,
  "valorArla": 120.50,
  "valorDiversos": 0,
  "km": 45000,
  "numeroDocumento": "NF-12345",
  "cheio": true
}
```

### Retorno:
```json
{
  "id": 234,
  "veiculoId": 5,
  "postoId": 10,
  "dataAbastecimento": "2025-12-09T10:30:00Z",
  "tipoCombustivel": "Diesel_S10",
  "litros": 150.5,
  "valorUnitario": 5.80,
  "valorTotal": 872.90,
  "valorArla": 120.50,
  "valorDiversos": 0,
  "km": 45000,
  "numeroDocumento": "NF-12345",
  "cheio": true,
  "createdAt": "2025-12-09T10:30:00Z",
  "updatedAt": "2025-12-09T10:30:00Z"
}
```

---

## valorDiversos (Não Salvo no Banco)

Campo usado apenas para **somar no valor total** do lançamento financeiro automático.

### Características:
- **Tipo**: `float64` (apenas em memória)
- **Obrigatório**: Não
- **Editável**: Não (não pode ser editado pois não é persistido)
- **Salvo no banco**: **NÃO**
- **Retornado nas consultas**: Sim (sempre será 0 nas consultas)
- **Uso**: Apenas no momento da criação para cálculo do financeiro

### Propósito:
Permite incluir valores extras (taxas, serviços, etc.) no lançamento financeiro **sem alterar o valor total do abastecimento** registrado no banco.

### Exemplo:

**Request:**
```json
POST /api/v1/fueling
{
  "veiculoId": 5,
  "postoId": 10,
  "dataAbastecimento": "2025-12-09T10:30:00Z",
  "tipoCombustivel": "Diesel_S10",
  "litros": 150.5,
  "valorUnitario": 5.80,
  "valorTotal": 872.90,
  "valorArla": 120.50,
  "valorDiversos": 50.00,
  "km": 45000,
  "numeroDocumento": "NF-12345",
  "cheio": true
}
```

**Resultado:**

1. **Abastecimento salvo no banco:**
```json
{
  "id": 234,
  "valorTotal": 872.90,
  "valorArla": 120.50,
  "valorDiversos": 0
}
```
*Nota: valorDiversos não é salvo, então sempre retorna 0 nas consultas*

2. **Lançamento financeiro criado automaticamente:**
```json
{
  "id": 567,
  "origem": "ABASTECIMENTO",
  "origemId": 234,
  "valor": 922.90,
  "valorParcela": 922.90,
  "observacao": "Lançamento automático de abastecimento"
}
```
*Nota: valor = valorTotal (872.90) + valorDiversos (50.00) = 922.90*

---

## Cálculo do Lançamento Financeiro

### Fórmula:
```
Valor do Lançamento Financeiro = valorTotal + valorDiversos
```

### Implementação no Service:

```go
func (s *FuelingService) Add(dieselAdd *entities.Fueling) error {
    // Adiciona o abastecimento e obtém o ID
    fuelingId, err := s.RepoManager.Fueling().Add(*dieselAdd)
    if err != nil {
        return err
    }

    // Cria automaticamente um lançamento a pagar
    // Soma valorDiversos ao valorTotal para o lançamento financeiro
    valorFinanceiro := dieselAdd.ValorTotal + dieselAdd.ValorDiversos
    
    origemId := fuelingId
    finance := entities.Finance{
        PessoaId:        pessoa.PessoaId,
        CategoriaId:     2,
        OrigemId:        &origemId,
        Origem:          "ABASTECIMENTO",
        Valor:           valorFinanceiro,  // <-- Soma de valorTotal + valorDiversos
        ValorParcela:    valorFinanceiro,
        NumeroParcela:   1,
        TotalParcelas:   1,
        NumeroDocumento: dieselAdd.NumeroDocumento,
        DataCompetencia: dieselAdd.Data,
        DataVencimento:  dieselAdd.Data,
        Observacao:      "Lançamento automático de abastecimento",
    }

    err = s.RepoManager.Finance().Add(finance)
    return nil
}
```

---

## Casos de Uso

### Caso 1: Abastecimento Simples (sem Arla, sem diversos)
```json
{
  "valorTotal": 872.90,
  "valorArla": 0,
  "valorDiversos": 0
}
```
- **Salvo no banco**: valorTotal = 872.90, valorArla = 0
- **Lançamento financeiro**: 872.90

---

### Caso 2: Abastecimento com Arla
```json
{
  "valorTotal": 872.90,
  "valorArla": 120.50,
  "valorDiversos": 0
}
```
- **Salvo no banco**: valorTotal = 872.90, valorArla = 120.50
- **Lançamento financeiro**: 872.90

---

### Caso 3: Abastecimento com Taxa Extra (valorDiversos)
```json
{
  "valorTotal": 872.90,
  "valorArla": 0,
  "valorDiversos": 30.00
}
```
- **Salvo no banco**: valorTotal = 872.90, valorArla = 0
- **Lançamento financeiro**: 902.90 (872.90 + 30.00)

---

### Caso 4: Abastecimento Completo (com Arla e diversos)
```json
{
  "valorTotal": 872.90,
  "valorArla": 120.50,
  "valorDiversos": 50.00
}
```
- **Salvo no banco**: valorTotal = 872.90, valorArla = 120.50
- **Lançamento financeiro**: 922.90 (872.90 + 50.00)

---

## UPDATE/Edição

### valorArla
Pode ser **editado** normalmente via PUT:
```json
PUT /api/v1/fueling
{
  "id": 234,
  "valorArla": 150.00
}
```

### valorDiversos
**Não pode ser editado** pois não é persistido no banco. Apenas influencia no momento da criação do lançamento financeiro.

---

## Migração do Banco de Dados

Execute o script SQL para adicionar a coluna `valor_arla`:

```sql
-- scripts/009_add_valor_arla_to_abastecimento.sql
ALTER TABLE abastecimento 
ADD COLUMN valor_arla DECIMAL(10,2) NOT NULL DEFAULT 0.00 
AFTER valor_total;
```

---

## Estrutura da Entidade

```go
type Fueling struct {
    Id              int64     `json:"id"`
    VeiculoId       int64     `json:"veiculoId"`
    PostoId         int64     `json:"postoId"`
    Data            time.Time `json:"dataAbastecimento"`
    TipoCombustivel string    `json:"tipoCombustivel"`
    Litros          float64   `json:"litros"`
    ValorUnitario   float64   `json:"valorUnitario"`
    ValorTotal      float64   `json:"valorTotal"`
    ValorArla       float64   `json:"valorArla"`       // Salvo no banco
    ValorDiversos   float64   `json:"valorDiversos"`   // NÃO salvo no banco
    Km              int64     `json:"km"`
    NumeroDocumento string    `json:"numeroDocumento"`
    Cheio           bool      `json:"cheio"`
    CreatedAt       time.Time `json:"createdAt"`
    UpdatedAt       time.Time `json:"updatedAt"`
}
```

---

## Diferenças entre valorArla e valorDiversos

| Característica | valorArla | valorDiversos |
|---------------|-----------|---------------|
| Salvo no banco | ✅ Sim | ❌ Não |
| Editável | ✅ Sim | ❌ Não |
| Retornado nas consultas | ✅ Sim (valor real) | ✅ Sim (sempre 0) |
| Usado no cálculo financeiro | ❌ Não | ✅ Sim |
| Persistido | ✅ Sim | ❌ Não (apenas em memória) |

---

## Observações Importantes

1. **valorDiversos é temporário**: Usado apenas no momento da criação do abastecimento para calcular o lançamento financeiro.

2. **valorArla é permanente**: Fica registrado no banco de dados para futuras consultas e relatórios.

3. **Lançamento financeiro não é atualizado**: Se você editar o abastecimento posteriormente, o lançamento financeiro criado automaticamente **não será atualizado**.

4. **Consultas retornam valorDiversos = 0**: Como não é salvo no banco, sempre retornará 0 em consultas GET.

5. **Rastreabilidade**: O lançamento financeiro possui `origemId` e `origem = "ABASTECIMENTO"` para rastrear o abastecimento original.
