-- +migrate Up
CREATE TABLE IF NOT EXISTS cves_products (
  cve_id     VARCHAR(64) NOT NULL,
  product_id INTEGER     NOT NULL
);

CREATE INDEX IF NOT EXISTS cves_products_cve_id
  ON cves_products (cve_id);
CREATE INDEX IF NOT EXISTS cves_products_product_id
  ON cves_products (product_id);

CREATE UNIQUE INDEX IF NOT EXISTS cves_products_unique
  ON cves_products (cve_id, product_id);

-- +migrate Down
DROP TABLE IF EXISTS cves_products;
