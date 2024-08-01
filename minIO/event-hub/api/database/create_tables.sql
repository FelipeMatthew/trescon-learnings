CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE admins (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) NOT NULL,
  password VARCHAR(100) NOT NULL
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    init_date TIMESTAMP,
    finish_date TIMESTAMP,
    location VARCHAR(255),
    ticket_price DECIMAL(10, 2),
    image_url VARCHAR(255)
);

CREATE TABLE participants (
    id SERIAL PRIMARY KEY,
    event_id INTEGER REFERENCES events(id) ON DELETE CASCADE,
    profile_image VARCHAR(255),
    document_file VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    age INTEGER,
    email VARCHAR(255),
    contact VARCHAR(255)
);

