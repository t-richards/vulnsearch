-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
  id      INTEGER PRIMARY KEY,
  name    VARCHAR(255) NOT NULL,
  vendor  VARCHAR(255) NOT NULL,
  version VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS products_unique
  ON products (name, vendor, version);

-- +migrate Down
DROP TABLE IF EXISTS products;
