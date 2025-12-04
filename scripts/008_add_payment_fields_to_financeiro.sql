-- Add payment fields to financeiro table
ALTER TABLE financeiro 
ADD COLUMN forma_pagamento_id BIGINT NULL,
ADD COLUMN valor_pago DECIMAL(10,2) NULL,
ADD COLUMN realizado BOOLEAN DEFAULT FALSE,
ADD CONSTRAINT fk_forma_pagamento FOREIGN KEY (forma_pagamento_id) REFERENCES forma_pagamento(id);

-- Create index for better performance
CREATE INDEX idx_financeiro_realizado ON financeiro(realizado);
CREATE INDEX idx_financeiro_forma_pagamento ON financeiro(forma_pagamento_id);
