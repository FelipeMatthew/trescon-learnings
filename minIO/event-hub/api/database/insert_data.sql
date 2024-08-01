INSERT INTO admins (username, password)
VALUES ('admin_user', crypt('123', gen_salt('bf', 4)));

INSERT INTO events (name, description, init_date, finish_date, location, ticket_price, image_url)
VALUES
('Music Festival', 'An exciting music festival with various artists.', '2024-09-15 18:00:00', '2024-09-15 23:00:00', 'Central Park', 99.99, 'http://example.com/images/music-festival.jpg'),
('Tech Conference', 'A conference featuring the latest in technology.', '2024-10-05 09:00:00', '2024-10-05 17:00:00', 'Convention Center', 299.99, 'http://example.com/images/tech-conference.jpg'),
('Art Exhibition', 'An exhibition showcasing contemporary art.', '2024-11-01 10:00:00', '2024-11-01 20:00:00', 'Art Gallery', 49.99, 'http://example.com/images/art-exhibition.jpg');

INSERT INTO participants (event_id, profile_image, document_file, first_name, last_name, age, email, contact)
VALUES
(1, 'http://example.com/images/profile1.jpg', 'documents/doc1.pdf', 'Alice', 'Johnson', 28, 'alice.johnson@example.com', '555-0101'),
(1, 'http://example.com/images/profile2.jpg', 'documents/doc2.pdf', 'Bob', 'Smith', 35, 'bob.smith@example.com', '555-0102'),
(2, 'http://example.com/images/profile3.jpg', 'documents/doc3.pdf', 'Charlie', 'Davis', 40, 'charlie.davis@example.com', '555-0103'),
(2, 'http://example.com/images/profile4.jpg', 'documents/doc4.pdf', 'Dana', 'Brown', 31, 'dana.brown@example.com', '555-0104'),
(3, 'http://example.com/images/profile5.jpg', 'documents/doc5.pdf', 'Eva', 'White', 26, 'eva.white@example.com', '555-0105');