CREATE TABLE orders (
                        id VARCHAR(100) PRIMARY KEY,
                        menu_id VARCHAR(100) REFERENCES menus(id),
                        quantity INT NOT NULL,
                        total_price DECIMAL(10, 2),
                        status VARCHAR(50) DEFAULT 'pending',
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);