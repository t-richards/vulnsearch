-- +micrate Up
CREATE TABLE IF NOT EXISTS products (
  id           INTEGER PRIMARY KEY,
  product_name VARCHAR(255),
  vendor       VARCHAR(255),
  version      VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS cves_products (
  cve_id     VARCHAR(64),
  product_id INTEGER
);

CREATE INDEX cves_products_cve_id
  ON cves_products (cve_id);
CREATE INDEX cves_products_product_id
  ON cves_products (product_id);

-- +micrate Down
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS cves_products;
