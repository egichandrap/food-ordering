CREATE TABLE menus (
                       id VARCHAR(100) PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       description TEXT,
                       price DECIMAL(10, 2) NOT NULL,
                       currency VARCHAR(100) DEFAULT 'Rp',
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
