CREATE TABLE IF NOT EXISTS Complaints (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    employee_id BIGINT,
    type ENUM('anonymous', 'not anonymous'),
    content TEXT,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (employee_id) REFERENCES Employees(id),
    INDEX (employee_id),
    INDEX(date)
);