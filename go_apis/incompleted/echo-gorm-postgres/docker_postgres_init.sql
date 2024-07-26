CREATE TABLE students {
  id BIGINT NOT NULL,
  name TEXT COLLATE pg_catalog."default",
  CONSTRAINT student_pkey PRIMARY KEY (id)
}

INSERT INTO students(id, name) VALUES
(1, "felipe"),
(2, "roberto"),
(3, "jose");