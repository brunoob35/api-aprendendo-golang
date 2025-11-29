DROP DATABASE IF EXISTS treehousedb;
CREATE DATABASE treehousedb;
USE treehousedb;

CREATE TABLE users (
                       id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       cpf VARCHAR(14) NOT NULL UNIQUE,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL
);