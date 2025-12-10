-- Adiciona coluna qtd_arla na tabela abastecimento
ALTER TABLE abastecimento 
ADD COLUMN qtd_arla DECIMAL(10,2) NOT NULL DEFAULT 0.00 
AFTER valor_diesel;
