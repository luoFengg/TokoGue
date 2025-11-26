-- Tabel Header Order
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(100) PRIMARY KEY, -- UUID
    user_id VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending', -- pending, paid, cancelled
    total_price BIGINT NOT NULL, -- Total belanjaan
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign Key ke tabel Users
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Tabel Detail Item (Keranjang Belanja)
CREATE TABLE IF NOT EXISTS order_items (
    id VARCHAR(100) PRIMARY KEY,
    order_id VARCHAR(100) NOT NULL,
    product_id VARCHAR(100) NOT NULL,
    quantity INTEGER NOT NULL,
    price INTEGER NOT NULL, -- HARGA SAAT BELI (Penting!)
    
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE SET NULL
);

-- migrate -path apps/databases/migrations -database "postgresql://postgres:postgres@localhost:5432/tokogue?sslmode=disable" up