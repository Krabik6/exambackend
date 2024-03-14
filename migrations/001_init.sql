GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO root;
ALTER USER root WITH PASSWORD 'root';

-- Создание базы данных, если еще не существует
CREATE DATABASE  your_db_name;

ALTER DATABASE your_db_name CONNECTION LIMIT -1;

-- Назначение прав пользователя на базу данных
CREATE USER postgres_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE your_db_name TO postgres_user;
ALTER USER postgres_user WITH SUPERUSER;

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
                                          status VARCHAR(50) NOT NULL,
                                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                          FOREIGN KEY (user_id) REFERENCES users (id)
);

