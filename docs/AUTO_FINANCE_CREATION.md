# Criação Automática de Lançamentos Financeiros

## Descrição

O sistema agora cria automaticamente lançamentos financeiros quando um **Abastecimento** ou **Frete** é cadastrado.

---

## Abastecimento → Conta a Pagar

Quando um novo **Abastecimento** é criado, o sistema automaticamente:

1. Salva o registro de abastecimento na tabela `abastecimento`
2. Obtém o ID do abastecimento recém-criado
3. Cria um lançamento financeiro em **Contas a Pagar** com:
   - **PessoaId**: `PostoId` (fornecedor do posto de combustível)
   - **CategoriaId**: `1` (categoria de "Abastecimento")
   - **OrigemId**: ID do abastecimento
   - **Origem**: `"ABASTECIMENTO"`
   - **Valor**: `ValorTotal` do abastecimento
   - **ValorParcela**: `ValorTotal` (pagamento único)
   - **NumeroParcela**: `1`
   - **TotalParcelas**: `1`
   - **NumeroDocumento**: Mesmo do abastecimento
   - **DataCompetencia**: Data do abastecimento
   - **DataVencimento**: Data do abastecimento
   - **Observacao**: "Lançamento automático de abastecimento"

### Exemplo de Fluxo

```json
POST /api/v1/fueling
{
  "veiculoId": 5,
  "postoId": 10,
  "dataAbastecimento": "2025-12-08T10:30:00Z",
  "tipoCombustivel": "Diesel_S10",
  "litros": 150.5,
  "valorUnitario": 5.80,
  "valorTotal": 872.90,
  "km": 45000,
  "numeroDocumento": "NF-12345",
  "cheio": true
}
```

**Resultado:**
1. Abastecimento criado com ID `234`
2. Lançamento financeiro criado automaticamente:

```json
{
  "id": 567,
  "pessoaId": 10,
  "categoriaId": 1,
  "origemId": 234,
  "origem": "ABASTECIMENTO",
  "valor": 872.90,
  "valorParcela": 872.90,
  "numeroParcela": 1,
  "totalParcelas": 1,
  "numeroDocumento": "NF-12345",
  "dataCompetencia": "2025-12-08T10:30:00Z",
  "dataVencimento": "2025-12-08T10:30:00Z",
  "observacao": "Lançamento automático de abastecimento",
  "status": "Em Aberto"
}
```

---

## Frete → Conta a Receber

Quando um novo **Frete** é criado, o sistema automaticamente:

1. Salva o registro de frete na tabela `frete`
2. Obtém o ID do frete recém-criado
3. Cria um lançamento financeiro em **Contas a Receber** com:
   - **PessoaId**: `ClienteId` (cliente do frete)
   - **CategoriaId**: `2` (categoria de "Frete")
   - **OrigemId**: ID do frete
   - **Origem**: `"FRETE"`
   - **Valor**: `ValorFrete`
   - **ValorParcela**: `ValorFrete` (recebimento único)
   - **NumeroParcela**: `1`
   - **TotalParcelas**: `1`
   - **NumeroDocumento**: Mesmo do frete
   - **DataCompetencia**: Data de carregamento
   - **DataVencimento**: Data de entrega
   - **Observacao**: "Lançamento automático de frete"

### Exemplo de Fluxo

```json
POST /api/v1/trip
{
  "cavaloPlaca": "ABC1234",
  "clienteId": 15,
  "origemId": 100,
  "destinoId": 200,
  "motoristaId": 8,
  "dataCarregamento": "2025-12-08T08:00:00Z",
  "dataEntrega": "2025-12-10T18:00:00Z",
  "numeroDocumento": "CTRC-98765",
  "valorAgenciamento": 500,
  "valorFrete": 5000,
  "valorPedagio": 200,
  "observacoes": "Carga frágil"
}
```

**Resultado:**
1. Frete criado com ID `345`
2. Lançamento financeiro criado automaticamente:

```json
{
  "id": 678,
  "pessoaId": 15,
  "categoriaId": 2,
  "origemId": 345,
  "origem": "FRETE",
  "valor": 5000.00,
  "valorParcela": 5000.00,
  "numeroParcela": 1,
  "totalParcelas": 1,
  "numeroDocumento": "CTRC-98765",
  "dataCompetencia": "2025-12-08T08:00:00Z",
  "dataVencimento": "2025-12-10T18:00:00Z",
  "observacao": "Lançamento automático de frete",
  "status": "Em Aberto"
}
```

---

## Status do Lançamento

O campo `status` é calculado dinamicamente:

- **"PAGO"**: Quando o abastecimento possui `dataRealizacao` preenchida
- **"RECEBIDO"**: Quando o frete possui `dataRealizacao` preenchida
- **"Em Atraso"**: Quando não possui `dataRealizacao` e já passou do vencimento
- **"EM_ABERTO"**: Quando não possui `dataRealizacao` e o vencimento ainda não passou

---

## Implementação Técnica

### Mudanças nas Interfaces

```go
// Antes
type FuelingRepository interface {
    Add(diesel entities.Fueling) error
}

type TripRepository interface {
    Add(trip entities.Trip) error
}

// Depois
type FuelingRepository interface {
    Add(diesel entities.Fueling) (int64, error)
}

type TripRepository interface {
    Add(trip entities.Trip) (int64, error)
}
```

### FuelingService.Add

```go
func (s *FuelingService) Add(dieselAdd *entities.Fueling) error {
    // Adiciona o abastecimento e obtém o ID
    fuelingId, err := s.RepoManager.Fueling().Add(*dieselAdd)
    if err != nil {
        return err
    }

    // Cria automaticamente um lançamento a pagar
    origemId := fuelingId
    finance := entities.Finance{
        PessoaId:        dieselAdd.PostoId,
        CategoriaId:     1, // Categoria de Abastecimento
        OrigemId:        &origemId,
        Origem:          "ABASTECIMENTO",
        Valor:           dieselAdd.ValorTotal,
        ValorParcela:    dieselAdd.ValorTotal,
        NumeroParcela:   1,
        TotalParcelas:   1,
        NumeroDocumento: dieselAdd.NumeroDocumento,
        DataCompetencia: dieselAdd.Data,
        DataVencimento:  dieselAdd.Data,
        Observacao:      "Lançamento automático de abastecimento",
    }

    err = s.RepoManager.Finance().Add(finance)
    // Se falhar ao criar o lançamento, não impede a criação do abastecimento

    return nil
}
```

### TripService.Add

```go
func (s *TripService) Add(tripAdd *entities.Trip) error {
    // Adiciona o frete e obtém o ID
    tripId, err := s.RepoManager.Trip().Add(*tripAdd)
    if err != nil {
        return err
    }

    // Cria automaticamente um lançamento a receber
    origemId := tripId
    finance := entities.Finance{
        PessoaId:        tripAdd.ClienteId,
        CategoriaId:     2, // Categoria de Frete
        OrigemId:        &origemId,
        Origem:          "FRETE",
        Valor:           float64(tripAdd.ValorFrete),
        ValorParcela:    float64(tripAdd.ValorFrete),
        NumeroParcela:   1,
        TotalParcelas:   1,
        NumeroDocumento: tripAdd.NumeroDocumento,
        DataCompetencia: tripAdd.DataCarregamento,
        DataVencimento:  tripAdd.DataEntrega,
        Observacao:      "Lançamento automático de frete",
    }

    err = s.RepoManager.Finance().Add(finance)
    // Se falhar ao criar o lançamento, não impede a criação do frete

    return nil
}
```

---

## Rastreabilidade

Com o campo `OrigemId` e `Origem`, é possível:

1. **Identificar a origem** do lançamento financeiro
2. **Voltar ao registro original** (abastecimento ou frete)
3. **Filtrar lançamentos** por tipo de origem
4. **Gerar relatórios** específicos por origem

### Exemplos de Consulta

```sql
-- Todos os lançamentos de abastecimento
SELECT * FROM financeiro WHERE origem = 'ABASTECIMENTO';

-- Lançamento específico de um abastecimento
SELECT * FROM financeiro WHERE origemId = 234 AND origem = 'ABASTECIMENTO';

-- Todos os lançamentos de frete
SELECT * FROM financeiro WHERE origem = 'FRETE';

-- Valor total a receber de fretes
SELECT SUM(valor) FROM financeiro 
WHERE origem = 'FRETE' AND dataRealizacao IS NULL;
```

---

## Tratamento de Erros

Se a criação do lançamento financeiro falhar:
- O **abastecimento/frete** será criado normalmente
- O erro na criação do lançamento **não bloqueia** o processo
- Recomenda-se implementar um **sistema de logs** para rastrear esses erros

---

## TODO / Melhorias Futuras

1. **Categorias Dinâmicas**: 
   - Buscar as categorias corretas do banco ao invés de usar IDs fixos (1 e 2)
   - Criar categorias padrão se não existirem

2. **Logs de Erro**:
   - Implementar sistema de log quando falhar a criação do lançamento financeiro

3. **Transações**:
   - Considerar uso de transações SQL para garantir atomicidade

4. **Configurações**:
   - Permitir configurar quais tipos de origem devem criar lançamentos automaticamente

5. **Webhooks/Eventos**:
   - Emitir eventos quando lançamentos automáticos são criados
