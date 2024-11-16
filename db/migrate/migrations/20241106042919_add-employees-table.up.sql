CREATE TABLE IF NOT EXISTS Employees (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    lastname VARCHAR(100),
    email VARCHAR(255) UNIQUE,
    department VARCHAR(100),
    position VARCHAR(100),
    INDEX(department),
    INDEX(position)
);