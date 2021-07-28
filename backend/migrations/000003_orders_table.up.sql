CREATE TABLE IF NOT EXISTS orders(
  seller_id VARCHAR(100) REFERENCES users(email),
  user_id VARCHAR(100) REFERENCES users(email),
  product_id CHARACTER(27) REFERENCES products(id),
  units INT NOT  NULL DEFAULT 1,
  created_at TIMESTAMPTZ  NOT NULL DEFAULT (NOW()),
  PRIMARY KEY (user_id,product_id)
);
