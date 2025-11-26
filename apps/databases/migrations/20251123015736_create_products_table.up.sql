CREATE TABLE products (
    id VARCHAR(100) PRIMARY KEY,
    
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price BIGINT NOT NULL DEFAULT 0,
    stock INTEGER NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_products_deleted_at ON products(deleted_at);

-- migrate -path apps/databases/migrations -database "postgresql://postgres:postgres@localhost:5432/tokogue?sslmode=disable" up