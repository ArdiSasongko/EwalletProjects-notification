CREATE TABLE IF NOT EXISTS notification_history (
    id SERIAL PRIMARY KEY,
    recipient VARCHAR(255) NOT NULL,
    template_id INT NOT NULL,
    status VARCHAR(255) NOT NULL,
    error_message TEXT,
    created_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(0) DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_template_id FOREIGN KEY (template_id) REFERENCES notification_template(id)
);