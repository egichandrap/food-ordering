CREATE TABLE payments (
                          id VARCHAR(100) PRIMARY KEY,
                          order_id VARCHAR(100) REFERENCES orders(id),
                          amount DECIMAL(10, 2),
                          payment_status VARCHAR(50) DEFAULT 'unpaid',
                          payment_method VARCHAR(50),
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);