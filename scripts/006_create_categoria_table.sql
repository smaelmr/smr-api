-- Create table for finance categories
CREATE TABLE IF NOT EXISTS categoria (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    descricao VARCHAR(255) NOT NULL,
    tipo CHAR(1) NOT NULL CHECK (tipo IN ('R', 'D')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_tipo (tipo)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert some default categories
INSERT INTO categoria (descricao, tipo) VALUES 
    ('Venda de Mercadorias', 'R'),
    ('Prestação de Serviços', 'R'),
    ('Juros Recebidos', 'R'),
    ('Combustível', 'D'),
    ('Manutenção', 'D'),
    ('Salários', 'D'),
    ('Aluguel', 'D'),
    ('Impostos', 'D');
