CREATE TABLE IF NOT EXISTS products(
  id CHARACTER(27)  PRIMARY KEY,
  seller_id VARCHAR(100)  REFERENCES users(email),
  name VARCHAR (50) NOT NULL,
  stock INT NOT NULL DEFAULT 0,
  category VARCHAR(50) NOT NULL,
  created_at TIMESTAMPTZ  NOT NULL DEFAULT (NOW())
);

ALTER TABLE products ADD description VARCHAR(1000);

