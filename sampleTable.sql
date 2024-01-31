CREATE TABLE operator_logs (
    id SERIAL PRIMARY KEY,
    operator_name VARCHAR(100),
    log_time TIMESTAMP,
    log_file TEXT
);

SELECT * FROM operator_logs;
