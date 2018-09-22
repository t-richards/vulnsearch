-- +micrate Up
BEGIN;
CREATE TABLE IF NOT EXISTS products (
  id      INTEGER PRIMARY KEY,
  name    VARCHAR(255) NOT NULL,
  vendor  VARCHAR(255) NOT NULL,
  version VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS cves_products (
  cve_id     VARCHAR(64) NOT NULL,
  product_id INTEGER     NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS products_name_vendor_version
  ON products (name, vendor, version);

CREATE INDEX IF NOT EXISTS cves_products_cve_id
  ON cves_products (cve_id);
CREATE INDEX IF NOT EXISTS cves_products_product_id
  ON cves_products (product_id);
END;

-- +micrate Down
BEGIN;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS cves_products;
END;
