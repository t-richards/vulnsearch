module Vulnsearch
  class Db
    SCHEMA = <<-SQL
      -- 1: Cves
      CREATE TABLE IF NOT EXISTS cves (
        id            VARCHAR(64) NOT NULL PRIMARY KEY,
        description   TEXT        NOT NULL,
        cwe_id        VARCHAR(32) NOT NULL,
        cvss_v2_score REAL        NOT NULL,
        cvss_v3_score REAL        NOT NULL,
        published     DATETIME    NOT NULL,
        last_modified DATETIME    NOT NULL
      );

      -- 2: Products
      CREATE TABLE IF NOT EXISTS products (
        id      INTEGER PRIMARY KEY,
        vendor  VARCHAR(255) NOT NULL,
        name    VARCHAR(255) NOT NULL,
        version VARCHAR(255) NOT NULL
      );

      CREATE UNIQUE INDEX IF NOT EXISTS products_unique
        ON products (vendor, name, version);

      --3: Cves <-> Products
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
    SQL

    def self.migrate
      db.transaction do |txn|
        SCHEMA.split(";\n").each do |query|
          query = query.strip
          query += ";" unless query.ends_with?(";")

          Log.info { "Migrating: \n" + query }
          txn.connection.exec(query)
        end
      end
    end
  end
end
