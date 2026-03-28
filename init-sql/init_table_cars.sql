-- 1. Удаляем старую таблицу (БЕЗОПАСНО!)
DROP TABLE IF EXISTS cars CASCADE;

-- 2. Создаем новую таблицу под models.Car
CREATE TABLE cars (
    car_id BIGINT PRIMARY KEY,
    department TEXT DEFAULT '',
    resource TEXT DEFAULT '',
    from_country TEXT DEFAULT '',
    link TEXT DEFAULT '',
    vehicle_type TEXT DEFAULT '',
    vin TEXT DEFAULT '',
    make TEXT DEFAULT '',
    model TEXT DEFAULT '',
    month BIGINT DEFAULT 0,
    year BIGINT DEFAULT 0,
    age BIGINT DEFAULT 0,
    body_type TEXT DEFAULT '',
    is_right_steering BOOLEAN DEFAULT false,
    color TEXT DEFAULT '',
    trim TEXT DEFAULT '',
    mileage BIGINT DEFAULT 0,
    fuel TEXT DEFAULT '',
    engine_volume BIGINT DEFAULT 0,
    horse_power BIGINT DEFAULT 0,
    transmission TEXT DEFAULT '',
    drive_type TEXT DEFAULT '',
    photos TEXT[] DEFAULT '{}',
    price BIGINT DEFAULT 0,
    price_currency TEXT DEFAULT 'RUB',
    additional_context TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. Индексы
CREATE INDEX idx_cars_car_id ON cars(car_id);
CREATE INDEX idx_cars_make_model ON cars(make, model);
CREATE INDEX idx_cars_year_price ON cars(year, price);

-- 4. Тестовые данные (ПОЛНОСТЬЮ заполненные!)
INSERT INTO cars (car_id, department, resource, from_country, link, vehicle_type, vin, 
    make, model, month, year, age, body_type, is_right_steering, color, trim, mileage,
    fuel, engine_volume, horse_power, transmission, drive_type, photos, price, 
    price_currency, additional_context) VALUES 
    -- Toyota Camry
    (123456, 'Легковые', 'avito.ru', 'Россия', 'https://avito.ru/toyota_camry_123456', 
     'passenger', 'XWV93ABC123456789', 'Toyota', 'Camry', 6, 2023, 2, 'седан', 
     false, 'белый', '2.5 AT Premium', 45000, 'бензин', 2500, 181, 'автомат', 
     'передний', ARRAY['photo1.jpg', 'photo2.jpg', 'photo3.jpg'], 3500000, 'RUB', 
     'ПТС оригинал, 1 владелец, не битая'),
     
    -- BMW X5  
    (123457, 'Внедорожники', 'avito.ru', 'Россия', 'https://avito.ru/bmw_x5_123457', 
     'suv', 'WBA12345678901234', 'BMW', 'X5', 3, 2024, 1, 'внедорожник', 
     false, 'черный', 'xDrive40i', 18000, 'бензин', 3000, 360, 'автомат', 
     'полный', ARRAY['photo4.jpg', 'photo5.jpg'], 8500000, 'RUB', 
     'xDrive, панорама, состояние идеал'),
     
    -- Lada Vesta (бюджет)
    (123458, 'Легковые', 'avito.ru', 'Россия', 'https://avito.ru/lada_vesta_123458', 
     'passenger', 'XTA12345678901234', 'Lada', 'Vesta', 12, 2022, 3, 'седан', 
     false, 'серый', '1.6 Lux', 65000, 'бензин', 1600, 106, 'автомат', 
     'передний', ARRAY['photo6.jpg'], 1650000, 'RUB', 
     'Автомат, климат, состояние отличное'),
     
    -- Hyundai Solaris (популярный)
    (123459, 'Легковые', 'avito.ru', 'Россия', 'https://avito.ru/hyundai_solaris_123459', 
     'passenger', 'KMH12345678901234', 'Hyundai', 'Solaris', 9, 2023, 2, 'седан', 
     false, 'синий', '1.6 Comfort', 32000, 'бензин', 1600, 123, 'механика', 
     'передний', ARRAY['photo7.jpg', 'photo8.jpg'], 1850000, 'RUB', 
     'Механика, 2 владельца')
ON CONFLICT (car_id) DO NOTHING;