-- Таблица cars
CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    brand VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    year INTEGER NOT NULL CHECK (year > 1880 AND year <= 2026),
    price_rub INTEGER NOT NULL CHECK (price_rub >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для производительности
CREATE INDEX IF NOT EXISTS idx_cars_brand ON cars(brand);
CREATE INDEX IF NOT EXISTS idx_cars_year ON cars(year);

-- Тестовые данные
INSERT INTO cars (brand, model, year, price_rub) VALUES 
    ('Toyota', 'Camry', 2023, 3500000),
    ('BMW', 'X5', 2024, 8500000),
    ('Lada', 'Vesta', 2022, 1500000)
ON CONFLICT DO NOTHING;
