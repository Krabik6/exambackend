
-- Выбор базы данных для выполнения последующих команд
\c your_db_name

-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     login VARCHAR(255) UNIQUE NOT NULL,
                                     password VARCHAR(255) NOT NULL,
                                     full_name VARCHAR(255) NOT NULL,
                                     phone VARCHAR(20) NOT NULL,
                                     email VARCHAR(255) UNIQUE NOT NULL,
                                     role VARCHAR(50) DEFAULT 'user'
);

-- Создание таблицы заявлений
CREATE TABLE IF NOT EXISTS violations (
                                          id SERIAL PRIMARY KEY,
                                          user_id INT NOT NULL,
                                          car_number VARCHAR(20) NOT NULL,
                                          description TEXT NOT NULL,
                                          status VARCHAR(50) NOT NULL DEFAULT 'новое',
                                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                          FOREIGN KEY (user_id) REFERENCES users (id)
);


-- Добавление пользователя-администратора с хешированным паролем
INSERT INTO users (login, password, full_name, phone, email, role)
VALUES ('copp', '$2a$10$IuqabjnB3AY.FyJE8qmhCO.aLhRuVAwo60vlFqJ87pdWxgS167Epe', 'Администратор Системы', '+00000000000', 'admin@example.com', 'admin')
ON CONFLICT (login) DO NOTHING;
