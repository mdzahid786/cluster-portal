
CREATE TABLE IF NOT EXISTS clusters(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(200),
    servers INT
);

INSERT INTO clusters (name, servers) VALUES
('Amazon', 850),
('Michigan', 904),
('Rhino', 1208),
('Tahoe', 150);