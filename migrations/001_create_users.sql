-- Создаем таблицу users.
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,              -- Уникальный ID, автоматически увеличивается
                       telegram_id BIGINT UNIQUE NOT NULL, -- Telegram ID, уникальный и обязательный
                       username VARCHAR(255),             -- Имя пользователя, до 255 символов
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Время создания, по умолчанию текущее
);
