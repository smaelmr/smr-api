-- Adiciona coluna valor_arla na tabela abastecimento
ALTER TABLE abastecimento 
ADD COLUMN valor_arla DECIMAL(10,2) NOT NULL DEFAULT 0.00 
AFTER valor_total;
