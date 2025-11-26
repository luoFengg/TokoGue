CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(100) PRIMARY KEY,

    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'customer',
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT check_role CHECK (role IN ('admin', 'customer'))
);

-- migrate -path apps/databases/migrations -database "postgresql://postgres:postgres@localhost:5432/tokogue?sslmode=disable" up