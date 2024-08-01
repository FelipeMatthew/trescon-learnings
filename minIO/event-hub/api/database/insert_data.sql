INSERT INTO admins (username, password)
VALUES ('admin_user', crypt('123', gen_salt('bf', 4)));


