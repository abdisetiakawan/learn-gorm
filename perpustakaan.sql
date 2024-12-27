-- Active: 1724149523236@@127.0.0.1@3306@perpustakaan
CREATE TABLE author (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    nationality VARCHAR(50),
    birth_year INT
);


CREATE TABLE buku (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    genre VARCHAR(50),
    publication_year INT,
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES author(id)
);


INSERT INTO author (name, nationality, birth_year) VALUES
('J.K. Rowling', 'British', 1965),
('George Orwell', 'British', 1903),
('Haruki Murakami', 'Japanese', 1949),
('Agatha Christie', 'British', 1890),
('Gabriel García Márquez', 'Colombian', 1927);

INSERT INTO buku (title, genre, publication_year, author_id) VALUES
('Harry Potter and the Philosopher Stone', 'Fantasy', 1997, 1),
('1984', 'Dystopian', 1949, 2),
('Norwegian Wood', 'Fiction', 1987, 3),
('Murder on the Orient Express', 'Mystery', 1934, 4),
('One Hundred Years of Solitude', 'Magic Realism', 1967, 5);

DELETE FROM buku;
DELETE FROM author;

SELECT * FROM buku;
SELECT * FROM author;

DROP DATABASE perpustakaan;
CREATE DATABASE perpustakaan;
USE perpustakaan;

