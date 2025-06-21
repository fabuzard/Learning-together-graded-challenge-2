-- ddl.sql

-- Drop tables in reverse order of dependency to allow recreation without foreign key conflicts
DROP TABLE IF EXISTS loans CASCADE;
DROP TABLE IF EXISTS book_genres CASCADE;
DROP TABLE IF EXISTS books CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS authors CASCADE;
DROP TABLE IF EXISTS genres CASCADE;

-- Table: users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL, -- Store hashed password
    address TEXT NOT NULL,
    date_of_birth DATE NOT NULL, -- Changed from VARCHAR to DATE
    role VARCHAR(50) DEFAULT 'user' NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Table: authors
CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Table: genres
CREATE TABLE IF NOT EXISTS genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Table: books
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    min_age_restriction INT NOT NULL,
    cover_url VARCHAR(255) NOT NULL,
    author_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_author
        FOREIGN KEY (author_id)
        REFERENCES authors(id)
        ON DELETE RESTRICT
);

-- Junction Table for Many-to-Many relationship between Books and Genres
-- Table: book_genres
CREATE TABLE IF NOT EXISTS book_genres (
    book_id INT NOT NULL,
    genre_id INT NOT NULL,
    PRIMARY KEY (book_id, genre_id),
    CONSTRAINT fk_book_genres_book
        FOREIGN KEY (book_id)
        REFERENCES books(id)
        ON DELETE CASCADE, -- If a book is deleted, remove its genre associations
    CONSTRAINT fk_book_genres_genre
        FOREIGN KEY (genre_id)
        REFERENCES genres(id)
        ON DELETE CASCADE -- If a genre is deleted, remove its book associations
);


-- Table: loans
CREATE TABLE IF NOT EXISTS loans (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL,
    start_date TIMESTAMPTZ NOT NULL, -- Changed from LoanDate to StartDate to match your model
    due_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_book
        FOREIGN KEY (book_id)
        REFERENCES books(id)
        ON DELETE RESTRICT
);

-- Seeding Data (Example Data)
-- IMPORTANT: Replace 'hashed_password_...' with actual bcrypt hashed passwords
-- You should generate these hashes using a tool or your Go application's hashing utility.

-- Insert Users
INSERT INTO users (first_name, last_name, email, password, address, date_of_birth, role) VALUES
('Budi', 'Santoso', 'budi.s@example.com', '$2a$10$HASH_FOR_PASSWORD1', 'Jl. Merdeka No.1, Jakarta', '1992-03-10', 'user'),
('Siti', 'Aminah', 'siti.a@example.com', '$2a$10$HASH_FOR_PASSWORD2', 'Jl. Kenanga No.5, Bandung', '1988-11-25', 'user'),
('Admin', 'Super', 'admin@example.com', '$2a$10$HASH_FOR_ADMIN_PASSWORD', 'Jl. Admin Raya, Cyber City', '1980-01-01', 'admin');


-- Insert Authors
INSERT INTO authors (first_name, last_name) VALUES
('Andrea', 'Hirata'),
('Tere', 'Liye'),
('Dewi', 'Lestari');

-- Insert Genres
INSERT INTO genres (name) VALUES
('Fiksi'),
('Romansa'),
('Fantasi'),
('Petualangan'),
('Sejarah');

-- Insert Books
-- Get author IDs first for foreign key insertion
INSERT INTO books (title, description, min_age_restriction, cover_url, author_id) VALUES
('Laskar Pelangi', 'Novel inspiratif tentang perjuangan anak-anak di Belitung.', 13, 'http://example.com/laskar_pelangi.jpg', (SELECT id FROM authors WHERE first_name = 'Andrea' AND last_name = 'Hirata')),
('Orang-Orang Biasa', 'Kisah perjuangan orang-orang biasa.', 15, 'http://example.com/orang_biasa.jpg', (SELECT id FROM authors WHERE first_name = 'Andrea' AND last_name = 'Hirata')),
('Hafalan Shalat Delisa', 'Kisah mengharukan tentang seorang gadis kecil.', 10, 'http://example.com/delisa.jpg', (SELECT id FROM authors WHERE first_name = 'Tere' AND last_name = 'Liye')),
('Supernova: Ksatria, Puteri, dan Bintang Jatuh', 'Novel fiksi ilmiah dan spiritual.', 17, 'http://example.com/supernova.jpg', (SELECT id FROM authors WHERE first_name = 'Dewi' AND last_name = 'Lestari'));

-- Insert into book_genres (many-to-many)
-- Laskar Pelangi (Fiksi, Petualangan)
INSERT INTO book_genres (book_id, genre_id) VALUES
((SELECT id FROM books WHERE title = 'Laskar Pelangi'), (SELECT id FROM genres WHERE name = 'Fiksi')),
((SELECT id FROM books WHERE title = 'Laskar Pelangi'), (SELECT id FROM genres WHERE name = 'Petualangan'));

-- Orang-Orang Biasa (Fiksi)
INSERT INTO book_genres (book_id, genre_id) VALUES
((SELECT id FROM books WHERE title = 'Orang-Orang Biasa'), (SELECT id FROM genres WHERE name = 'Fiksi'));

-- Hafalan Shalat Delisa (Fiksi)
INSERT INTO book_genres (book_id, genre_id) VALUES
((SELECT id FROM books WHERE title = 'Hafalan Shalat Delisa'), (SELECT id FROM genres WHERE name = 'Fiksi'));

-- Supernova (Fiksi, Romansa, Fantasi)
INSERT INTO book_genres (book_id, genre_id) VALUES
((SELECT id FROM books WHERE title = 'Supernova: Ksatria, Puteri, dan Bintang Jatuh'), (SELECT id FROM genres WHERE name = 'Fiksi')),
((SELECT id FROM books WHERE title = 'Supernova: Ksatria, Puteri, dan Bintang Jatuh'), (SELECT id FROM genres WHERE name = 'Romansa')),
((SELECT id FROM books WHERE title = 'Supernova: Ksatria, Puteri, dan Bintang Jatuh'), (SELECT id FROM genres WHERE name = 'Fantasi'));


-- Insert Loans
INSERT INTO loans (user_id, book_id, start_date, due_date) VALUES
((SELECT id FROM users WHERE email = 'budi.s@example.com'), (SELECT id FROM books WHERE title = 'Hafalan Shalat Delisa'), '2025-06-15 10:00:00+07', '2025-06-29 10:00:00+07'),
((SELECT id FROM users WHERE email = 'siti.a@example.com'), (SELECT id FROM books WHERE title = 'Laskar Pelangi'), '2025-06-18 11:30:00+07', '2025-07-02 11:30:00+07'),
((SELECT id FROM users WHERE email = 'budi.s@example.com'), (SELECT id FROM books WHERE title = 'Supernova: Ksatria, Puteri, dan Bintang Jatuh'), '2025-06-20 14:00:00+07', '2025-07-04 14:00:00+07');