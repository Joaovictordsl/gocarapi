CREATE TABLE IF NOT EXISTS CARROS (
    id SERIAL PRIMARY KEY,
    marca VARCHAR(50) NOT NULL,
    modelo VARCHAR(50) NOT NULL,
    preco NUMERIC(7) NOT NULL,
    created timestamp DEFAULT NOW()
)