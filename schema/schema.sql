-- Drop database if exists
DROP DATABASE IF EXISTS ia03_db;

-- Create database
CREATE DATABASE ia03_db;

-- Set timezone to Vietnam
SET TIME ZONE 'Asia/Ho_Chi_Minh';

CREATE TABLE tbl_users (
    user_id       SERIAL PRIMARY KEY,
    full_name     VARCHAR(30) NOT NULL, 
    email         VARCHAR(40) NOT NULL, 
    password      VARCHAR(100) NOT NULL,
    salt          VARCHAR(100) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
