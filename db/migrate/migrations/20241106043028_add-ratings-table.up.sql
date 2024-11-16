CREATE TABLE IF NOT EXISTS Ratings (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    manager_id BIGINT,
    rating INT,
    comment TEXT,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (manager_id) REFERENCES Employees(id),
    INDEX (manager_id),
    INDEX (date)
);