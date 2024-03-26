CREATE TABLE user_classes(
    id SERIAL PRIMARY KEY,
    user_id uuid  NOT NULL  REFERENCES users(id),
    class_id uuid  NOT NULL REFERENCES classes(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
