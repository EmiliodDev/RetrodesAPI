CREATE TABLE IF NOT EXISTS Credentials (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    employee_id BIGINT UNIQUE,
    username VARCHAR(50) UNIQUE,
    password CHAR(60),
    FOREIGN KEY (employee_id) REFERENCES Employees(id)
);