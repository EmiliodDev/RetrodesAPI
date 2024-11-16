CREATE TABLE IF NOT EXISTS SupportServices (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    employee_id BIGINT,
    type_support VARCHAR(100),
    description TEXT,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (employee_id) REFERENCES Employees(id),
    INDEX (employee_id),
    INDEX (date)
);