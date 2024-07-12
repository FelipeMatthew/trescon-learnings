-- CREATE DATABASE

CREATE TABLE library (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    author VARCHAR(100),
    genre VARCHAR(50)
);

INSERT INTO library (name, author, genre) VALUES 
('1984', 'George Orwell', 'Dystopian'),
('To Kill a Mockingbird', 'Harper Lee', 'Fiction'),
('The Great Gatsby', 'F. Scott Fitzgerald', 'Classic');